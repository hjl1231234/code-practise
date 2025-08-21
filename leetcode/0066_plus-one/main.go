package main

import (
	"fmt"
	// "unsafe"

	"github.com/hjl1231234/code-practise/leetcode/0066_plus-one/simulation1"
	"github.com/hjl1231234/code-practise/leetcode/0066_plus-one/simulation2"
)

func main() {
	// 测试数据
	testCases := [][]int{
		{2, 2, 3, 4},
		{9, 9, 9, 9},
		{0},
		{1, 2, 9},
		{9, 9},
	}

	// 遍历所有测试用例
	for i, digits := range testCases {

		/*
			testbyte := "hello"
			bytes := []byte(testbyte) // 转为可变的字节切片
			ptr := &bytes[0]          // 现在可以取地址了
			fmt.Println(bytes[0])
			fmt.Printf("ptr: %v, value: %v\n", ptr, *ptr) // 输出地址和字符 'h'
			fmt.Printf("ptr: %p, value: %c\n", ptr, *ptr) // 输出地址和字符 'h'

			testStr := []string{"test1", "test2", "test3"}
			fmt.Printf(" %v\n", testStr[0])
			fmt.Printf(" %v\n", testStr[1])
			fmt.Printf(" %v\n", testStr[2])

			ptrstr := &testStr[0]

			fmt.Println(ptrstr, *ptrstr, &ptrstr)
			fmt.Println(unsafe.Pointer(ptrstr), uintptr(unsafe.Pointer(ptrstr)), unsafe.Pointer(uintptr(unsafe.Pointer(ptrstr))+unsafe.Sizeof(testStr[0])))

			// nextPtr0 := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(ptrstr)) + unsafe.Sizeof(testStr[0])))
			nextPtr0 := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(ptrstr)) + uintptr(8)))
			nextPtr1 := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptrstr)) + unsafe.Sizeof(testStr[0])))

			// num1, _ := strconv.Atoi(*nextPtr0)
			// num2 := strconv.Itoa(*nextPtr1)

			// fmt.Println(nextPtr0, *nextPtr0,strconv.Atoi(*nextPtr0), &nextPtr0)
			fmt.Println(nextPtr0, *nextPtr0, &nextPtr0)
			fmt.Println(nextPtr1, *nextPtr1, &nextPtr1)

			fmt.Printf(" %v\n", &testStr[0])
			fmt.Printf(" %v\n", &testStr[1])
			fmt.Printf(" %v\n", &testStr[2])

		*/

		fmt.Printf("\n测试用例 %d: %v\n", i+1, digits)
		fmt.Printf("模拟进位解法1结果: %v\n", simulation1.PlusOne(append([]int{}, digits...)))
		fmt.Printf("模拟进位解法2结果: %v\n", simulation2.PlusOne(append([]int{}, digits...)))
		fmt.Println("----------------------------------------")
	}
}
