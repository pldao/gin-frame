package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PLDao/gin-frame/data"
	model "github.com/PLDao/gin-frame/internal/model/mongo"
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

func (s *SocialCache) GetUserSocialListCache(userName string) ([]*model.SocialModel, bool, error) {
	cachedData, err := s.Get(context.Background(), userName).Result()
	if err != nil {
		return nil, false, nil
	}
	var social []*model.SocialModel
	err = json.Unmarshal([]byte(cachedData), &social)
	if err != nil {
		fmt.Println("Failed to unmarshal cached data:", err)
	}
	return social, true, nil
}

func (s *SocialCache) UpdateSocialListCache(data *resources.SocialCollection) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = s.Set(context.Background(), data.UserName, string(jsonData), 0).Err()
	if err != nil {
		return err
	}
	return nil
}
