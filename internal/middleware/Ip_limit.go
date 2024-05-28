package middleware

import (
	"github.com/PLDao/gin-frame/config"
	"github.com/PLDao/gin-frame/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"sync"
	"time"
)

var (
	requestInfoMap = make(map[string]*RequestInfo)
	mutex          = &sync.Mutex{}
)

type RequestInfo struct {
	LastAccessTime time.Time
	RequestNum     int
}

// IpLimit IP限流器中间件
func IpLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		mutex.Lock()
		defer mutex.Unlock()

		info, exists := requestInfoMap[ip]
		if !exists {
			requestInfoMap[ip] = &RequestInfo{LastAccessTime: time.Now(), RequestNum: 1}
			c.Next()
			return
		}

		if time.Since(info.LastAccessTime) > config.Config.Limit.TimeWindow {
			info.RequestNum = 1
			info.LastAccessTime = time.Now()
			c.Next()
			return
		}

		info.RequestNum++
		if info.RequestNum > config.Config.Limit.MaxRequests {
			response.Fail(c, http.StatusTooManyRequests, "请求过于频繁，请稍后再试！")
			c.Abort()
			return
		}

		info.LastAccessTime = time.Now()
		c.Next()
	}
}

// RateLimit 基于令牌桶的限流器中间件
func RateLimit(interval time.Duration, capacity, quantum int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(interval, capacity, quantum)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			response.Fail(c, http.StatusTooManyRequests, "请求过于频繁，请稍后再试！")
			c.Abort()
			return
		}
		c.Next()
	}
}
