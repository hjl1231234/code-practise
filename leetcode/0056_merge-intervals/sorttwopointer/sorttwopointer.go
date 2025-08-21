package sorttwopointer

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
	for i := 0; i < len(intervals); {
		end := intervals[i][1]
		j := i + 1
		for j < len(intervals) && intervals[j][0] <= end {
			if intervals[j][1] > end {
				end = intervals[j][1]
			}
			j++
		}
		res = append(res, []int{intervals[i][0], end})
		i = j
	}
	return res
}
