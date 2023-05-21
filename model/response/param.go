package response

type ParamSelectList struct {
	Id               int    `json:"id"`
	AppId            int    `json:"app_id"`
	Name             string `json:"name"`
	Title            string `json:"title"`
	Value            string `json:"value"`
	AssociateUrl     string `json:"associate_url"`
	AssociateUrlStop int    `json:"associate_url_stop"`
	CreatorId        int    `json:"creator_id"`
	UpdaterId        int    `json:"updater_id"`
	Creator          string `json:"creator"`
	Updater          string `json:"updater"`
}
