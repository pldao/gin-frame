package external

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

// Client 结构体，包含 Resty 客户端和基础 URL
type Client struct {
	RestyClient *resty.Client
	APIKey      string
	BaseURL     string
}

// NewClient 返回一个新的 Client 实例
func NewClient(baseURL, apiKey string) *Client {
	client := resty.New()
	client.SetTimeout(10*time.Second).
		SetRetryCount(3).
		SetRetryWaitTime(5*time.Second).
		SetHeader("Content-Type", "application/json").
		SetHeader("Ok-Access-Key", apiKey)

	return &Client{
		RestyClient: client,
		APIKey:      apiKey,
		BaseURL:     baseURL,
	}
}

// SendRequest 发送 HTTP 请求并返回响应
func (c *Client) SendRequest(method, endpoint string, payload interface{}) (string, error) {
	url := c.BaseURL + endpoint

	var resp *resty.Response
	var err error

	switch method {
	case "GET":
		resp, err = c.RestyClient.R().
			SetQueryParams(payload.(map[string]string)).
			Get(url)
	case "POST":
		resp, err = c.RestyClient.R().
			SetBody(payload).
			Post(url)
	case "PUT":
		resp, err = c.RestyClient.R().
			SetBody(payload).
			Put(url)
	case "DELETE":
		resp, err = c.RestyClient.R().
			SetBody(payload).
			Delete(url)
	default:
		return "", fmt.Errorf("unsupported HTTP method: %s", method)
	}

	if err != nil {
		return "", err
	}

	if resp.IsError() {
		return "", fmt.Errorf("request failed with status: %s", resp.Status())
	}

	return resp.String(), nil
}
