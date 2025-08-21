package main

import (
	"fmt"
	"sync"
)

// Counter 包含一个计数器和一个互斥锁
type Counter struct {
	mu    sync.Mutex
	count int
}

// Increment 增加计数器的值，使用互斥锁保护
func (c *Counter) Increment() {
	c.mu.Lock()         // 加锁
	defer c.mu.Unlock() // 确保在函数返回时解锁
	c.count++
}

func main() {
	// 创建计数器实例
	counter := &Counter{}
	fmt.Printf("初始值: %d\n", counter.count)

	// 创建WaitGroup以等待所有协程完成
	var wg sync.WaitGroup

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 每个协程增加10次计数
			for j := 0; j < 100; j++ {
				counter.Increment()
				fmt.Printf("协程%d, 增加后: %d\n", i, counter.count)
			}
		}()
	}

	// 等待所有协程完成
	wg.Wait()

	// 输出最终计数
	fmt.Printf("最终计数器值: %d\n", counter.count)
}
