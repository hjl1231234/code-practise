package trie

import (
	"math"
)

// 移除main函数

var trie [][]int
var index int

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	trie = make([][]int, 2000)
	for k := range trie {
		value := make([]int, 26)
		trie[k] = value
	}
	insert(strs[0])

	minValue := math.MaxInt32
	for i := 1; i < len(strs); i++ {
		retValue := insert(strs[i])
		if minValue > retValue {
			minValue = retValue
		}
	}
	return strs[0][:minValue]
}

func insert(str string) int {
	p := 0
	count := 0
	for i := 0; i < len(str); i++ {
		ch := str[i] - 'a'
		if value := trie[p][ch]; value == 0 {
			index++
			trie[p][ch] = index
		} else {
			count++
		}
		p = trie[p][ch]
	}
	return count
}
