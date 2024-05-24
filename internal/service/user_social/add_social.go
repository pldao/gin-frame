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
	userData, err := socialModel.ListSocial(param.UserName)
	if err != nil {
		err := e.NewBusinessError(e.SocialListError, "social list error")
		return err
	}
	for _, v := range userData {
		if v.SocialName == param.SocialName {
			err := e.NewBusinessError(e.SocialAllReadyExist, "social name is exist")
			return err
		}
	}

	err = socialModel.AddSocial(param)
	if err != nil {
		err := e.NewBusinessError(e.SocialAddError, "social add error")
		return err
	}
	//newData := append(userData, &model.SocialModel{
	//	SocialLink: param.SocialLink,
	//	SocialName: param.SocialName,
	//})
	// update to redis
	return nil
}

func (s UserAddSocialController) ListSocial(param *form.UserName) (*resources.SocialCollection, error) {
	// first get from redis
	socialModel := model.NewSocial()
	res, err := socialModel.ListSocial(param.UserName)
	if err != nil {
		err := e.NewBusinessError(e.SocialListError, "mongo error")
		return nil, err
	}
	data := resources.NewSocialCollection()
	data.UserName = param.UserName
	data.Socials = res
	return data, nil
}

func (s UserAddSocialController) UpdateSocial(param *form.UserSocial) error {
	socialModel := model.NewSocial()
	userData, err := socialModel.ListSocial(param.UserName)
	if err != nil {
		err := e.NewBusinessError(e.SocialListError, "social list error")
		return err
	}
	for _, v := range userData {
		if v.SocialName == param.SocialName {
			err = socialModel.UpdateSocial(param)
			if err != nil {
				err := e.NewBusinessError(e.SocialUpdateError, "social update error")
				return err
			}
			//newData := append(userData, &model.SocialModel{
			//	SocialLink: param.SocialLink,
			//	SocialName: param.SocialName,
			//})
			// update to redis
			return nil
		}
	}
	err = e.NewBusinessError(e.SocialNameNotExist, "social name is not exist")
	return err
}
