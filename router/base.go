package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lsy88/jsonwizard/api/v1"
)

type BaseRouter struct{}

func (u *BaseRouter) InitUserRouter(router *gin.RouterGroup) gin.IRoutes {
	baseRouter := router.Group("base")
	baseApi := v1.ApiGroupApp.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("register", baseApi.Register)
		baseRouter.POST("getCaptcha", baseApi.GenCaptcha)
		baseRouter.POST("verifyCaptcha", baseApi.VerifyCaptcha)
	}
	return baseRouter
}
