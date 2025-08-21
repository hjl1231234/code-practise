package main

import (
	"fmt"

	"github.com/hjl1231234/code-practise/leetcode/0014_longest-common-prefix/brute"
	"github.com/hjl1231234/code-practise/leetcode/0014_longest-common-prefix/divide"
	"github.com/hjl1231234/code-practise/leetcode/0014_longest-common-prefix/horizontal"
	"github.com/hjl1231234/code-practise/leetcode/0014_longest-common-prefix/shortest"
	"github.com/hjl1231234/code-practise/leetcode/0014_longest-common-prefix/sort"
	"github.com/hjl1231234/code-practise/leetcode/0014_longest-common-prefix/trie"
)

func main() {
	// 测试数据
	testCases := [][]string{
		{"flower", "flow", "flight"},
		{"dog", "racecar", "car"},
		{"a"},
		{""},
		{"abc", "abc", "abc"},
	}

	// 遍历所有测试用例
	for i, strs := range testCases {
		fmt.Printf("\n测试用例 %d: %v\n", i+1, strs)
		fmt.Printf("最短字符串遍历解法结果: %s\n", shortest.LongestCommonPrefix(strs))
		fmt.Printf("暴力解法结果: %s\n", brute.LongestCommonPrefix(strs))
		fmt.Printf("排序比较解法结果: %s\n", sort.LongestCommonPrefix(strs))
		fmt.Printf("Trie树解法结果: %s\n", trie.LongestCommonPrefix(strs))
		fmt.Printf("水平扫描解法结果: %s\n", horizontal.LongestCommonPrefix(strs))
		fmt.Printf("分治解法结果: %s\n", divide.LongestCommonPrefix(strs))
		fmt.Println("----------------------------------------")
	}
}