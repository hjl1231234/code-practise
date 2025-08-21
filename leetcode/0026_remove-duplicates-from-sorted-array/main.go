package main

import (
	"fmt"

	"github.com/hjl1231234/code-practise/leetcode/0026_remove-duplicates-from-sorted-array/counting"
	"github.com/hjl1231234/code-practise/leetcode/0026_remove-duplicates-from-sorted-array/twopointer"
)

func main() {
	// 测试数据
	testCases := [][]int{
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		{1, 1, 2},
		{},
		{1},
		{1, 1, 1, 1},
	}

	// 遍历所有测试用例
	for i, nums := range testCases {
		// 创建测试用例的副本，避免修改原始数据
		twopointerNums := make([]int, len(nums))
		copy(twopointerNums, nums)
		countingNums := make([]int, len(nums))
		copy(countingNums, nums)

		fmt.Printf("\n测试用例 %d: %v\n", i+1, nums)
		
		// 双指针解法
		twopointerResult := twopointer.RemoveDuplicates(twopointerNums)
		fmt.Printf("双指针解法结果: %d", twopointerResult)
		if twopointerResult > 0 {
			fmt.Printf(", 修改后的数组: %v\n", twopointerNums[:twopointerResult])
		} else {
			fmt.Println(", 修改后的数组: []")
		}
		
		// 计数法解法
		countingResult := counting.RemoveDuplicates(countingNums)
		fmt.Printf("计数法解法结果: %d", countingResult)
		if countingResult > 0 {
			fmt.Printf(", 修改后的数组: %v\n", countingNums[:countingResult])
		} else {
			fmt.Println(", 修改后的数组: []")
		}
		
		fmt.Println("----------------------------------------")
	}
}