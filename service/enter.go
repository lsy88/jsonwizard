package service

import "github.com/lsy88/jsonwizard/service/system"

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	UserService        UserService
	ApplicationService ApplicationService
	ModuleService      ModuleService
}

var ServiceGroupApp = new(ServiceGroup)
