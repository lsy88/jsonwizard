package service

import "github.com/lsy88/jsonwizard/service/system"

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	UserService        UserService
	ApplicationService ApplicationService
	ModuleService      ModuleService
	ParamService       ParamService
}

var ServiceGroupApp = new(ServiceGroup)
