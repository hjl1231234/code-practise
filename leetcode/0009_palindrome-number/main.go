package main

import (
	"fmt"

	"github.com/hjl1231234/code-practise/leetcode/0009_palindrome-number/array"
	"github.com/hjl1231234/code-practise/leetcode/0009_palindrome-number/math"
	"github.com/hjl1231234/code-practise/leetcode/0009_palindrome-number/string_method"
)

func main() {
	// 测试数据
	testCases := []int{
		12321,
		-121,
		10,
		12345,
		1221,
	}

	// 遍历所有测试用例
	for i, num := range testCases {
		fmt.Printf("\n测试用例 %d: %d\n", i+1, num)
		fmt.Printf("数学反转解法结果: %t\n", math.IsPalindrome(num))
		fmt.Printf("字符串处理解法结果: %t\n", string_method.IsPalindrome(num))
		fmt.Printf("数组处理解法结果: %t\n", array.IsPalindrome(num))
		fmt.Println("----------------------------------------")
	}
}