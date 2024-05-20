package middleware

import (
	"fmt"
	"github.com/PLDao/gin-frame/config"
	"github.com/PLDao/gin-frame/internal/global"
	e "github.com/PLDao/gin-frame/internal/pkg/errors"
	"github.com/PLDao/gin-frame/internal/pkg/response"
	"github.com/PLDao/gin-frame/internal/pkg/utils/token"
	"github.com/PLDao/gin-frame/internal/service/admin_auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

func AdminAuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		accessToken, err := token.GetAccessToken(authorization)
		if err != nil {
			response.Fail(c, e.NotLogin, err.Error())
			return
		}
		adminCustomClaims := new(token.AdminCustomClaims)
		// 解析Token
		err = token.Parse(accessToken, adminCustomClaims, jwt.WithSubject(global.Subject))
		if err != nil || adminCustomClaims == nil {
			response.FailCode(c, e.NotLogin)
			return
		}

		exp, err := adminCustomClaims.GetExpirationTime()
		// 获取过期时间返回err,或者exp为nil都返回错误
		if err != nil || exp == nil {
			response.FailCode(c, e.NotLogin)
			return
		}

		// 刷新时间大于0则判断剩余时间小于刷新时间时刷新Token并在Response header中返回
		if config.Config.Jwt.RefreshTTL > 0 {
			now := time.Now()
			diff := exp.Time.Sub(now)
			refreshTTL := config.Config.Jwt.RefreshTTL * time.Second
			fmt.Println(diff.Seconds(), refreshTTL)
			if diff < refreshTTL {
				tokenResponse, _ := admin_auth.NewLoginService().Refresh(adminCustomClaims.UserID)
				c.Writer.Header().Set("refresh-access-token", tokenResponse.AccessToken)
				c.Writer.Header().Set("refresh-exp", strconv.FormatInt(tokenResponse.ExpiresAt, 10))
			}
		}

		c.Set("a_uid", adminCustomClaims.UserID)
		c.Set("a_mobile", adminCustomClaims.Mobile)
		c.Set("a_nickname", adminCustomClaims.Nickname)
		c.Set("admin_user_info", adminCustomClaims.AdminUserInfo)
		c.Next()
	}
}
