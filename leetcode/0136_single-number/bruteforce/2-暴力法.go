package bruteforce

// SingleNumber 使用暴力方法找出只出现一次的数字
func SingleNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		flag := false
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] && i != j {
				flag = true
				break
			}
		}
		if flag == false {
			return nums[i]
		}
	}
	return 0 // 保证函数总是有返回值
}
