package user_social

import (
	model "github.com/PLDao/gin-frame/internal/model/redis"
	log "github.com/PLDao/gin-frame/internal/pkg/logger"
	"github.com/PLDao/gin-frame/internal/resources"
)

func (s *UserAddSocialController) listSocialCache(userName string) (*resources.SocialCollection, bool) {
	// first get from redis
	socialModel := model.NewSocialCache()
	socialList, ok, _ := socialModel.GetUserSocialListCache(userName)
	if ok {
		return socialList, true
	}
	return nil, false
}

func (s *UserAddSocialController) updateSocialCache(data *resources.SocialCollection) {
	socialModel := model.NewSocialCache()
	err := socialModel.UpdateSocialListCache(data)
	if err != nil {
		log.Logger.Sugar().Error(err)
	}
}
