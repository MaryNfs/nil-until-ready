package longest_ones

import "testing"

func TestLongestOnes(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		out  int
	}{
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
		{[]int{0, 0}, 2, 2},
	}
	for _, c := range cases {
		res := longestOnes(c.nums, c.k)
		if c.out != res {
			t.Errorf("longestOnes(%v,%v)=%v. but expected %v", c.nums, c.k, res, c.out)
		}
	}
}
