package max_area

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	cases := []struct {
		heights []int
		output  int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{0, 2}, 0},
		{[]int{8, 7, 2, 1}, 7},
	}

	for _, c := range cases {
		res := maxArea(c.heights)
		if res != c.output {
			t.Errorf("maxArea(%v)=%v, expected %v", c.heights, res, c.output)
		}
	}
}
