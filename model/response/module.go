package response

//模块信息查询实例
type ModuleInfoResp struct {
	Id             int    `json:"id"`
	AppId          int    `json:"app_id"`
	NameCN         string `json:"name_cn"`
	NameEN         string `json:"name_en"`
	Definition     string `json:"definition"`
	UISchema       string `json:"ui_schema"`
	AssociationUrl string `json:"association_url"`
	Sort           int    `json:"sort"`
	CreatorId      int    `json:"creator_id"`
	UpdaterId      int    `json:"updater_id"`
	IsStop         int    `json:"is_stop"`
}

type ModuleStatusListResp struct {
	Id             int    `json:"id"`
	AppId          int    `json:"app_id"`
	NameCn         string `json:"name_cn"`
	NameEn         string `json:"name_en"`
	AssociationUrl string `json:"association_url"`
	Sort           int    `json:"sort"`
	CreatorId      int    `json:"creator_id"`
	UpdaterId      int    `json:"updater_id"`
	IsStop         int    `json:"is_stop"`
	Key            string `json:"key"`
	ReviewStatus   int    `json:"review_status"` //审核状态
	Creator        string `json:"creator"`
	Updater        string `json:"updater"`
}
