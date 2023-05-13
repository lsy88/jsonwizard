package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lsy88/jsonwizard/global"
	"github.com/lsy88/jsonwizard/global/response"
	"github.com/lsy88/jsonwizard/model"
	"github.com/lsy88/jsonwizard/model/request"
	"github.com/lsy88/jsonwizard/utils"
	"go.uber.org/zap"
	"strconv"
)

type ModuleApi struct{}

//创建模块
func (m *ModuleApi) Create(c *gin.Context) {
	c.Request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	var mu request.ModuleCreateReq
	if err := c.ShouldBindJSON(&mu); err != nil {
		global.JW_LOG.Error("绑定参数错误", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := utils.Verify(&mu, utils.ModuleVerify)
	if err != nil {
		global.JW_LOG.Error("校验参数错误", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	id, _ := strconv.Atoi(mu.AppId)
	module := model.JW_Module{
		JW_Model: global.JW_Model{
			CreateGUID: utils.GenGUID(),
		},
		AppId:          id,
		NameCN:         mu.NameCN,
		NameEN:         mu.NameEN,
		Definition:     mu.Definition,
		UISchema:       mu.UISchema,
		Sort:           mu.Sort,
		AssociationUrl: mu.AssociationUrl,
	}
	err = moduleService.Create(&module)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("创建模块成功", c)
}
