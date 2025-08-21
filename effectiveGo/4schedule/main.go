package main

import (
	"fmt"
	"sync"
	"time"
)

// schedule 调度器函数：接收任务切片，并发执行并统计耗时
func schedule(tasks []func()) {
	var wg sync.WaitGroup
	
	for i, task := range tasks {
		wg.Add(1)
		go func(index int, t func()) {
			defer wg.Done()
			start := time.Now()
			t() // 执行任务
			duration := time.Since(start)
			fmt.Printf("任务%d执行耗时:%v\n", index+1, duration)
		}(i, task)
	}
	
	wg.Wait()
}

func main() {
	// 定义任务：模拟耗时任务
	task1 := func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("任务1完成")
	}
	
	task2 := func() {
		time.Sleep(800 * time.Millisecond)
		fmt.Println("任务2完成")
	}
	
	task3 := func() {
		time.Sleep(300 * time.Millisecond)
		fmt.Println("任务3完成")
	}
	
	// 调用调度器函数
	schedule([]func(){task1, task2, task3})
}