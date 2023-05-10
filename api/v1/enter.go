package v1

import "github.com/lsy88/jsonwizard/service"

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	BaseApi
}

var (
	userService = service.ServiceGroupApp.UserService
)
