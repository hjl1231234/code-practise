package main

import (
	"fmt"
	"time"

	"github.com/Jeffail/tunny"
)

func main() {
	// 创建一个协程池，最多5个协程
	// 从 Jeffail/tunny 库的版本来看，可能使用的是旧版本，新版本使用 tunny.NewFunc 创建协程池
	pool := tunny.NewFunc(5, func(data interface{}) interface{} {
		return struct{}{}
	})
	defer pool.Close()

	for i := 0; i < 10; i++ {
		i := i
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Printf("任务 %d 完成\n", i)
		}()
		pool.Process(i)
	}

	// 等待所有任务完成
	time.Sleep(5 * time.Second)
}
