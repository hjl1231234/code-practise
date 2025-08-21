package main

import (
	"fmt"

	"github.com/hjl1231234/code-practise/leetcode/0056_merge-intervals/sorttraversal"
	"github.com/hjl1231234/code-practise/leetcode/0056_merge-intervals/sorttwopointer"
)

func main() {
	// 测试数据
	testCases := [][][]int{
		{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
		{{1, 4}, {4, 5}},
		{{1, 4}, {0, 4}},
		{{1, 4}, {0, 0}},
		{{}},
	}

	// 遍历所有测试用例
	for i, intervals := range testCases {
		fmt.Printf("\n测试用例 %d: %v\n", i+1, intervals)
		
		// 排序遍历解法
		sorttraversalResult := sorttraversal.Merge(intervals)
		fmt.Printf("排序遍历解法结果: %v\n", sorttraversalResult)
		
		// 排序双指针解法
		sorttwopointerResult := sorttwopointer.Merge(intervals)
		fmt.Printf("排序双指针解法结果: %v\n", sorttwopointerResult)
		
		fmt.Println("----------------------------------------")
	}
}