package request

type CreateAppReq struct {
	NameCN      string `json:"name_cn"`
	NameEN      string `json:"name_en"`
	Description string `json:"description"`
}
