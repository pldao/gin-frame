package external

import (
	"encoding/json"
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
		SetHeader("Ok-Access-Key", apiKey).
		SetHeader("X-API-KEY", apiKey)
	//Authorization

	return &Client{
		RestyClient: client,
		APIKey:      apiKey,
		BaseURL:     baseURL,
	}
}

// SendRequest 发送 HTTP 请求并返回响应
func (c *Client) SendRequest(method, endpoint string, payload interface{}) ([]byte, error) {
	url := c.BaseURL + endpoint

	var resp *resty.Response
	var err error

	switch method {
	case "GET":
		request := c.RestyClient.R()
		if payload != nil {
			request.SetQueryParams(payload.(map[string]string))
		}
		resp, err = request.Get(url)
	case "POST":
		request := c.RestyClient.R()
		if payload != nil {
			request.SetBody(payload)
		}
		resp, err = request.Post(url)
	case "PUT":
		request := c.RestyClient.R()
		if payload != nil {
			request.SetBody(payload)
		}
		resp, err = request.Put(url)
	case "DELETE":
		request := c.RestyClient.R()
		if payload != nil {
			request.SetBody(payload)
		}
		resp, err = request.Delete(url)
	default:
		return nil, fmt.Errorf("unsupported HTTP method: %s", method)
	}

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("request failed with status: %s", resp.Status())
	}
	fmt.Println("Response Body:", resp.String())

	return resp.Body(), nil
}

func (c *Client) SendRequestAndParseJSON(method, endpoint string, payload interface{}, result interface{}) error {
	body, err := c.SendRequest(method, endpoint, payload)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		return fmt.Errorf("failed to parse JSON: %v", err)
	}

	return nil
}
