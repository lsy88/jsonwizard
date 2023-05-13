package v1

import "github.com/lsy88/jsonwizard/service"

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	BaseApi
	AppApi
	ModuleApi
}

var (
	userService   = service.ServiceGroupApp.UserService
	appService    = service.ServiceGroupApp.ApplicationService
	moduleService = service.ServiceGroupApp.ModuleService
)
