package model

import "github.com/lsy88/jsonwizard/global"

//应用请求参数表
type JW_Param struct {
	global.JW_Model
	AppId            int    `json:"app_id" gorm:"type:int(11) NOT NULL; index:app_id; COMMENT:'所属应用id'"`
	Name             string `json:"name" gorm:"type:varchar(64) DEFAULT NULL; COMMENT:'请求参数名'"`
	Value            string `json:"value" gorm:"type:text; COMMENT:'参数值,e.g:城市id/城市名'"`
	Title            string `json:"title" gorm:"type:varchar(64) DEFAULT NULL; COMMENT:'表单展示标题,e.g.城市'"`
	AssociateUrl     string `json:"associate_url" gorm:"column:associate_url; type:varchar(1024) DEFAULT NULL; COMMENT:'参数关联网址'"`
	AssociateUrlStop int    `json:"associate_url_stop" gorm:"column:associate_url_stop; type:tinyint(3) unsigned NOT NULL DEFAULT '0'; COMMENT:'未停用/已停用：0/1'"`
	SoftDelete       int    `json:"soft_delete" gorm:"type:tinyint(3) unsigned NOT NULL DEFAULT '0'; index:app_id; COMMENT:'未删除/已删除：0/1'"`
}
