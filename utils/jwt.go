package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/lsy88/jsonwizard/global"
	"github.com/lsy88/jsonwizard/model/request"
	"time"
)

var (
	TokenExpired     = errors.New("Token is expired")            //token过期
	TokenNotValidYet = errors.New("Token not active yet")        //token未生效
	TokenMalformed   = errors.New("That's not even a token")     //token为生成
	TokenInvalid     = errors.New("Couldn't handle this token:") //无法处理此令牌
)

type JWT struct {
	SigningKey []byte
}

func NewJwt() *JWT {
	return &JWT{
		SigningKey: []byte(global.JW_CONFIG.JWT.SigningKey),
	}
}

func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	bf, _ := ParseDuration(global.JW_CONFIG.JWT.BufferTime)
	ep, _ := ParseDuration(global.JW_CONFIG.JWT.ExpiresTime)
	claims := request.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,    // 签名生效时间
			ExpiresAt: time.Now().Add(ep).Unix(),   // 过期时间 7天  配置文件
			Issuer:    global.JW_CONFIG.JWT.Issuer, // 签名的发行者
		},
	}
	return claims
}

//CreateToken 根据claims生成token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧的token换成新的token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims request.CustomClaims) (string, error) {
	v, err, _ := global.JW_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				fmt.Println(tokenString)
				fmt.Println(token)
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
		
	} else {
		return nil, TokenInvalid
		
	}
	
}
