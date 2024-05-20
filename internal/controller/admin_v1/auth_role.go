package admin_v1

import "github.com/PLDao/gin-frame/internal/controller"

type RoleController struct {
	controller.Api
}

func NewRoleController() *RoleController {
	return &RoleController{}
}
