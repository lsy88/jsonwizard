package request

import "github.com/dgrijalva/jwt-go"

// JWT Custom claims structure
type CustomClaims struct {
	Id         int
	UserName   string
	BufferTime int64
	jwt.StandardClaims
}

type GetUserInfoRequest struct {
	Id int `json:"id"`
}

type LoginUserRequest struct {
}
