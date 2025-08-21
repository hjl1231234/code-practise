package main

import (
	"fmt"
	"time"
)

// producer 函数: 生成1到10的整数并发送到通道
func producer(ch chan<- int) {
	for i := 1; i <= 10; i++ {

		ch <- i
		fmt.Printf("发送: %d\n", i)
		time.Sleep(100 * time.Millisecond) // 短暂休眠，使输出更清晰
	}
	close(ch) // 发送完成后关闭通道
	fmt.Println("生产者: 已关闭通道")
}

// consumer 函数: 从通道接收整数并打印
func consumer(ch <-chan int, done chan<- bool) {
	for num := range ch {
		fmt.Printf("接收到: %d\n", num)
	}
	fmt.Println("消费者: 通道已关闭，退出")
	done <- true // 通知主协程消费者已完成
}

func main() {
	// 1. 基本的生产者-消费者模式
	fmt.Println("=== 基本通道通信演示 ===")
	ch := make(chan int)    // 创建无缓冲通道
	done := make(chan bool) // 创建用于通知完成的通道

	go producer(ch)       // 启动生产者协程
	go consumer(ch, done) // 启动消费者协程

	<-done // 等待消费者完成

	// 非阻塞发送
	// select {}

	fmt.Println("程序正常结束")
}
