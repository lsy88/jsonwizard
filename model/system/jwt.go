package system

import "github.com/lsy88/jsonwizard/global"

type JwtBlacklist struct {
	global.JW_Model
	Jwt string `gorm:"not null;type:text;comment:jwt"`
}

func (r *JwtBlacklist) TableName() string {
	return "jwt_blacklist"
}
