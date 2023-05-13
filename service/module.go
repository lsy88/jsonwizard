package service

import (
	"github.com/lsy88/jsonwizard/global"
	"github.com/lsy88/jsonwizard/model"
)

type ModuleService struct{}

func (m *ModuleService) Create(mu *model.JW_Module) error {
	err := global.JW_DB.DB.Table(model.JW_Module{}.TableName()).Create(&mu).Error
	return err
}
