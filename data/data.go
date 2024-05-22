package data

import (
	c "github.com/PLDao/gin-frame/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"sync"
)

var (
	once    sync.Once
	MysqlDB *gorm.DB
	Rdb     *redis.Client
)

func InitData() {
	once.Do(func() {
		if c.Config.Mysql.Enable {
			// 初始化 mysql
			initMysql()
		}

		if c.Config.Redis.Enable {
			// 初始化 redis
			initRedis()
		}
	})
}
