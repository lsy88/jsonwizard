package model

import "github.com/lsy88/jsonwizard/global"

//应用模块表
type JW_Module struct {
	global.JW_Model
	AppId          int    `json:"app_id" gorm:"type:int(11) NOT NULL; index:app_id; COMMENT:'所属应用id'"`
	NameCN         string `json:"name_cn" gorm:"name_cn; type:varchar(64) NOT NULL; COMMENT:'模块中文名'"`
	NameEN         string `json:"name_en" gorm:"name_en; type:varchar(64) NOT NULL; COMMENT:'模块英文名'"`
	Definition     string `json:"definition" gorm:"type:JSON; COMMENT:'模块定义'"`
	UISchema       string `json:"ui_schema" gorm:"type:JSON; COMMENT:'模块UI Schema'"`
	Sort           int    `json:"sort" gorm:"type:tinyint(3) unsigned NOT NULL DEFAULT '0'; COMMENT:'模块排序'"`
	AssociationUrl string `json:"association_url" gorm:"column:association_url; type:varchar(1024) DEFAULT NULL; COMMENT:'模块关联审核地址'"`
	IsStop         int    `json:"is_stop" gorm:"type:tinyint(3) unsigned NOT NULL DEFAULT '0'; COMMENT:'未停用/已停用：0/1'"`
	SoftDelete     int    `json:"soft_delete" gorm:"type:tinyint(3) unsigned NOT NULL DEFAULT '0'; index:app_id; COMMENT:'未删除/已删除：0/1'"`
}

func (JW_Module) TableName() string {
	return "jw_module"
}
