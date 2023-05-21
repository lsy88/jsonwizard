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

type GetModuleInfo struct {
	AppId    string `json:"app_id"`
	ModuleId string `json:"module_id"`
}

type ModuleModifyReq struct {
	Definition string `json:"definition" form:"definition"`
	ModuleId   string `json:"module_id" form:"module_id"`
	UISchema   string `json:"ui_schema" form:"ui_schema"`
}
