package curl

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	logger "github.com/learn/init_order/log"
)

// APITester 定义API测试器结构体，用于链式调用API
type APITester struct {
	client   *resty.Client
	token    string
	username string
	password string
	email    string
	err      error
}

// NewAPITester 创建一个新的API测试器实例
func NewAPITester(username, password, email string) *APITester {
	client := resty.New()
	client.SetTimeout(10 * time.Second)
	return &APITester{
		client:   client,
		username: username,
		password: password,
		email:    email,
	}
}

// BatchUser 定义批量注册中的用户结构体
type BatchUser struct {
	Username string
	Password string
	Email    string
}

// Register 注册用户
func (t *APITester) Register() *APITester {
	if t.err != nil {
		return t
	}

	logger.Info("=== 步骤1: 注册用户 ===")
	resp, err := t.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"username": t.username,
			"password": t.password,
			"email":    t.email,
		}).
		SetResult(map[string]interface{}{}).
		Post("http://localhost:8080/api/register")

	if err != nil {
		logger.Errorf("用户注册失败: %v", err)
		logger.Info("将跳过注册步骤，继续执行登录流程...")
	} else {
		logger.Infof("注册响应状态码: %d", resp.StatusCode())
		if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
			result := resp.Result().(*map[string]interface{})
			demoJSONOutput(*result)
		} else {
			logger.Errorf("注册失败，响应内容: %s", resp.String())
			logger.Info("用户名可能已被占用，将跳过注册步骤，继续执行登录流程...")
		}
	}

	return t
}

// BatchRegister 批量注册多个用户（并发）
func (t *APITester) BatchRegister(users []BatchUser) *APITester {
	if t.err != nil {
		return t
	}

	logger.Info("=== 批量注册用户 ===")
	logger.Infof("开始批量注册 %d 个用户...", len(users))

	// 记录开始时间
	startTime := time.Now()

	// 并发发送请求
	var wg sync.WaitGroup
	for i, user := range users {
		wg.Add(1)
		go func(index int, u BatchUser) {
			defer wg.Done()

			// 发送POST请求
			resp, err := t.client.R().
				SetHeader("Content-Type", "application/json").
				SetBody(map[string]interface{}{
					"username": u.Username,
					"password": u.Password,
					"email":    u.Email,
				}).
				SetResult(map[string]interface{}{}).
				Post("http://localhost:8080/api/register")

			if err != nil {
				logger.Errorf("发送用户%s注册请求失败: %v", u.Username, err)
				return
			}

			// 检查响应状态码
			if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
				logger.Errorf("用户%s注册失败，状态码: %d", u.Username, resp.StatusCode())
				logger.Errorf("响应内容: \n%s", resp.String())
				return
			}

			// 解析响应结果
			resultPtr := resp.Result().(*map[string]interface{})
			data := *resultPtr

			// 输出成功信息
			logger.Infof("用户%s注册成功", u.Username)
			// 输出格式化的JSON响应
			demoJSONOutput(data)
		}(i, user)
	}

	// 等待所有请求完成
	wg.Wait()

	// 计算总耗时
	totalTime := time.Since(startTime)
	logger.Info("所有注册请求已发送完成")
	logger.Infof("所有请求完成共花费时间: %.2f ms", float64(totalTime.Milliseconds()))

	return t
}

// Login 登录获取token
func (t *APITester) Login() *APITester {
	if t.err != nil {
		return t
	}

	logger.Info("\n=== 步骤2: 登录获取token ===")
	resp, err := t.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"username": t.username,
			"password": t.password,
		}).
		SetResult(map[string]interface{}{}).
		Post("http://localhost:8080/api/login")

	if err != nil {
		logger.Errorf("用户登录失败: %v", err)
		t.err = err
		return t
	}

	logger.Infof("登录响应状态码: %d", resp.StatusCode())
	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		result := resp.Result().(*map[string]interface{})
		demoJSONOutput(*result)

		// 从响应中提取token
		if tokenData, ok := (*result)["data"].(map[string]interface{}); ok {
			if tokenValue, ok := tokenData["token"].(string); ok {
				t.token = tokenValue
				logger.Infof("成功获取token: %s", t.token)
			}
		}
	} else {
		logger.Errorf("登录失败，响应内容: %s", resp.String())
		t.err = fmt.Errorf("login failed")
	}

	return t
}

// CreatePost 发布文章
func (t *APITester) CreatePost(title, content string) *APITester {
	if t.err != nil || t.token == "" {
		return t
	}

	logger.Info("\n=== 步骤3: 使用token发帖 ====")
	resp, err := t.client.R().
		SetHeader("token", t.token).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"title":   title,
			"content": content,
		}).
		SetResult(map[string]interface{}{}).
		Post("http://localhost:8080/api/posts")

	if err != nil {
		logger.Errorf("发布文章失败: %v", err)
		t.err = err
		return t
	}

	logger.Infof("发布文章响应状态码: %d", resp.StatusCode())
	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		result := resp.Result().(*map[string]interface{})
		demoJSONOutput(*result)
	} else {
		logger.Errorf("发布文章失败，响应内容: %s", resp.String())
		t.err = fmt.Errorf("create post failed")
	}

	return t
}

// UpdatePost 更新文章
func (t *APITester) UpdatePost(postID string, title, content string) *APITester {
	if t.err != nil || t.token == "" {
		return t
	}

	logger.Info("\n=== 步骤4: 更新文章 ===")
	resp, err := t.client.R().
		SetHeader("token", t.token).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"title":   title,
			"content": content,
		}).
		SetResult(map[string]interface{}{}).
		Put(fmt.Sprintf("http://localhost:8080/api/posts/%s", postID))

	if err != nil {
		logger.Errorf("更新文章失败: %v", err)
	} else {
		logger.Infof("更新文章响应状态码: %d", resp.StatusCode())
		if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
			result := resp.Result().(*map[string]interface{})
			demoJSONOutput(*result)
		} else {
			logger.Errorf("更新文章失败，响应内容: %s", resp.String())
		}
	}

	return t
}

// DeletePost 删除文章
func (t *APITester) DeletePost(postID string) *APITester {
	if t.err != nil || t.token == "" {
		return t
	}

	logger.Info("\n=== 步骤5: 删除文章 ===")
	resp, err := t.client.R().
		SetHeader("token", t.token).
		SetResult(map[string]interface{}{}).
		Delete(fmt.Sprintf("http://localhost:8080/api/posts/%s", postID))

	if err != nil {
		logger.Errorf("删除文章失败: %v", err)
	} else {
		logger.Infof("删除文章响应状态码: %d", resp.StatusCode())
		if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
			result := resp.Result().(*map[string]interface{})
			demoJSONOutput(*result)
		} else {
			logger.Errorf("删除文章失败，响应内容: %s", resp.String())
		}
	}

	return t
}

// AddComments 添加多条评论
func (t *APITester) AddComments(postID string, contents []string, intervalMs int) *APITester {
	if t.err != nil || t.token == "" {
		return t
	}

	logger.Info("\n=== 步骤6: 为文章新增评论 ===")
	for i, content := range contents {
		logger.Infof("\n--- 添加第%d条评论 ---", i+1)
		resp, err := t.client.R().
			SetHeader("token", t.token).
			SetHeader("Content-Type", "application/json").
			SetBody(map[string]interface{}{
				"content": content,
			}).
			SetResult(map[string]interface{}{}).
			Post(fmt.Sprintf("http://localhost:8080/api/posts/%s/comments", postID))

		if err != nil {
			logger.Errorf("添加评论失败: %v", err)
		} else {
			logger.Infof("添加评论响应状态码: %d", resp.StatusCode())
			if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
				result := resp.Result().(*map[string]interface{})
				demoJSONOutput(*result)
			} else {
				logger.Errorf("添加评论失败，响应内容: %s", resp.String())
			}
		}

		// 每条评论之间暂停指定时间
		time.Sleep(time.Duration(intervalMs) * time.Millisecond)
	}

	return t
}

// Post 定义文章结构体

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// BatchCreatePosts 批量创建多篇文章（并发）
func (t *APITester) BatchCreatePosts(posts []Post) *APITester {
	if t.err != nil || t.token == "" {
		return t
	}

	logger.Info("\n=== 批量创建文章 ====")
	logger.Infof("开始批量创建 %d 篇文章...", len(posts))

	// 记录开始时间
	startTime := time.Now()

	// 并发发送请求
	var wg sync.WaitGroup
	for i, post := range posts {
		wg.Add(1)
		go func(index int, p Post) {
			defer wg.Done()

			// 发送POST请求
			resp, err := t.client.R().
				SetHeader("token", t.token).
				SetHeader("Content-Type", "application/json").
				SetBody(map[string]interface{}{
					"title":   p.Title,
					"content": p.Content,
				}).
				SetResult(map[string]interface{}{}).
				Post("http://localhost:8080/api/posts")

			if err != nil {
				logger.Errorf("发送创建文章%d请求失败: %v", index+1, err)
				return
			}

			// 检查响应状态码
			if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
				logger.Errorf("创建文章%d失败，状态码: %d", index+1, resp.StatusCode())
				logger.Errorf("响应内容: \n%s", resp.String())
				return
			}

			// 解析响应结果
			resultPtr := resp.Result().(*map[string]interface{})
			data := *resultPtr

			// 输出成功信息
			logger.Infof("创建文章%d成功", index+1)
			// 输出格式化的JSON响应
			demoJSONOutput(data)
		}(i, post)
	}

	// 等待所有请求完成
	wg.Wait()

	// 计算总耗时
	totalTime := time.Since(startTime)
	logger.Info("所有创建文章请求已发送完成")
	logger.Infof("所有请求完成共花费时间: %.2f ms", float64(totalTime.Milliseconds()))

	return t
}

// LoginWithForm 使用application/x-www-form-urlencoded格式登录
func (t *APITester) LoginWithForm() *APITester {
	if t.err != nil {
		return t
	}

	logger.Info("\n=== 测试enctype方式(application/x-www-form-urlencoded)登录请求 ===")

	// 准备登录数据
	loginData := map[string]string{
		"username": t.username,
		"password": t.password,
	}

	logger.Infof("登录请求数据: %+v", loginData)

	// 发送POST请求 - 使用enctype方式
	resp, err := t.client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(loginData). // 使用SetFormData替代SetBody
		SetResult(map[string]interface{}{}).
		Post("http://localhost:8080/api/login")

	if err != nil {
		logger.Errorf("发送enctype登录请求失败: %v", err)
		t.err = err
		return t
	}

	// 检查响应状态码
	logger.Debugf("响应内容: \n%s", resp.String())
	logger.Infof("响应状态码: %d", resp.StatusCode())

	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		logger.Errorf("登录失败，状态码: %d", resp.StatusCode())
		t.err = fmt.Errorf("login failed")
		return t
	}

	// 解析响应结果
	resultPtr := resp.Result().(*map[string]interface{})
	data := *resultPtr

	// 从响应中提取token
	if tokenData, ok := data["data"].(map[string]interface{}); ok {
		if tokenValue, ok := tokenData["token"].(string); ok {
			t.token = tokenValue
			logger.Infof("成功获取token: %s", t.token)
		}
	}

	// 输出成功信息
	logger.Info("enctype方式登录成功")
	// 输出格式化的JSON响应
	demoJSONOutput(data)
	logger.Info("=== enctype登录测试完成 ====")

	return t
}

// GetPosts 获取文章列表
func (t *APITester) GetPosts(page, limit string) *APITester {
	if t.err != nil {
		return t
	}

	logger.Info("\n=== 获取文章列表 ====")

	// 发送GET请求
	resp, err := t.client.R().
		SetHeader("User-Agent", "resty-example").
		SetQueryParams(map[string]string{
			"page":  page,
			"limit": limit,
		}).
		SetResult(map[string]interface{}{}).
		Get("http://localhost:8080/api/posts")

	if err != nil {
		logger.Errorf("使用resty发送请求失败: %v", err)
		t.err = err
		return t
	}

	// 打印HTTP响应码
	logger.Infof("当前的http响应码: %d", resp.StatusCode())

	// 检查响应状态码
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		logger.Errorf("请求失败，状态码: %d", resp.StatusCode())
		logger.Errorf("响应内容: \n%s", resp.String())
		t.err = fmt.Errorf("get posts failed")
		return t
	}

	// 解析响应结果
	resultPtr := resp.Result().(*map[string]interface{})
	data := *resultPtr

	// 输出格式化的JSON响应
	demoJSONOutput(data)

	return t
}

// GetPostByID 获取指定ID的文章
func (t *APITester) GetPostByID(postID string) *APITester {
	if t.err != nil {
		return t
	}

	logger.Info("\n=== 获取指定ID的文章 ====")

	// 发送GET请求
	resp, err := t.client.R().
		SetHeader("User-Agent", "resty-example-single-post").
		SetResult(map[string]interface{}{}).
		Get(fmt.Sprintf("http://localhost:8080/api/posts/%s", postID))

	if err != nil {
		logger.Errorf("根据ID发送请求失败: %v", err)
		t.err = err
		return t
	}

	// 打印HTTP响应码
	logger.Infof("根据ID查询的HTTP响应码: %d", resp.StatusCode())

	// 检查响应状态码
	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		resultPtr := resp.Result().(*map[string]interface{})
		data := *resultPtr
		logger.Info("根据ID查询的响应:")
		demoJSONOutput(data)
	} else {
		logger.Errorf("请求失败，状态码: %d", resp.StatusCode())
		logger.Errorf("响应内容: \n%s", resp.String())
		t.err = fmt.Errorf("get post by id failed")
	}

	return t
}

// GetPostsWithURLParams 使用URL直接包含查询参数的方式获取文章列表
func (t *APITester) GetPostsWithURLParams(url string) *APITester {
	if t.err != nil {
		return t
	}

	logger.Info("\n=== 使用URL直接包含查询参数的方式获取文章列表 ====")

	// 发送GET请求
	resp, err := t.client.R().
		SetHeader("User-Agent", "resty-example-url-params").
		SetResult(map[string]interface{}{}).
		Get(url)

	if err != nil {
		logger.Errorf("使用URL直接包含参数发送请求失败: %v", err)
		t.err = err
		return t
	}

	// 打印HTTP响应码
	logger.Infof("当前的http响应码: %d", resp.StatusCode())

	// 检查响应状态码
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		logger.Errorf("请求失败，状态码: %d", resp.StatusCode())
		logger.Errorf("响应内容: \n%s", resp.String())
		t.err = fmt.Errorf("get posts with url params failed")
		return t
	}

	// 解析响应结果
	resultPtr := resp.Result().(*map[string]interface{})
	data := *resultPtr

	// 输出格式化的JSON响应
	demoJSONOutput(data)

	logger.Info("\n--- 使用URL直接包含查询参数的方式演示完成 ---")

	return t
}

// Pause 暂停指定秒数
func (t *APITester) Pause(seconds int) *APITester {
	logger.Infof("暂停%d秒...", seconds)
	time.Sleep(time.Duration(seconds) * time.Second)
	return t
}

// Complete 完成测试流程
func (t *APITester) Complete() {
	if t.err == nil {
		logger.Info("\n=== 完整流程执行完成 ====")
	} else {
		logger.Errorf("\n=== 流程执行中断，原因: %v ===", t.err)
	}
}
