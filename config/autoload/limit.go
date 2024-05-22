package autoload

import "time"

type LimitConfig struct {
	MaxRequests int           `mapstructure:"max_requests"` // 允许的最大请求数
	TimeWindow  time.Duration `mapstructure:"time_window"`  // 时间窗口
}

var Limit = LimitConfig{
	MaxRequests: 5,
	TimeWindow:  1 * time.Second,
}
