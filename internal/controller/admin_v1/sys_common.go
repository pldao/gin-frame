package admin_v1

import (
	"github.com/PLDao/gin-frame/internal/controller"
)

type CommonController struct {
	controller.Api
}

func NewCommonController() *CommonController {
	return &CommonController{}
}
