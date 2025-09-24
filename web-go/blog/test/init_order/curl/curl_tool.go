package curl

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/cenkalti/backoff/v4"
)

// Client 自定义HTTP客户端
type Client struct {
	client      *http.Client    // 底层HTTP客户端
	baseURL     string          // 基础URL（如 "https://api.example.com"）
	headers     http.Header     // 全局请求头
	debug       bool            // 调试模式
	retry       int             // 重试次数
	backoff     backoff.BackOff // 退避策略
	middlewares []Middleware    // 拦截器
}

// Middleware 拦截器接口
type Middleware func(req *http.Request, resp *http.Response, err error) error

// Option 配置选项
type Option func(*Client)

// NewClient 创建客户端
func NewClient(opts ...Option) *Client {
	// 默认配置
	transport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     30 * time.Second,
	}

	c := &Client{
		client: &http.Client{
			Transport: transport,
			Timeout:   30 * time.Second, // 默认总超时
		},
		headers:     make(http.Header),
		retry:       0,
		backoff:     backoff.NewConstantBackOff(1 * time.Second),
		middlewares: []Middleware{},
	}

	// 应用选项
	for _, opt := range opts {
		opt(c)
	}

	return c
}

// 配置选项：设置基础URL
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// 配置选项：设置超时
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.client.Timeout = timeout
	}
}

// 设置全局请求头
func (c *Client) SetHeader(key, value string) {
	c.headers.Set(key, value)
}

// 发送GET请求
func (c *Client) Get(ctx context.Context, path string, opts ...RequestOption) (*Response, error) {
	return c.Do(ctx, http.MethodGet, path, nil, opts...)
}

// 发送POST请求（JSON body）
func (c *Client) PostJSON(ctx context.Context, path string, body interface{}, opts ...RequestOption) (*Response, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return c.Do(ctx, http.MethodPost, path, bytes.NewReader(data), append(opts, WithHeader("Content-Type", "application/json"))...)
}

// 核心请求方法
func (c *Client) Do(ctx context.Context, method, path string, body io.Reader, opts ...RequestOption) (*Response, error) {
	// 构建URL
	reqURL, err := c.buildURL(path)
	if err != nil {
		return nil, err
	}

	// 创建请求
	req, err := http.NewRequestWithContext(ctx, method, reqURL, body)
	if err != nil {
		return nil, err
	}

	// 设置全局头
	req.Header = c.headers.Clone()

	// 应用请求级选项（如查询参数、临时头）
	reqOpts := &requestOptions{}
	for _, opt := range opts {
		opt(reqOpts)
	}
	// 设置查询参数
	q := req.URL.Query()
	for k, v := range reqOpts.query {
		q.Set(k, v)
	}
	req.URL.RawQuery = q.Encode()
	// 设置请求头
	for k, v := range reqOpts.headers {
		req.Header.Set(k, v)
	}

	// 执行请求（带重试和拦截器）
	var resp *http.Response
	var respErr error

	// 重试逻辑
	operation := func() error {
		// 执行前置拦截器
		for _, m := range c.middlewares {
			if err := m(req, nil, nil); err != nil {
				return err
			}
		}

		// 发送请求
		resp, respErr = c.client.Do(req)

		// 执行后置拦截器
		for _, m := range c.middlewares {
			if err := m(req, resp, respErr); err != nil {
				return err
			}
		}

		// 判断是否需要重试
		if shouldRetry(resp, respErr) {
			return respErr
		}
		return nil
	}

	// 带重试执行
	if c.retry > 0 {
		bo := backoff.WithMaxRetries(c.backoff, uint64(c.retry))
		if err := backoff.Retry(operation, bo); err != nil {
			return nil, err
		}
	} else {
		if err := operation(); err != nil {
			return nil, err
		}
	}

	return &Response{resp}, nil
}

// 辅助方法：构建完整URL
func (c *Client) buildURL(path string) (string, error) {
	if c.baseURL == "" {
		return path, nil
	}
	base, err := url.Parse(c.baseURL)
	if err != nil {
		return "", err
	}
	u, err := base.Parse(path)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

// Response 响应封装
type Response struct {
	*http.Response
}

// ParseJSON 解析JSON响应到结构体
func (r *Response) ParseJSON(v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

// 其他辅助方法：判断是否需要重试
func shouldRetry(resp *http.Response, err error) bool {
	// 网络错误重试
	if err != nil {
		return true
	}
	// 5xx状态码重试
	if resp.StatusCode >= 500 && resp.StatusCode < 600 {
		return true
	}
	return false
}

// 请求级选项（如查询参数、临时头）
type requestOptions struct {
	query   map[string]string
	headers map[string]string
}

type RequestOption func(*requestOptions)

func WithQuery(key, value string) RequestOption {
	return func(o *requestOptions) {
		if o.query == nil {
			o.query = make(map[string]string)
		}
		o.query[key] = value
	}
}

func WithHeader(key, value string) RequestOption {
	return func(o *requestOptions) {
		if o.headers == nil {
			o.headers = make(map[string]string)
		}
		o.headers[key] = value
	}
}
