package core

import (
	"github.com/gin-gonic/gin"
	"github.com/lsy88/jsonwizard/global"
	"net/http"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r.StaticFS(global.JW_CONFIG.Local.Path, http.Dir(global.JW_CONFIG.Local.StorePath)) // 为用户头像和文件提供静态地址
	PublicGroup := r.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	return r
}
