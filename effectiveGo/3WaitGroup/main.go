package main

import (
	"fmt"
	"sync"
)

// 通用打印函数：接收参数控制打印逻辑
// 参数说明：
// - name: goroutine标识（用于打印区分）
// - start: 起始数字
// - max: 最大数字
// - step: 步长（控制奇偶打印，如1->3->5用步长2）
// - waitCh: 等待信号的channel（自身等待被唤醒）
// - sendCh: 发送信号的channel（唤醒另一个goroutine）
// - wg: WaitGroup指针（用于通知完成）
func printNumbers(name string, start, max, step int, waitCh, sendCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done() // 退出时通知WaitGroup

	for i := start; i <= max; i += step {
		<-waitCh // 等待对方发送的信号（阻塞直到有信号）
		fmt.Printf("%s: %d\n", name, i)

		// 判断是否需要发送下一个信号（避免最后一次发送导致阻塞）
		// 下一个数字若超过max，则无需发送
		if i < max {
			sendCh <- struct{}{}
		}
	}
}

func main() {
	var wg sync.WaitGroup
	maxNum := 10

	// 创建两个channel用于信号传递
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	// 启动第一个goroutine：打印奇数（1,3,5...）
	wg.Add(1)
	go printNumbers(
		"奇数",   // 标识名
		1,      // 起始数字
		maxNum, // 最大数字
		2,      // 步长（每次+2）
		ch1,    // 等待ch1的信号
		ch2,    // 向ch2发送信号
		&wg,    // WaitGroup指针
	)

	// 启动第二个goroutine：打印偶数（2,4,6...）
	wg.Add(1)
	go printNumbers(
		"偶数",   // 标识名
		2,      // 起始数字
		maxNum, // 最大数字
		2,      // 步长（每次+2）
		ch2,    // 等待ch2的信号
		ch1,    // 向ch1发送信号
		&wg,    // WaitGroup指针
	)

	// 发送初始信号，让第一个goroutine先执行
	ch1 <- struct{}{}

	// 等待所有goroutine完成
	wg.Wait()
	close(ch1)
	close(ch2)
	fmt.Println("程序结束")
}
