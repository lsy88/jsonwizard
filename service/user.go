package service

import (
	"errors"
	"fmt"
	"github.com/lsy88/jsonwizard/global"
	"github.com/lsy88/jsonwizard/model"
	"github.com/lsy88/jsonwizard/model/response"
	"github.com/lsy88/jsonwizard/utils"
	"gorm.io/gorm"
	"strconv"
)

type UserService struct{}

var (
	UserExistError    = errors.New("用户名已注册")
	UserNotExistError = errors.New("该用户不存在")
)

func (u *UserService) FindUserById(id int) error {
	_, err := u.GetUserInfo(id)
	return err
}

func (u *UserService) Login(user *model.JW_User) (us *model.JW_User, err error) {
	user.Password = utils.MD5([]byte(user.Password))
	err = global.JW_DB.DB.Where("user_name = ? AND password = ? AND soft_delete = 0", user.UserName, user.Password).First(&us).Error
	return
}

func (u *UserService) Register(us model.JW_User) (userInter model.JW_User, err error) {
	var user *model.JW_User
	//用户已存在
	if !errors.Is(global.JW_DB.DB.Where("user_name = ? AND soft_delete = 0", us.UserName).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, UserExistError
	}
	us.Password = utils.MD5([]byte(us.Password))
	err = global.JW_DB.DB.Create(&u).Error
	return us, err
}

//获取用户信息
func (u *UserService) GetUserInfo(id int) (user *response.UserInfoResp, err error) {
	err = global.JW_DB.DB.Table(model.JW_User{}.TableName()).Where("id = ? AND soft_delete = 0", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, UserNotExistError
	}
	return
}

//获取用户列表
func (u *UserService) GetUserInfoList(userType int) (users *[]response.UserInfoListResp, total int64, err error) {
	var us []model.JW_User
	err = global.JW_DB.DB.Table(model.JW_User{}.TableName()).Where("type < ? AND soft_delete = 0", userType).Count(&total).Find(&us).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}
	for _, u := range us {
		*users = append(*users, response.UserInfoListResp{
			Key:        "user-" + strconv.Itoa(u.ID),
			UserId:     u.ID,
			RealName:   u.RealName,
			EmployeeId: u.EmployeeId,
		})
	}
	return
}

//授权用户列表
func (u *UserService) GetUserAuthorizedList(appId int) (users *[]response.UserAuthorizedListResp, total int64, err error) {
	err = global.JW_DB.DB.Transaction(func(tx *gorm.DB) error {
		var auth []model.JW_Auth
		err = tx.Table(model.JW_Auth{}.TableName()).Where("app_id = ? AND soft_delete = 0", appId).Count(&total).Find(&auth).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		for _, u := range auth {
			var a model.JW_User
			err = tx.Table(model.JW_User{}.TableName()).Where("id = ? AND soft_delete = 0", u.UserId).First(&a).Error
			if err != nil {
				fmt.Println(u.UserId, "该用户出现问题")
				continue
			}
			t := response.UserAuthorizedListResp{
				Key:      "user-" + strconv.Itoa(a.ID),
				UserId:   a.ID,
				RealName: a.RealName,
			}
			*users = append(*users, t)
		}
		return nil
	})
	
	return
}
