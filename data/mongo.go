package data

import (
	"context"
	"fmt"
	c "github.com/PLDao/gin-frame/config"
	"github.com/PLDao/gin-frame/config/autoload"
	log "github.com/PLDao/gin-frame/internal/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// initMongoDB initializes the MongoDB client and connects to the database
func initMongoDB() {
	// 构建 MongoDB 连接 URI
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d",
		c.Config.Mongo.Username,
		c.Config.Mongo.Password,
		c.Config.Mongo.Host,
		c.Config.Mongo.Port,
	)
	// 设置客户端选项
	clientOptions := options.Client().ApplyURI(uri).
		SetMaxPoolSize(autoload.Mongo.MaxPoolSize).
		SetMinPoolSize(autoload.Mongo.MinPoolSize).
		SetMaxConnIdleTime(autoload.Mongo.MaxConnIdleTime)

	// 创建新的 MongoDB 客户端
	var err error
	MongoDB, err = mongo.NewClient(clientOptions)
	if err != nil {
		panic("Failed to create MongoDB client: " + err.Error())
	}

	// 连接到 MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = MongoDB.Connect(ctx)
	if err != nil {
		panic("Failed to connect to MongoDB: " + err.Error())
	}

	// Ping MongoDB 来验证连接
	err = MongoDB.Ping(ctx, nil)
	if err != nil {
		panic("Failed to ping MongoDB: " + err.Error())
	}

	log.Logger.Sugar().Info("MongoDB connected successfully")
}
