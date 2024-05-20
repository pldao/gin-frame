package admin_test

import (
	"fmt"
	e "github.com/PLDao/gin-frame/internal/pkg/errors"
	"github.com/PLDao/gin-frame/internal/pkg/response"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestPermissionEdit(t *testing.T) {
	route := fmt.Sprintf("%s/api/v1/permission/edit", ts.URL)

	body := `{"id":6,"name":"ping","desc":"","method":"GET","route":"/ping","func":"func1","func_path":"github.com/PLDao/gin-frame/internal/routers.SetRouters.func1","is_auth":2,"sort":100}`
	resp := postRequest(route, &body)

	assert.Nil(t, resp.Error)
	assert.Equal(t, http.StatusOK, resp.Response.StatusCode)
	result := new(response.Result)
	err := resp.ParseJson(result)
	assert.Nil(t, err)
	assert.Equal(t, e.SUCCESS, result.Code)
}

func TestPermissionList(t *testing.T) {
	route := fmt.Sprintf("%s/api/v1/permission/list", ts.URL)
	queryParams := &url.Values{}
	queryParams.Set("page", "1")
	queryParams.Set("per_page", "1")
	queryParams.Set("name", "ping")
	queryParams.Set("method", "GET")
	queryParams.Set("route", "/ping")
	queryParams.Set("is_auth", "2")
	resp := getRequest(route, queryParams)

	assert.Nil(t, resp.Error)
	assert.Equal(t, http.StatusOK, resp.Response.StatusCode)
	result := new(response.Result)
	err := resp.ParseJson(result)
	fmt.Println(result, err)
	assert.Nil(t, err)
	assert.Equal(t, e.SUCCESS, result.Code)
}
