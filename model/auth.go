package model

import "github.com/lsy88/jsonwizard/global"

//应用权限拥有者表
type JW_Auth struct {
	global.JW_Model
	AppId      int `json:"app_id" gorm:"type:int(11) NOT NULL; index:app_id_2; index:app_id; COMMENT:'所属应用id'"`
	UserId     int `json:"user_id" gorm:"type:int(11) NOT NULL; index:app_id_2; index:user_id; COMMENT:'用户id'"`
	SoftDelete int `json:"soft_delete" gorm:"type:tinyint(3) unsigned NOT NULL DEFAULT '0'; index:app_id; index:user_id; COMMENT:'未删除/已删除：0/1'"`
}

func (JW_Auth) TableName() string {
	return "jw_auth"
}
