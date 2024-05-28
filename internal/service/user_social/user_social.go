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

func (s *UserAddSocialController) AddSocial(param *form.UserSocial) error {
	userData, err := s.listSocial(&form.UserName{UserName: param.UserName})
	if err != nil {
		err := e.NewBusinessError(e.SocialListError, "social list error")
		return err
	}
	for _, v := range userData.Socials {
		if v.SocialName == param.SocialName {
			err := e.NewBusinessError(e.SocialAllReadyExist, "social name is exist")
			return err
		}
	}
	socialModel := model.NewSocial()
	err = socialModel.AddSocial(param)
	if err != nil {
		err := e.NewBusinessError(e.SocialAddError, "social add error")
		return err
	}
	s.updateSocialCache(s.argumentAddData(param.UserName, userData.Socials, param))
	// update to redis
	return nil
}

func (s *UserAddSocialController) UpdateSocial(param *form.UserSocial) error {
	socialModel := model.NewSocial()
	userData, err := s.listSocial(&form.UserName{UserName: param.UserName})
	if err != nil {
		err := e.NewBusinessError(e.SocialListError, "social list error")
		return err
	}
	for i, v := range userData.Socials {
		if v.SocialName == param.SocialName {
			err = socialModel.UpdateSocial(param)
			if err != nil {
				err := e.NewBusinessError(e.SocialUpdateError, "social update error")
				return err
			}
			userData.Socials[i].SocialLink = param.SocialLink
			s.updateSocialCache(userData)
			return nil
		}
	}
	err = e.NewBusinessError(e.SocialNameNotExist, "social name is not exist")
	return err
}

func (s *UserAddSocialController) ListSocial(param *form.UserName) (*resources.SocialCollection, error) {
	cache, ok := s.listSocialCache(param.UserName)
	if ok {
		return cache, nil
	}
	// first get from redis
	data, err := s.listSocial(param)
	if err != nil {
		err := e.NewBusinessError(e.SocialListError, "mongo error")
		return nil, err
	}
	return data, nil
}

// tools

func (s *UserAddSocialController) listSocial(param *form.UserName) (*resources.SocialCollection, error) {
	socialModel := model.NewSocial()
	res, err := socialModel.ListSocial(param.UserName)
	if err != nil {
		return nil, err
	}
	var data = resources.NewSocialCollection()
	data.UserName = param.UserName
	data.Socials = res
	return data, nil
}

// argument data
func (s *UserAddSocialController) argumentAddData(userName string, data []*model.SocialModel, newSocial *form.UserSocial) *resources.SocialCollection {
	newData := resources.NewSocialCollection()
	newData.UserName = userName
	newData.Socials = append(data, &model.SocialModel{
		SocialName: newSocial.SocialName,
		SocialLink: newSocial.SocialLink,
	})
	return newData
}
