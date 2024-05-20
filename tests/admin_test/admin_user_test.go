package admin_test

import (
	"fmt"
	e "github.com/PLDao/gin-frame/internal/pkg/errors"
	"github.com/PLDao/gin-frame/internal/pkg/response"
	"github.com/PLDao/gin-frame/pkg/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	route := fmt.Sprintf("%s/api/v1/admin/login", ts.URL)
	h := utils.HttpRequest{}

	body := `{"username": "super_admin","password": "123456"}`
	resp := h.JsonRequest("POST", route, strings.NewReader(body))

	assert.Nil(t, resp.Error)
	assert.Equal(t, http.StatusOK, resp.Response.StatusCode)
	result := new(response.Result)
	err := resp.ParseJson(result)
	assert.Nil(t, err)
	assert.Equal(t, e.SUCCESS, result.Code)
}

func TestGetAdminUser(t *testing.T) {
	route := fmt.Sprintf("%s/api/v1/admin-user/get", ts.URL)
	queryParams := &url.Values{}
	queryParams.Set("id", "1")
	resp := getRequest(route, queryParams)

	assert.Nil(t, resp.Error)
	assert.Equal(t, http.StatusOK, resp.Response.StatusCode)
	result := new(response.Result)
	err := resp.ParseJson(result)
	assert.Nil(t, err)
	assert.Equal(t, e.SUCCESS, result.Code)
}
