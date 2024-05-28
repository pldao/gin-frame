package middleware

import (
	"github.com/PLDao/gin-frame/internal/pkg/aes_crypto"
	e "github.com/PLDao/gin-frame/internal/pkg/errors"
	"github.com/PLDao/gin-frame/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

// DecryptMiddleware 解密中间件
func AecMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.URL.RawQuery
		if uri == "" {
			c.Next()
			return
		}

		req, err := aes_crypto.DePwdCode(uri)
		if err != nil {
			response.Fail(c, e.AecFail, "aec 解密失败！")
			c.Abort()
			return
		}

		// Update the request with the decrypted query
		c.Request.URL.RawQuery = req

		// Continue processing the request
		c.Next()
	}
}
