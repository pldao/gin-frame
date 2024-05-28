package middleware

import (
	"errors"
	"fmt"
	"github.com/PLDao/gin-frame/config"
	e "github.com/PLDao/gin-frame/internal/pkg/errors"
	"github.com/PLDao/gin-frame/internal/pkg/logger"
	"github.com/PLDao/gin-frame/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// CustomRecovery 自定义错误 (panic) 拦截中间件、对可能发生的错误进行拦截、统一记录
func CustomRecovery() gin.HandlerFunc {
	DefaultErrorWriter := &PanicExceptionRecord{}
	return gin.RecoveryWithWriter(DefaultErrorWriter, func(c *gin.Context, err interface{}) {
		// 这里针对发生的panic等异常进行统一响应即可
		errStr := ""
		if config.Config.Debug {
			errStr = fmt.Sprintf("%v", err)
		}
		response.Resp().SetHttpCode(http.StatusInternalServerError).FailCode(c, e.ServerError, errStr)
	})
}

// PanicExceptionRecord  panic等异常记录
type PanicExceptionRecord struct{}

func (p *PanicExceptionRecord) Write(b []byte) (n int, err error) {
	s1 := "An error occurred in the server's internal code："
	var build strings.Builder
	build.WriteString(s1)
	build.Write(b)
	errStr := build.String()
	logger.Logger.Error(errStr)
	return len(errStr), errors.New(errStr)
}
