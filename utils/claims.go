package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/lsy88/jsonwizard/global"
	"github.com/lsy88/jsonwizard/global/response"
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

// 从Gin的Context中获取从jwt解析出来的用户是不是管理员
func GetUserIsAdmin(c *gin.Context) bool {
	if claims, exists := c.Get("claims"); !exists {
		global.JW_LOG.Error("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件!")
		return false
	} else {
		//waitUse := claims.(*request.UserCache)
		//return waitUse.IsAdmin
	}
}

func CheckRole(c *gin.Context) bool {
	if !GetUserIsAdmin(c) {
		response.FailWithMessage("权限不足，请联系管理员进行操作", c)
		return false
	}
	return true
}
