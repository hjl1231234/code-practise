package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// 定义原子计数器
	var counter int64

	// 打印初始值
	fmt.Printf("初始值: %d\n", atomic.LoadInt64(&counter))

	// 创建WaitGroup以等待所有协程完成
	var wg sync.WaitGroup

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// 每个协程增加1000次计数
			for j := 0; j < 100; j++ {
				// 原子递增操作
				atomic.AddInt64(&counter, 1)
				fmt.Printf("协程%d, 增加后: %d\n", id, atomic.LoadInt64(&counter))
			}
		}(i)
	}

	// 等待所有协程完成
	wg.Wait()

	// 原子加载并输出最终计数
	finalCount := atomic.LoadInt64(&counter)
	fmt.Printf("最终计数器值: %d\n", finalCount)
}
