package request

import (
	"github.com/dgrijalva/jwt-go"
)

// JWT Custom claims structure
type CustomClaims struct {
	Id         int
	UserName   string
	BufferTime int64
	jwt.StandardClaims
	BaseClaims
}

type BaseClaims struct {
	ID       int
	Username string
	RealName string
	UserType int //用户角色
}

// User cache structure
type UserCache struct {
	ID   int `redis:"id"`
	Type int `redis:"type"`
}

// User cache structure
type UserCacheRedis struct {
	ID   int `redis:"id"`
	Type int `redis:"type"`
}

type GetUserInfoRequest struct {
	Id int `json:"id"`
}

type LoginUserRequest struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}
