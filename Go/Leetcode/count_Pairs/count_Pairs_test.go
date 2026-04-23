package count_Pairs

import "testing"

func TestCountPairs(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		out    int
	}{
		{[]int{-1, 1, 2, 3, 1}, 2, 3},
		{[]int{-6, 2, 5, -2, -7, -1, 3}, -2, 10},
	}
	for _, c := range cases {
		res := countPairs(c.nums, c.target)
		if res != c.out {
			t.Errorf("countPairs(%v,%v)=%v, but expected %v", c.nums, c.target, res, c.out)
		}
	}
}
