package array

import (
	"strconv"
)

// leetcode 9 回文数
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	arrs := []byte(strconv.Itoa(x))
	Len := len(arrs)
	for i := 0; i < Len/2; i++ {
		if arrs[i] != arrs[Len-i-1] {
			return false
		}
	}
	return true
}
