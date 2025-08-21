package main

import "fmt"

// 函数：接收整数切片的指针，将每个元素乘以2
func multiplyByTwo(nums *[]int) {
	for i := range *nums {
		(*nums)[i] *= 2
	}
}
func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("修改前:", numbers)
	multiplyByTwo(&numbers) //传递切片的指针
	fmt.Println("修改后:", numbers)
}
