package global

import (
	"time"
)

//基础字段表
type JW_Model struct {
	ID         int       `json:"id" gorm:"primarykey"`
	CreateGUID string    `json:"create_guid" gorm:"type:varchar(36) NOT NULL; COMMENT:'全局数据唯一标识'"`
	CreatorId  int       `json:"creator_id" gorm:"type:int(11) NOT NULL;COMMENT:'创建人ID'"`
	UpdaterId  int       `json:"updater_id" gorm:"type:int(11) NOT NULL; index:updater_id; COMMENT:'更新人ID'"`
	CreateTime time.Time `gorm:"type:datetime NOT NULL;COMMENT:'创建时间'"`
	UpdateTime time.Time `gorm:"type:datetime NOT NULL;COMMENT:'最后修改时间'"`
}
