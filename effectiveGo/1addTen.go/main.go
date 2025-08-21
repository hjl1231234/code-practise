package main

import (
	"fmt"
)

// 定义函数，接收一个整数指针参数
func addTen(num *int) {
	*num = *num + 10 //通过解引用修改指针指向的值
}
func main() {
	value := 20 //定义一个整数变量
	fmt.Println("修改前:", value)
	addTen(&value) //传入变量的地址(指针)
	fmt.Println("修改后:", value)
}
