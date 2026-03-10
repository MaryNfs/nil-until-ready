package find_max_average

import "testing"

func TestFindMaxAverage(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		out  float64
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75000},
		{[]int{5}, 1, 5.00000},
		{[]int{-1}, 1, -1.00000},
	}
	for _, c := range cases {
		res := findMaxAverage(c.nums, c.k)
		if res != c.out {
			t.Errorf("findMaxAverage(%v, %v)=%v, but expected %v", c.nums, c.k, res, c.out)
		}
	}
}
