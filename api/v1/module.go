package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lsy88/jsonwizard/global"
	"github.com/lsy88/jsonwizard/global/response"
	"github.com/lsy88/jsonwizard/model"
	"github.com/lsy88/jsonwizard/model/request"
	"github.com/lsy88/jsonwizard/utils"
	"go.uber.org/zap"
	"strconv"
	"strings"
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

//模块信息查询
func (m *ModuleApi) ModuleInfo(c *gin.Context) {
	var in = request.GetModuleInfo{
		AppId:    c.Query("app_id"),
		ModuleId: c.Query("module_id"),
	}
	err := utils.Verify(&in, utils.ModuleInfoVerify)
	if err != nil {
		global.JW_LOG.Error("校验参数错误", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	getInfo, err := moduleService.GetInfo(&in)
	if err != nil {
		response.FailWithDetailed(err.Error(), "查询模块信息失败", c)
		return
	}
	response.OkWithDetailed(getInfo, "查询成功", c)
}

//模块编辑
func (m *ModuleApi) Modify(c *gin.Context) {
	c.Request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	var mu request.ModuleModifyReq
	if err := c.ShouldBind(&mu); err != nil {
		global.JW_LOG.Error("绑定参数错误", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := utils.Verify(&mu, utils.ModuleModifyVerify)
	if err != nil {
		global.JW_LOG.Error("校验参数错误", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = moduleService.Modify(&mu)
	if err != nil {
		response.FailWithDetailed(err.Error(), "模块编辑失败", c)
		return
	}
	response.Ok(c)
}

//模块状态列表
func (m *ModuleApi) GetStatusList(c *gin.Context) {
	appId, _ := strconv.Atoi(c.Query("app_id"))
	params := c.Query("params")
	if appId < 0 || params == "" {
		global.JW_LOG.Error("绑定参数错误")
		response.FailWithMessage(errors.New("绑定参数错误").Error(), c)
		return
	}
	param := strings.Split(params, "/")
	//todo 参数格式校验
	if len(param) != 4 {
		response.FailWithMessage(errors.New("参数格式错误").Error(), c)
		return
	}
	moduleId, _ := strconv.Atoi(string(params[0]))
	status, _ := strconv.Atoi(string(params[3]))
	list, err := moduleService.GetStatusList(appId, moduleId, status)
	if err != nil {
		global.JW_LOG.Error("获取列表错误", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(list, "获取成功", c)
}
