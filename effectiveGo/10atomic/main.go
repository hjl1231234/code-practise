package main

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

var logFile, err = os.Create("/home/hjl/web3project/code-practise/effectiveGo/10atomic/go.log")

func main() {
	// 检查日志文件创建是否成功
	if err != nil {
		fmt.Printf("无法创建日志文件: %v\n", err)
		return
	}
	defer logFile.Close()

	// 记录开始时间
	startTime := time.Now()

	// 定义原子计数器
	var counter int64

	// 打印初始值
	msg := fmt.Sprintf("初始值: %d\n", atomic.LoadInt64(&counter))
	fmt.Print(msg)
	logFile.WriteString(msg)

	// 创建WaitGroup以等待所有协程完成
	var wg sync.WaitGroup

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 每个协程增加1000次计数
			for j := 0; j < 10000; j++ {
				// 原子递增操作
				atomic.AddInt64(&counter, 1)
				// 输出到控制台和日志文件
				msg := fmt.Sprintf("协程%d, 增加后: %d  -- %d\n", i, atomic.LoadInt64(&counter))
				// fmt.Print(msg)
				logFile.WriteString(msg)
			}
		}()
	}

	// 等待所有协程完成
	wg.Wait()

	// 记录结束时间并计算耗时
	elapsedTime := time.Since(startTime).Seconds()
	msg = fmt.Sprintf("程序执行时间: %.6f 秒\n", elapsedTime)
	fmt.Print(msg)
	logFile.WriteString(msg)
}
