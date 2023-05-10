package service

import "github.com/lsy88/jsonwizard/service/system"

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	UserService        UserService
}

var ServiceGroupApp = new(ServiceGroup)
