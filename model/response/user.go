package response

import "github.com/lsy88/jsonwizard/model"

type UserInfoResp struct {
	Id         int    `json:"id"`
	UserName   string `json:"user_name"`
	NickName   string `json:"nick_name"`
	RealName   string `json:"real_name"`
	Type       int    `json:"type"`
	Avatar     string `json:"avatar"`
	EmployeeId string `json:"employee_id"`
}

type UserInfoListResp struct {
	Key        string `json:"key"`
	UserId     int    `json:"user_id"`
	RealName   string `json:"real_name"`
	EmployeeId string `json:"employee_id"`
}

type UserAuthorizedListResp struct {
	Key      string `json:"key"`
	UserId   int    `json:"user_id"`
	RealName string `json:"real_name"`
}

type UserLoginResponse struct {
	User      model.JW_User `json:"user"`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expiresAt"`
}
