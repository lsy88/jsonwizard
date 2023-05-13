package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lsy88/jsonwizard/global"
	"github.com/lsy88/jsonwizard/global/response"
	"github.com/lsy88/jsonwizard/model"
	"github.com/lsy88/jsonwizard/model/request"
	"github.com/lsy88/jsonwizard/utils"
	"go.uber.org/zap"
)

type AppApi struct{}

//创建应用
func (a *AppApi) Create(c *gin.Context) {
	c.Request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	var app request.CreateAppReq
	if err := c.ShouldBindJSON(&app); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := utils.Verify(a, utils.AppVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	claims, ok := c.Get("tokenClaims")
	if !ok {
		global.JW_LOG.Error("token claims noe exist！", zap.Error(err))
		response.FailWithDetailed("token claims noe exist！", err.Error(), c)
		return
	}
	m := &model.JW_Application{
		JW_Model: global.JW_Model{
			CreateGUID: utils.GenGUID(),
		},
		NameCN:      app.NameCN,
		NameEN:      app.NameEN,
		Description: app.Description,
		OwnerId:     claims.(request.CustomClaims).ID,
	}
	_, ok = appService.FindByName(m.NameCN, m.NameEN)
	if ok {
		response.FailWithMessage("应用已存在", c)
		return
	}
	err = appService.Create(m)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("应用创建成功", c)
}
