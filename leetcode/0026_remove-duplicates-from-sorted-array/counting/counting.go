package counting

// 移除未使用的fmt导入
// 移除main函数

// RemoveDuplicates 删除排序数组中的重复项
func RemoveDuplicates(nums []int) int {
	// 处理空数组情况
	if len(nums) == 0 {
		return 0
	}
	
	count := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			nums[count] = nums[i+1]
			count++
		}
	}
	return count
}
