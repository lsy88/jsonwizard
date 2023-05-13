package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lsy88/jsonwizard/api/v1"
)

type ModuleRouter struct{}

func (m *ModuleRouter) InitModuleRouter(router *gin.RouterGroup) gin.IRoutes {
	moduleRouter := router.Group("module")
	moduleApi := v1.ApiGroupApp.ModuleApi
	{
		moduleRouter.POST("create", moduleApi.Create)
	}
	return moduleRouter
}
