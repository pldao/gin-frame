package middleware

import (
	"github.com/PLDao/gin-frame/internal/controller"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type RateLimiter struct {
	controller.Api
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{}
}

type RequestInfo struct {
	LastAccessTime time.Time // 上次访问时间
	RequestNum     int       // 请求计数
}

var (
	requestInfoMap = make(map[string]*RequestInfo) // IP到请求信息的映射
	mutex          = &sync.Mutex{}                 // 用于保护requestInfoMap的互斥锁
	maxRequests    = 5                             // 允许的最大请求数
	timeWindow     = 1 * time.Second               // 时间窗口
)

// IP限流器
func (l RateLimiter) IpLimit(c *gin.Context) {
	ip := c.ClientIP()
	mutex.Lock()
	defer mutex.Unlock()
	// 检查IP是否在map中
	info, exists := requestInfoMap[ip]
	// 如果IP不存在，初始化并添加到map中
	if !exists {
		requestInfoMap[ip] = &RequestInfo{LastAccessTime: time.Now(), RequestNum: 1}
		return
	}
	// 如果IP存在，检查时间窗口
	if time.Since(info.LastAccessTime) > timeWindow {
		// 如果超过时间窗口，重置请求计数
		info.RequestNum = 1
		info.LastAccessTime = time.Now()
		return
	}
	info.RequestNum++ // 如果在时间窗口内，增加请求计数
	// 如果请求计数超过限制，禁止访问
	if info.RequestNum > maxRequests {
		l.Fail(c, http.StatusTooManyRequests, "请求过于频繁，请稍后再试！")
		c.Abort()
		return
	}
	// 更新最后访问时间
	info.LastAccessTime = time.Now()
	c.Next()
}

// ratelimit限流器
func (l RateLimiter) RateLimit(time time.Duration, originNum, pushNum int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(time, originNum, pushNum)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			l.Fail(c, http.StatusTooManyRequests, "请求过于频繁，请稍后再试！")
			c.Abort()
			return
		}
		c.Next()
	}
}
