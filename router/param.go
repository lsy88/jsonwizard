package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lsy88/jsonwizard/api/v1"
)

type ParamRouter struct{}

func (p *ParamRouter) InitParamRouter(router *gin.RouterGroup) gin.IRoutes {
	paramRouter := router.Group("param")
	paramApi := v1.ApiGroupApp.ParamApi
	{
		paramRouter.GET("selectList", paramApi.ParamSelectList)
	}
	return paramRouter
}
