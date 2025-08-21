package main

import (
	"fmt"

	"github.com/hjl1231234/code-practise/leetcode/0136_single-number/bruteforce"
	"github.com/hjl1231234/code-practise/leetcode/0136_single-number/hashmap"
	"github.com/hjl1231234/code-practise/leetcode/0136_single-number/sort"
	"github.com/hjl1231234/code-practise/leetcode/0136_single-number/xor"
)

func main() {
	// 测试数据
	testCases := [][]int{
		{4, 1, 2, 1, 2},
		{2, 2, 1},
		{1},
	}

	// 遍历所有测试用例
	for i, nums := range testCases {
		fmt.Printf("\n测试用例 %d: %v\n", i+1, nums)
		fmt.Printf("异或解法结果: %d\n", xor.SingleNumber(nums))
		fmt.Printf("暴力解法结果: %d\n", bruteforce.SingleNumber(nums))
		fmt.Printf("哈希表解法结果: %d\n", hashmap.SingleNumber(nums))
		fmt.Printf("排序解法结果: %d\n", sort.SingleNumber(nums))
		fmt.Println("----------------------------------------")
	}
}
