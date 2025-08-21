package main

import (
	"fmt"
	"sync"
	"time"
)

// 生产者函数：向通道发送100个整数
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Printf("生产: %d\n", i)
	}
	close(ch) // 关闭通道，表示不再发送数据
}

// 消费者函数：从通道接收并打印数据
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Printf("消费: %d\n", num)
		// 模拟处理时间
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// 创建带缓冲的通道，缓冲区大小为20
	ch := make(chan int, 20)

	// 创建WaitGroup以同步协程
	var wg sync.WaitGroup

	// 增加WaitGroup计数，表示有一个生产者协程
	wg.Add(2)
	// 启动生产者协程
	go producer(ch, &wg)

	// 增加WaitGroup计数，表示有一个消费者协程
	// 启动消费者协程
	go consumer(ch, &wg)

	// 等待所有协程完成
	wg.Wait()

	fmt.Println("所有数据处理完成")
}
