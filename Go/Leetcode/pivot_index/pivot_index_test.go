package pivot_index

import "testing"

func TestSolve(t *testing.T) {
	cases := []struct {
		nums []int
		out  int
	}{
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1, 2, 3}, -1},
		{[]int{2, 1, -1}, 0},
	}

	for _, c := range cases {
		res := pivotIndex(c.nums)
		if res != c.out {
			t.Errorf("pivotIndex(%v)=%v but expected %v.", c.nums, res, c.out)
		}
	}
}
