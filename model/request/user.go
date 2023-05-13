package request

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JWT Custom claims structure
type CustomClaims struct {
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

//用户注册
type RegisterUserRequest struct {
	EmployeeId  string    `json:"employee_id"`
	NickName    string    `json:"nick_name"`
	RealName    string    `json:"real_name"`
	PhoneNumber string    `json:"phone_number" `
	UserName    string    `json:"user_name"`
	Password    string    `json:"password"`
	Type        int       `json:"type"`
	Avatar      string    `json:"avatar"`
	WeiBo       string    `json:"wei_bo"`
	WeChat      string    `json:"wechat"`
	QQ          string    `json:"qq"`
	Email       string    `json:"email"`
	Profile     string    `json:"profile"`
	Birthday    time.Time `json:"birthday"`
	Gender      int       `json:"gender"`
}
