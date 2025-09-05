package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// Counter 包含一个计数器和一个互斥锁
type Counter struct {
	mu    sync.Mutex
	count int
}

// Increment 增加计数器的值，使用互斥锁保护
func (c *Counter) Increment(i int) {
	c.mu.Lock()         // 加锁
	defer c.mu.Unlock() // 确保在函数返回时解锁
	c.count++
	msg := fmt.Sprintf("协程%d, 增加后: %d\n", i, c.count)
	// fmt.Print(msg)
	logFile.WriteString(msg)

}

var logFile, err = os.Create("/home/hjl/web3project/code-practise/effectiveGo/9Mutex/go.log")

func main() {
	// 创建日志文件，使用绝对路径
	if err != nil {
		fmt.Printf("无法创建日志文件: %v\n", err)
		return
	}
	defer logFile.Close()

	// 记录开始时间
	startTime := time.Now()

	// 创建计数器实例
	counter := &Counter{}
	msg := fmt.Sprintf("初始值: %d\n", counter.count)
	// fmt.Print(msg)
	logFile.WriteString(msg)

	// 创建WaitGroup以等待所有协程完成
	var wg sync.WaitGroup

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 每个协程增加10次计数
			for j := 0; j < 100; j++ {
				counter.Increment(i)

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
