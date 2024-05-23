package autoload

import (
	"time"
)

type MongoConfig struct {
	Enable          bool
	Host            string
	Port            uint16
	Username        string
	Password        string
	Database        string
	MaxPoolSize     uint64
	MinPoolSize     uint64
	MaxConnIdleTime time.Duration
}

var Mongo = MongoConfig{
	Enable:          false,
	Host:            "127.0.0.1",
	Port:            27017,
	Username:        "admin",
	Password:        "adminpassword",
	Database:        "testdb",
	MaxPoolSize:     100,              // 连接池最大数量
	MinPoolSize:     10,               // 连接池最小数量
	MaxConnIdleTime: 30 * time.Minute, // 连接池中连接的最大空闲时间
}
