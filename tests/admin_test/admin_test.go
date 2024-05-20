package admin_test

import (
	"fmt"
	c "github.com/PLDao/gin-frame/config"
	"github.com/PLDao/gin-frame/internal/global"
	"github.com/PLDao/gin-frame/internal/pkg/utils/token"
	"github.com/PLDao/gin-frame/pkg/utils"
	"github.com/PLDao/gin-frame/tests"
	"github.com/golang-jwt/jwt/v5"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

var (
	ts            *httptest.Server
	authorization string
)

func TestMain(m *testing.M) {
	ts = httptest.NewServer(tests.SetupRouter())
	now := time.Now()
	expiresAt := now.Add(time.Second * c.Config.Jwt.TTL)
	claims := token.AdminCustomClaims{
		AdminUserInfo: token.AdminUserInfo{
			UserID:   1,
			Mobile:   "13200000000",
			Nickname: "admin",
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			Issuer:    global.Issuer, // 签发人
			//IssuedAt:  jwt.NewNumericDate(now),       // 签发时间
			Subject: global.Subject, // 签发主体
			//NotBefore: jwt.NewNumericDate(now),       // 生效时间
		},
	}
	accessToken, err := token.Generate(claims)
	authorization = fmt.Sprintf("%s%s", "Bearer ", accessToken)
	if err != nil {
		panic("创建管理员Token失败")
	}
	m.Run()
}

func postRequest(route string, body *string) *utils.HttpRequest {
	options := map[string]string{
		"Authorization": authorization,
	}
	return tests.Request("POST", route, body, options)
}

func getRequest(route string, queryParams *url.Values) *utils.HttpRequest {
	options := map[string]string{
		"Authorization": authorization,
	}
	return tests.GetRequest(route, queryParams, options)
}
