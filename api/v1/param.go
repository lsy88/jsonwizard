package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lsy88/jsonwizard/global"
	"github.com/lsy88/jsonwizard/global/response"
	"github.com/lsy88/jsonwizard/utils"
	"go.uber.org/zap"
)

type ParamApi struct{}

func (p *ParamApi) ParamSelectList(c *gin.Context) {
	appId := c.Query("app_id")
	err := utils.Verify(appId, utils.ParamSelectVerify)
	if err != nil {
		global.JW_LOG.Error("校验参数错误", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, count, err := paramService.GetParamList(appId)
	if err != nil {
		global.JW_LOG.Error("查询错误", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{
		"data":  list,
		"count": count,
	}, "查询成功", c)
}
