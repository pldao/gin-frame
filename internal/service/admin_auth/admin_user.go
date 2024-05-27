package admin_auth

import (
	"github.com/PLDao/gin-frame/internal/model"
	"github.com/PLDao/gin-frame/internal/pkg/errors"
	"github.com/PLDao/gin-frame/internal/resources"
	"github.com/PLDao/gin-frame/internal/service"
	"github.com/PLDao/gin-frame/internal/validator/form"
)

// AdminUserService 授权服务
type AdminUserService struct {
	service.Base
}

func NewAdminUserService() *AdminUserService {
	return &AdminUserService{}
}

// GetUserInfo 获取用户信息
func (a *AdminUserService) GetUserInfo(id uint) (*resources.AdminUserResources, error) {
	// 查询用户是否存在
	adminUsersModel := model.NewAdminUsers()
	user := adminUsersModel.GetUserById(id)
	if user != nil {
		return resources.NewAdminUserResources(user), nil
	}
	return nil, errors.NewBusinessError(errors.FAILURE, "获取用户信息失败")
}

func (a *AdminUserService) Register(params *form.AddAdminUser) error {
	adminUsersModel := model.NewAdminUsers()
	if adminUsersModel.GetUserInfo(params.UserName) != nil {
		return errors.NewBusinessError(errors.FAILURE, "用户名已存在")
	}
	adminUsersModel.Username = params.UserName
	adminUsersModel.Nickname = "simple_user"
	adminUsersModel.Password = params.PassWord
	adminUsersModel.Mobile = params.Mobile
	adminUsersModel.Email = params.Email
	adminUsersModel.IsAdmin = 0
	adminUsersModel.Status = 1
	return adminUsersModel.Register()
}
