package admin_v1

import (
	"github.com/PLDao/gin-frame/internal/controller"
	"github.com/PLDao/gin-frame/internal/service/user_social"
	"github.com/PLDao/gin-frame/internal/validator"
	"github.com/PLDao/gin-frame/internal/validator/form"
	"github.com/gin-gonic/gin"
)

type UserSocialController struct {
	controller.Api
}

func NewUserSocialController() *UserSocialController {
	return &UserSocialController{}
}

func (api UserSocialController) AddUserSocial(c *gin.Context) {
	// 初始化参数结构体
	socialForm := form.NewUserSocialForm()
	// 绑定参数并使用验证器验证参数
	if err := validator.CheckPostParams(c, &socialForm); err != nil {
		return
	}
	err := user_social.NewUserAddSocialController().AddSocial(socialForm)
	if err != nil {
		api.Err(c, err)
		return
	}
	api.Success(c, nil)
}

func (api UserSocialController) GetAllUserSocial(c *gin.Context) {
	userName := form.NewUserNameForm()
	if err := validator.CheckQueryParams(c, &userName); err != nil {
		return
	}
	res, err := user_social.NewUserAddSocialController().ListSocial(userName)
	if err != nil {
		api.Err(c, err)
		return
	}
	api.Success(c, res)
}

func (api UserSocialController) UpdateSocial(c *gin.Context) {
	socialForm := form.NewUserSocialForm()
	if err := validator.CheckPostParams(c, &socialForm); err != nil {
		return
	}
	err := user_social.NewUserAddSocialController().UpdateSocial(socialForm)
	if err != nil {
		api.Err(c, err)
		return
	}
	api.Success(c, nil)
}
