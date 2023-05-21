package service

import (
	"github.com/lsy88/jsonwizard/global"
	"github.com/lsy88/jsonwizard/model"
	"github.com/lsy88/jsonwizard/model/response"
)

type ParamService struct{}

func (p *ParamService) GetParamList(appId string) (list []*response.ParamSelectList, count int64, err error) {
	err = global.JW_DB.DB.Table(model.JW_Param{}.TableName()).Where("app_id = ?", appId).Count(&count).Find(&list).Error
	return
}
