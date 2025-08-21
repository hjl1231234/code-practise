package sorttraversal

import (
	"sort"
)

// 移除main函数

// Merge 合并区间
func Merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	if len(intervals) == 0 {
		return nil
	}
	
	// 过滤空区间和无效区间
	validIntervals := make([][]int, 0)
	for _, interval := range intervals {
		if len(interval) >= 2 {
			validIntervals = append(validIntervals, interval)
		}
	}
	
	if len(validIntervals) == 0 {
		return res
	}
	
	sort.Slice(validIntervals, func(i, j int) bool {
		return validIntervals[i][0] < validIntervals[j][0]
	})
	intervals = validIntervals
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		arr := res[len(res)-1]
		if intervals[i][0] > arr[1] {
			res = append(res, intervals[i])
		} else if intervals[i][1] > arr[1] {
			res[len(res)-1][1] = intervals[i][1]
		}
	}
	return res
}
