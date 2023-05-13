package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lsy88/jsonwizard/global"
	"github.com/lsy88/jsonwizard/global/response"
	"github.com/lsy88/jsonwizard/model"
	"github.com/lsy88/jsonwizard/model/request"
	systemresp "github.com/lsy88/jsonwizard/model/response"
	"github.com/lsy88/jsonwizard/utils"
	"go.uber.org/zap"
)

type BaseApi struct{}

func (b *BaseApi) Login(c *gin.Context) {
	var user request.LoginUserRequest
	if err := c.ShouldBind(&user); err != nil {
		response.FailWithDetailed("参数绑定失败", err.Error(), c)
		return
	}
	err := utils.Verify(&user, utils.LoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if store.Verify(user.CaptchaId, user.Captcha, true) {
		u := model.JW_User{
			UserName: user.Username,
			Password: user.Password,
		}
		user, err := userService.Login(&u)
		if err != nil {
			global.JW_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			response.FailWithMessage("用户名不存在或者密码错误", c)
			return
		}
		b.tokenNext(c, *user)
		return
	}
	response.FailWithMessage("验证码错误", c)
}

//登录合法之后签发token
func (b *BaseApi) tokenNext(c *gin.Context, user model.JW_User) {
	j := &utils.JWT{SigningKey: []byte(global.JW_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(request.BaseClaims{
		ID:       user.ID,
		Username: user.UserName,
		RealName: user.RealName,
		UserType: user.Type,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.JW_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	//单点登录
	if !global.JW_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(systemresp.UserLoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000, //返回过期时间
		}, "登录成功", c)
	}
	
}

func (b *BaseApi) Register(c *gin.Context) {
	var u request.RegisterUserRequest
	if err := c.ShouldBindJSON(&u); err != nil {
		response.FailWithDetailed("参数绑定失败", err.Error(), c)
		return
	}
	err := utils.Verify(u, utils.RegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user := model.JW_User{
		JW_Model: global.JW_Model{
			CreateGUID: utils.GenGUID(),
		},
		EmployeeId:  u.EmployeeId,
		NickName:    u.NickName,
		RealName:    u.RealName,
		PhoneNumber: u.PhoneNumber,
		UserName:    u.UserName,
		Password:    u.Password,
		Type:        u.Type,
		Avatar:      u.Avatar,
		WeiBo:       u.WeiBo,
		WeChat:      u.WeChat,
		QQ:          u.QQ,
		Email:       u.Email,
		Profile:     u.Profile,
		Birthday:    u.Birthday,
		Gender:      u.Gender,
		SoftDelete:  0,
	}
	userReturn, err := userService.Register(user)
	if err != nil {
		global.JW_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(userReturn, "注册失败", c)
		return
	}
	response.OkWithMessage("注册成功", c)
}
