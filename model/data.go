package model

//模块真实数据
type JW_Data struct {
	Params     string `json:"params" gorm:"type:varchar(128) DEFAULT NULL; index:params; COMMENT:'请求参数的唯一标示：:isTemp(y/n)/app_id/module_id/param_1/value_1/param_2/value_2'"`
	ModuleId   int    `json:"module_id" gorm:"type:int(11) NOT NULL; COMMENT:'对应模块id'"`
	Data       string `json:"data" gorm:"type:JSON NOT NULL; COMMENT:'真实数据'"`
	IsStop     int    `json:"is_stop" gorm:"type:tinyint(3) unsigned NOT NULL DEFAULT '0'; COMMENT:'未停用/已停用：0/1'"`
	SoftDelete int    `json:"soft_delete" gorm:"type:tinyint(3) unsigned NOT NULL DEFAULT '0'; index:params; COMMENT:'未删除/已删除：0/1'"`
}

func (JW_Data) TableName() string {
	return "jw_data"
}
