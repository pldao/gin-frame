package resources

import model "github.com/PLDao/gin-frame/internal/model/mongo"

type SocialCollection struct {
	UserName string               `json:"username"`
	Socials  []*model.SocialModel `json:"socials"`
}

func NewSocialCollection() *SocialCollection {
	return &SocialCollection{}
}
