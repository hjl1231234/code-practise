package xor

// SingleNumber 使用异或解法找出只出现一次的数字
func SingleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		res = res ^ n
	}
	return res
}
