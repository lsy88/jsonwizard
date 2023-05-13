package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lsy88/jsonwizard/api/v1"
)

type ApplicationRouter struct{}

func (u *ApplicationRouter) InitAppRouter(router *gin.RouterGroup) gin.IRoutes {
	appRouter := router.Group("application")
	appApi := v1.ApiGroupApp.AppApi
	{
		appRouter.POST("appCreate", appApi.Create)
	}
	return appRouter
}
