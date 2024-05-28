package middleware

import (
	"errors"
	e "github.com/PLDao/gin-frame/internal/pkg/errors"
	"github.com/PLDao/gin-frame/internal/pkg/response"
	"time"

	"github.com/gin-gonic/gin"
)

// TimestampMiddleware 中间件函数，用于验证时间戳
func TimestampMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取时间戳
		timestamp := c.GetHeader("X-Timestamp")
		// 验证时间戳
		if err := validateTimestamp(timestamp); err != nil {
			// 如果时间戳无效，则返回错误响应
			response.Fail(c, e.NotTimeout, err.Error())
			c.Abort()
			return
		}
		// 继续处理请求
		c.Next()
	}
}

// validateTimestamp 验证时间戳
func validateTimestamp(timestamp string) error {
	// 检查时间戳是否为空
	if timestamp == "" {
		return errors.New("timestamp is empty")
	}
	// 解析时间戳
	requestTime, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return errors.New("invalid timestamp format")
	}
	// 检查时间戳是否在允许的时间范围内（例如 5 分钟内）
	if time.Since(requestTime) > 5*time.Minute {
		return errors.New("timestamp is expired")
	}
	return nil
}
