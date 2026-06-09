package erase_overlap_intervals

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := 0
	e := intervals[0][1]

	for _, arr := range intervals[1:] {
		if arr[0] >= e {
			e = arr[1]
		} else {
			res++
			e = min(e, arr[1])
		}
	}

	return res
}
