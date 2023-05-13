package service

import (
	"github.com/lsy88/jsonwizard/global"
	"github.com/lsy88/jsonwizard/model"
)

type ApplicationService struct{}

func (a *ApplicationService) FindByID(id int) (app *model.JW_Application, ok bool) {
	err := global.JW_DB.DB.Table(model.JW_Application{}.TableName()).Where("id = ?", id).First(&app).Error
	if err != nil {
		return nil, false
	}
	return app, true
}

func (a *ApplicationService) FindByName(cn, en string) (app *model.JW_Application, ok bool) {
	err := global.JW_DB.DB.Table(model.JW_Application{}.TableName()).Where("name_cn = ? AND name_en = ?", cn, en).First(&app).Error
	if err != nil {
		return nil, false
	}
	return app, true
}

func (a *ApplicationService) Create(app *model.JW_Application) error {
	err := global.JW_DB.DB.Where(model.JW_Application{}.TableName()).Create(app).Error
	return err
}
