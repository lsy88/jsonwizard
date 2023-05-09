package model

import (
	"github.com/lsy88/jsonwizard/global"
)

// 应用管理表

type JW_Application struct {
	global.JW_Model
	NameCN             string `json:"name_cn" gorm:"name_cn; index:name_cn; type:varchar(64) not null;COMMENT:'应用中文名'"`
	NameEN             string `json:"name_en" gorm:"name_en; index:name_en; type:varchar(64) not null;COMMENT:'应用英文名'"`
	Description        string `json:"description" gorm:"description; type:varchar(2048) default null;COMMENT:'应用描述'"`
	OwnerId            int    `json:"owner_id" gorm:"owner_id; index:owner_id; type:int(11) not null;COMMENT:'应用所有者ID'"`
	OperationManagerId int    `json:"operation_manager_id" gorm:"type:int(11) DEFAULT '0';COMMENT:'运营人员ID'"`
	AssociateUrl       string `json:"associate_url" gorm:"type:varchar(1024) DEFAULT NULL;COMMENT:'关联网址'"`
	SoftDelete         int    `json:"soft_delete" gorm:"type:tinyint(3) unsigned NOT NULL DEFAULT '0'; index:updater_id; index:name_cn; index:name_en; index:owner_id; COMMENT:'未删除/已删除：0/1'"`
}

func (JW_Application) TableName() string {
	return "jw_application"
}
