package user_social

import (
	model "github.com/PLDao/gin-frame/internal/model/mongo"
	e "github.com/PLDao/gin-frame/internal/pkg/errors"
	"github.com/PLDao/gin-frame/internal/resources"
	"github.com/PLDao/gin-frame/internal/service"
	"github.com/PLDao/gin-frame/internal/validator/form"
)

type UserAddSocialController struct {
	service.Base
}

func NewUserAddSocialController() *UserAddSocialController {
	return &UserAddSocialController{}
}

func (s UserAddSocialController) AddSocial(param *form.UserSocial) error {
	socialModel := model.NewSocial()
	err := socialModel.AddSocial(param)
	if err != nil {
		err := e.NewBusinessError(e.MONGOERROR, "mongo error")
		return err
	}
	return nil
}

func (s UserAddSocialController) ListSocial(param *form.UserName) (*resources.SocialCollection, error) {
	socialModel := model.NewSocial()
	res, err := socialModel.ListSocial(param.UserName)
	if err != nil {
		err := e.NewBusinessError(e.MONGOERROR, "mongo error")
		return nil, err
	}
	data := resources.NewSocialCollection()
	data.UserName = param.UserName
	data.Socials = res
	return data, nil
}
