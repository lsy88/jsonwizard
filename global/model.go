package global

import (
	"database/sql"
	"time"
)

type JW_Model struct {
	ID         int          `json:"id" gorm:"id; primaryKey; COMMENT:'所属应用ID'"`
	CreateGUID string       `json:"create_guid" gorm:"type:varchar(36) NOT NULL; COMMENT:'全局数据唯一标识'"`
	CreatorId  int          `json:"creator_id" gorm:"type:int(11) NOT NULL;COMMENT:'应用创建人ID'"`
	UpdaterId  int          `json:"updater_id" gorm:"index:updater_id;type:int(11) NOT NULL;COMMENT:'最后更新人ID'"`
	CreateTime time.Time    `gorm:"type:datetime NOT NULL;COMMENT:'创建时间'"`
	UpdateTime time.Time    `gorm:"type:datetime NOT NULL;COMMENT:'最后修改时间'"`
	Del        int          `json:"del" gorm:"del;index:name_cn name_en owner_id updater_id;type:tinyint(3) unsigned NOT NULL DEFAULT '0';COMMENT:'未删除/已删除：0/1'"`
	DeletedAt  sql.NullTime `gorm:"index"`
}
