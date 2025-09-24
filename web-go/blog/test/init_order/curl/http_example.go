package curl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/go-resty/resty/v2"
	logger "github.com/learn/init_order/log"
)

func TestRegisterLoginPostMessagesPOST() {
	// TestRegisterLoginPostMessagesPOST 测试完整的用户流程：注册、登录获取token、使用token发帖等，采用链式调用方式

	// 创建API测试器并执行完整流程
	NewAPITester("test2", "password123", "test2@example.com").
		Register().
		Pause(1).
		Login().
		Pause(1).
		CreatePost("测试文章标题 - test2用户", "这是test2用户发布的测试文章内容。").
		Pause(1).
		UpdatePost("30", "更新后的文章标题", "更新后的文章内容，长度需要至少10个字符。").
		Pause(1).
		DeletePost("32").
		Pause(1).
		AddComments("30", []string{
			"这是第1条测试评论，长度需要在5-500个字符之间。",
			"这是第2条测试评论，内容丰富一些，可以包含更多信息。",
			"这是第3条测试评论，作为这篇文章的评论内容。",
		}, 500).
		Pause(1).
		Complete()
}

// TestBatchRegister 测试批量注册功能，采用链式调用方式
func TestBatchRegister() {
	// 创建用户数据数组
	users := make([]BatchUser, 10)
	for i := 0; i < 10; i++ {
		users[i] = BatchUser{
			Username: fmt.Sprintf("batchtest%d", i+1),
			Password: "password123",
			Email:    fmt.Sprintf("batchtest%d@example.com", i+1),
		}
	}

	// 使用链式调用进行批量注册
	NewAPITester("", "", "").
		BatchRegister(users).
		Complete()
}
func TestBatchPostMessagesPOST() {
	// TestBatchPostMessagesPOST 测试并发发送POST请求创建多篇文章（需要认证）- 使用链式调用方式，通过Login动态获取token

	// 创建文章数据数组 - 使用append动态扩容
	posts := make([]Post, 0, 10) // 长度为0，预分配容量为10以优化性能
	username := "test3"
	for i := 0; i < 1000; i++ {
		// 使用append动态添加元素，会自动处理扩容
		title := fmt.Sprintf("测试文章标题%d", i+1)
		post := Post{
			Title:   title,
			Content: fmt.Sprintf("这是用户[%s]创建的第%d篇文章，内容需要至少%d个字符。文章标题: %s", username, i+1, i+1, title),
		}
		posts = append(posts, post)
	}

	// 使用链式调用进行登录并批量创建文章
	// 通过传入用户名和密码，Login方法会动态获取token
	NewAPITester(username, "password123", "test3@example.com").
		Login().                 // 登录并动态获取token
		BatchCreatePosts(posts). // 使用获取到的token批量创建文章
		Complete()               // 完成测试流程
}

func TestLoginEnctypePOST() {
	// TestLoginEnctypePOST 测试不使用Content-Type: application/json，而是使用enctype方式(application/x-www-form-urlencoded)请求登录接口 - 使用链式调用方式

	// 使用链式调用进行form表单格式登录
	NewAPITester("test3", "password123", "test3@example.com").
		LoginWithForm().
		Complete()
}

func TestLoginPOST() {
	// TestLoginPOST 测试登录接口，采用链式调用方式
	NewAPITester("test3", "password123", "test3@example.com").
		Login().
		Complete()
}

func TestPostsGET2() {
	// TestPostsGET2 使用resty库发送GET请求获取文章列表，不同封装方式
	// 创建resty客户端
	client := resty.New()

	// 配置客户端
	client.
		SetTimeout(10*time.Second).
		SetHeader("User-Agent", "resty-example")

	// 发送GET请求
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"page":  "1",
			"limit": "10",
		}).
		SetResult(map[string]interface{}{}). // 设置响应结果的类型为map，用于动态解析JSON
		Get("http://localhost:8080/api/posts")

	if err != nil {
		logger.Errorf("使用resty发送请求失败: %v", err)
		return
	}
	// 打印下当前的http响应码
	logger.Infof("当前的http响应码: %d", resp.StatusCode())

	// 检查响应状态码

	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {

		logger.Errorf("请求失败，状态码: %d", resp.StatusCode())
		logger.Errorf("响应内容: \n%s", resp.String())
		return
	}

	// 解析响应结果
	// 注意：resty库的Result()方法返回的是指针类型*map[string]interface{}
	// 这就是为什么需要先断言为指针，然后再解引用
	resultPtr := resp.Result().(*map[string]interface{})
	data := *resultPtr

	// 调用演示函数，将响应数据序列化为真正的JSON格式输出
	demoJSONOutput(data)

}

func TestSinglePostsGET() {
	// TestSinglePostsGET 测试获取指定文章，采用链式调用方式
	NewAPITester("", "", "").
		GetPostByID("56").
		Complete()
}

func TestPostsGET3() {
	// TestPostsGET3 测试获取文章列表，演示如何直接在URL中包含查询参数，采用链式调用方式
	NewAPITester("", "", "").
		GetPosts("1", "10").
		Complete()
}

// demoJSONOutput 展示如何将响应数据序列化为真正的JSON格式输出
func demoJSONOutput(data map[string]interface{}) {
	// 将map数据序列化为JSON字符串
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		logger.Errorf("JSON序列化失败: %v", err)
		return
	}

	// 输出格式化的JSON
	logger.Info("\n=== 格式化的JSON响应 ===")
	logger.Debug(string(jsonData))
}

// 演示如何在TestPostsGET2函数中调用demoJSONOutput函数的示例
// 在实际代码中，应该在data变量定义后添加以下代码：
// demoJSONOutput(data)

// TestPostsGET4 使用自定义的curl_tool.go实现与TestPostsGET3类似的功能
func TestPostsGET4() {
	// 使用自定义的Client替代resty
	client := NewClient(
		WithTimeout(10 * time.Second),
	)

	// 设置请求头
	client.SetHeader("User-Agent", "custom-curl-example-url-params")

	// 直接在URL中包含查询参数
	resp, err := client.Get(
		context.Background(),
		"http://localhost:8080/api/posts?page=1&limit=10",
	)

	if err != nil {
		logger.Errorf("使用自定义客户端发送请求失败: %v", err)
		return
	}

	// 打印HTTP响应码
	logger.Infof("当前的http响应码: %d", resp.StatusCode)

	// 检查响应状态码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		logger.Errorf("请求失败，状态码: %d", resp.StatusCode)
		// 读取并打印响应内容
		bodyBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close() // 确保关闭Body
		logger.Errorf("响应内容: \n%s", string(bodyBytes))
		return
	}

	// 解析响应结果
	var data map[string]interface{}
	if err := resp.ParseJSON(&data); err != nil {
		logger.Errorf("JSON解析失败: %v", err)
		return
	}

	// 输出格式化的JSON响应
	demoJSONOutput(data)

	/*

		对比curl_tool.go与go-resty/resty/v2的功能差异
			功能对比分析

		- 核心功能对比：两者都支持基本 HTTP 方法、请求头设置、查询参数和 JSON 解析
		- curl_tool.go 欠缺的功能：链式 API 风格、自动 cookie 管理、响应缓存、丰富的响应对象方法、文件上传、HTTP/2 支持、代理支持等
		- curl_tool.go 的优势：代码简单、体积轻量、支持中间件和自定义重试逻辑

	*/

}
