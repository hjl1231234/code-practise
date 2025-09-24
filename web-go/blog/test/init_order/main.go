package main

import (
	"fmt"

	"github.com/learn/init_order/curl"
	logger "github.com/learn/init_order/log"
)

func init() {
	fmt.Println("main init method invoked")
}

func main() {
	// 初始化日志系统
	logger.InitLogger()
	logger.Info("main method invoked!")
	// curl.ExampleJSONPostRequest()
	// curl.TestRegisterPOST()
	// curl.TestLoginPOST()
	// curl.TestSinglePostsGET()
	// curl.TestLoginEnctypePOST()
	// curl.TestRegisterLoginPostMessagesPOST()
	curl.TestBatchPostMessagesPOST()
	// curl.TestPostsGET3()
	// curl.TestPostsGET4()

}
