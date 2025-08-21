package sort

// SingleNumber 使用排序方法找出只出现一次的数字
func SingleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	
	// 首先对数组进行排序
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	
	// 然后遍历排序后的数组
	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}