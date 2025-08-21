package main

import (
	"fmt"

	arr2 "github.com/hjl1231234/code-practise/leetcode/0020_valid-parentheses/array2"
	arr3 "github.com/hjl1231234/code-practise/leetcode/0020_valid-parentheses/array3"
	stk "github.com/hjl1231234/code-practise/leetcode/0020_valid-parentheses/stack"
)

func main() {
	// 测试数据
	testCases := []string{
		"()[]{",
		"()[]{})",
		"({})",
		"",
		"(",
	}

	// 遍历所有测试用例
	for i, s := range testCases {
		fmt.Printf("\n测试用例 %d: %s\n", i+1, s)
		fmt.Printf("栈结构解法结果: %t\n", stk.IsValid(s))
		fmt.Printf("数组实现解法1结果: %t\n", arr2.IsValid(s))
		fmt.Printf("数组实现解法2结果: %t\n", arr3.IsValid(s))
		fmt.Println("----------------------------------------")
	}
}