package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lsy88/jsonwizard/global"
	"github.com/lsy88/jsonwizard/global/response"
	"github.com/lsy88/jsonwizard/utils"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

// JWTAuth
// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息
//这里前端需要把token存储到cookie或者本地localStorage中
//不过需要跟后端协商过期时间
//可以约定刷新令牌或者重新登录
func JWTAuth() gin.HandlerFunc {
	reload := gin.H{"reload": true}
	tokenPrefix := "Bearer"
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.FailWithDetailed(reload, "未登录或者非法访问", c)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == tokenPrefix) {
			response.FailWithDetailed(reload, "请求头中auth格式有误", c)
			c.Abort()
			return
		}
		token := parts[1]
		//if jwtService.IsBlacklist(token) {
		//	response.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
		//	c.Abort()
		//	return
		//}
		j := utils.NewJwt()
		// parts[1]是获取到的token，我们使用之前定义好的解析JWT的函数来解析它获取claims
		claims, err := j.ParseToken(token)
		if err != nil {
			response.FailWithDetailed(reload, err.Error(), c)
			c.Abort()
			return
		}
		//校验ID是否有效
		if err, _ = userService.FindUserById(claims.ID); err != nil {
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		//token 续约
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.JW_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = time.Now().Add(dr).Unix()
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new Token", tokenPrefix+" "+newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			//	多点登录拦截
			if !global.JW_CONFIG.System.UseMultipoint {
				err, RedisJwtToken := jwtService.GetRedisJWT(newClaims.Username)
				if err != nil {
					global.JW_LOG.Error("get redis jwt failed", zap.Any("err", err), utils.GetRequestID(c))
				} else {
					_ = jwtService.JoinInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
				}
			}
			_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
		}
		if userInfo, err := jwtService.GetRedisUserInfo(claims.ID); err != nil {
			fmt.Println(err.Error())
			response.FailWithDetailed(reload, "登录过期，请重新登录", c)
			c.Abort()
		} else {
			c.Set("claims", &userInfo)
			c.Set("tokenClaims", claims)
			c.Next()
		}
	}
}
