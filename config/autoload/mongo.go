package autoload

import (
	"time"
)

type MongoConfig struct {
	Enable          bool          `mapstructure:"enable"`
	Host            string        `mapstructure:"host"`
	Port            uint16        `mapstructure:"port"`
	Username        string        `mapstructure:"username"`
	Password        string        `mapstructure:"password"`
	Database        string        `mapstructure:"database"`
	MaxPoolSize     uint64        `mapstructure:"max_pool_size"`
	MinPoolSize     uint64        `mapstructure:"min_pool_size"`
	MaxConnIdleTime time.Duration `mapstructure:"max_conn_idle_time"`
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
