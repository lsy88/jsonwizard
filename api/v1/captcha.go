package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lsy88/jsonwizard/global"
	"github.com/lsy88/jsonwizard/global/response"
	systemReq "github.com/lsy88/jsonwizard/model/request"
	systemResp "github.com/lsy88/jsonwizard/model/response"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

//设置验证码的共享存储
var store = base64Captcha.DefaultMemStore

//生成验证码
func (b *BaseApi) GenCaptcha(c *gin.Context) {
	//支持字符、公式、验证码配置,生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.JW_CONFIG.Captcha.ImgHeight, global.JW_CONFIG.Captcha.ImgWidth, global.JW_CONFIG.Captcha.KeyLong, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, store)
	//生成随机的id，base64图像字符串
	id, b64s, err := captcha.Generate()
	if err != nil {
		global.JW_LOG.Error("验证码获取失败", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
		return
	}
	response.OkWithDetailed(systemResp.CaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.JW_CONFIG.Captcha.KeyLong,
	}, "验证码获取成功", c)
}

//验证验证码
func (b *BaseApi) VerifyCaptcha(c *gin.Context) {
	captcha := systemReq.CaptchaRequest{}
	if err := c.ShouldBind(captcha); err != nil {
		global.JW_LOG.Error("验证码传参失败", zap.Error(err))
		response.FailWithMessage("验证码绑定参数失败", c)
		return
	}
	//clear是否内存清理掉图片
	if verify := store.Verify(captcha.CaptchaId, captcha.PicPath, true); !verify {
		response.FailWithMessage("验证码验证失败", c)
		return
	}
	response.OkWithMessage("验证码验证成功", c)
}
