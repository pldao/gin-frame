package data

import (
	c "github.com/PLDao/gin-frame/config"
	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"sync"
)

var (
	once        sync.Once
	MysqlDB     *gorm.DB
	Rdb         *redis.Client
	MongoDB     *mongo.Client
	MinioClient *minio.Client
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

		if c.Config.Mongo.Enable {
			// 初始化 mongo
			initMongoDB()
		}

		if c.Config.Minio.Enable {
			// 初始化 minio
			initMinIO()
		}
	})
}
