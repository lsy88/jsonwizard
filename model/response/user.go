package response

type UserInfoResp struct {
	UserId     int    `json:"user_id"`
	UserName   string `json:"user_name"`
	NickName   string `json:"nick_name"`
	RealName   string `json:"real_name"`
	Type       int    `json:"type"`
	Avatar     string `json:"avatar"`
	EmployeeId string `json:"employee_id"`
}
