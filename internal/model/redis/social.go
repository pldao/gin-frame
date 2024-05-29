package redis

import (
	"context"
	"encoding/json"
	"github.com/PLDao/gin-frame/data"
	model "github.com/PLDao/gin-frame/internal/model/mongo"
	log "github.com/PLDao/gin-frame/internal/pkg/logger"
	"github.com/PLDao/gin-frame/internal/resources"
	"github.com/go-redis/redis/v8"
)

type SocialCache struct {
	UserName    string
	SocialCache []*model.SocialModel
	*redis.Client
}

func NewSocialCache() *SocialCache {
	return &SocialCache{
		UserName:    "",
		SocialCache: nil,
		Client:      data.Rdb,
	}
}

func (s *SocialCache) GetUserSocialListCache(userName string) (*resources.SocialCollection, bool, error) {
	cachedData, err := s.Get(context.Background(), userName).Result()
	if err != nil {
		return nil, false, err
	}
	var social *resources.SocialCollection
	err = json.Unmarshal([]byte(cachedData), &social)
	if err != nil {
		log.Logger.Sugar().Error(err)
	}
	return social, true, nil
}

func (s *SocialCache) UpdateSocialListCache(data *resources.SocialCollection) error {
	// 先尝试删除旧的缓存项
	if err := s.Del(context.Background(), data.UserName).Err(); err != nil && err != redis.Nil {
		return err
	}
	// 然后添加新的数据
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = s.Set(context.Background(), data.UserName, string(jsonData), 3600).Err()
	return err
}
