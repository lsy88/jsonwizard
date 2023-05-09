package model

import "time"

//用户表
type JW_User struct {
	EmployeeId  string    `json:"employee_id" gorm:"type:varchar(64) NOT NULL; COMMENT:'工号'"`
	NickName    string    `json:"nick_name" gorm:"type:varchar(64) DEFAULT NULL; index:nickname; COMMENT:'昵称(自动生成)'"`
	RealName    string    `json:"real_name" gorm:"type:varchar(64) DEFAULT NULL; COMMENT '真实姓名'"`
	PhoneNumber string    `json:"phone_number" gorm:"type:varchar(32) DEFAULT NULL; index:phone_number; index:phone_number_2; COMMENT '手机号'"`
	UserName    string    `json:"user_name"  gorm:"type:varchar(64) NOT NULL; index:username; index:username_2; COMMENT:'用户名'"`
	Password    string    `json:"password" gorm:"type:varchar(32) NOT NULL; index:username_2; index:phone_number_2; COMMENT:'登录密码'"`
	Type        int       `json:"type" gorm:"type:tinyint(3) unsigned NOT NULL DEFAULT '4';COMMENT:'用户类型：1：超管，2：开发主管；3：运营主管；4：普通开发；5：普通运营'"`
	Avatar      string    `json:"avatar" gorm:"type:varchar(1024) DEFAULT NULL; COMMENT:'头像图片地址'"`
	WeiBo       string    `json:"wei_bo" gorm:"type:varchar(1024) DEFAULT NULL; COMMENT:'绑定微博'"`
	WeChat      string    `json:"wechat" gorm:"type:varchar(1024) DEFAULT NULL; COMMENT:'绑定微信'"`
	QQ          string    `json:"qq" gorm:"type:varchar(1024) DEFAULT NULL; COMMENT:'绑定qq'"`
	Email       string    `json:"email" gorm:"type:varchar(1024) DEFAULT NULL; COMMENT:'绑定邮箱'"`
	Profile     string    `json:"profile" gorm:"type:varchar(1024) DEFAULT NULL; COMMENT:'个人简介'"`
	Birthday    time.Time `json:"birthday" gorm:"type:datetime DEFAULT NULL; COMMENT:'生日'"`
	Gender      int       `json:"gender" gorm:"type:tinyint(3) unsigned NOT NULL DEFAULT '0'; COMMENT:'性别：1：男，2：女'"`
	SoftDelete  int       `json:"soft_delete" gorm:"type:tinyint(3) unsigned NOT NULL DEFAULT '0'; index:username; index:nickname; index:username_2; index:phone_number; index:phone_number_2; COMMENT:'未删除/已删除：0/1'"`
}
