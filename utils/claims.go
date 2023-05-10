package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/lsy88/jsonwizard/global"
	"github.com/lsy88/jsonwizard/model/request"
	"go.uber.org/zap"
)

//从Gin的Context中获取requestId
func GetRequestID(c *gin.Context) zap.Field {
	return zap.String("requestId", c.GetString("requestId"))
}

// 从Gin的Context中获取从jwt解析出来的用户信息
func GetClaim(c *gin.Context) *request.CustomClaims {
	if claims, exists := c.Get("tokenClaims"); !exists {
		global.JW_LOG.Error("从Gin的Context中获取从jwt解析出来的用户失败, 请检查路由是否使用jwt中间件!", GetRequestID(c))
		return nil
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse
	}
}
