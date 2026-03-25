package top_K_frequent

import (
	"slices"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		out  []int
	}{
		{[]int{3, 3, 3, 1}, 2, []int{3, 1}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, 2, []int{1, 2}},
	}

	for _, c := range cases {
		res := topKFrequent(c.nums, c.k)
		slices.Sort(res)
		slices.Sort(c.out)
		if !slices.Equal(res, c.out) {
			t.Errorf("topKFrequent(%v,%v)=%v, but it should be %v", c.nums, c.k, res, c.out)
		}
	}
}
