package service

import (
	"github.com/lsy88/jsonwizard/global"
	"github.com/lsy88/jsonwizard/model"
	"github.com/lsy88/jsonwizard/model/request"
	"github.com/lsy88/jsonwizard/model/response"
	"gorm.io/gorm"
)

type ModuleService struct{}

func (m *ModuleService) Create(mu *model.JW_Module) error {
	err := global.JW_DB.DB.Table(model.JW_Module{}.TableName()).Create(&mu).Error
	return err
}

func (m *ModuleService) GetInfo(req *request.GetModuleInfo) (info *response.ModuleInfoResp, err error) {
	module := model.JW_Module{}
	err = global.JW_DB.DB.Table(model.JW_Module{}.TableName()).Where("id = ? AND app_id = ?", req.ModuleId, req.AppId).First(&module).Error
	if err != nil {
		return
	}
	info = &response.ModuleInfoResp{
		Id:             module.ID,
		AppId:          module.AppId,
		NameCN:         module.NameCN,
		NameEN:         module.NameEN,
		Definition:     module.Definition,
		UISchema:       module.UISchema,
		AssociationUrl: module.AssociationUrl,
		Sort:           module.Sort,
		CreatorId:      module.CreatorId,
		UpdaterId:      module.UpdaterId,
		IsStop:         module.IsStop,
	}
	return
}

func (m *ModuleService) Modify(modify *request.ModuleModifyReq) (err error) {
	mu := model.JW_Module{}
	err = global.JW_DB.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Table(mu.TableName()).Where("id = ?", modify.ModuleId).Error
		if err != nil {
			return err
		}
		err = tx.Table(mu.TableName()).Where("id = ?", modify.ModuleId).Update("definition", modify.Definition).
			Update("ui_schema", modify.UISchema).Error
		return err
	})
	return
}

func (m *ModuleService) GetStatusList(appId int, moduleId int, status int) (statusList []response.ModuleStatusListResp, err error) {
	//list := []*model.JW_Module{}
	err = global.JW_DB.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Table(model.JW_Module{}.TableName()).Where("app_id = ?", appId).Find(&statusList).Error
		if err != nil {
			return err
		}
		for i := range statusList {
			if statusList[i].Id == moduleId {
				statusList[i].ReviewStatus = status
			}
		}
		return nil
	})
	return
}
