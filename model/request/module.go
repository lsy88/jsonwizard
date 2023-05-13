package request

type ModuleCreateReq struct {
	NameCN         string `json:"name_cn"`
	NameEN         string `json:"name_en"`
	AppId          string `json:"app_id"`
	Definition     string `json:"definition"`
	UISchema       string `json:"ui_schema"`
	Sort           int    `json:"sort"`
	AssociationUrl string `json:"association_url"`
}
