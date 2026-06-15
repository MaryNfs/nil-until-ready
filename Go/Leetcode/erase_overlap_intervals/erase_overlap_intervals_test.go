package erase_overlap_intervals

import (
	"fmt"
	"testing"
)

func TestEraseOverlapIntervals(t *testing.T) {
	cases := []struct {
		intervals [][]int
		out       int
	}{
		{[][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}, []int{1, 3}}, 1},
		{[][]int{[]int{1, 2}, []int{1, 2}, []int{1, 2}}, 2},
		{[][]int{[]int{1, 2}, []int{2, 3}}, 0},
	}
	for _, c := range cases {
		res := eraseOverlapIntervals(c.intervals)

		if res != c.out{
			fmt.Printf("eraseOverlapIntervals(%v)=%v, but expected %v.",c.intervals,res,c.out)
		}
	}
}
