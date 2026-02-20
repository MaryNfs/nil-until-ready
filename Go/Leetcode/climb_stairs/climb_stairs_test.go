package climb_stairs

import "testing"

func TestClimbStairs(t *testing.T) {
	cases := []struct {
		n int
		o int
	}{
		{4, 5},
		{2, 2},
		{3, 3},
	}
	for _, c := range cases {
		res := climbStairs(c.n)
		if res != c.o {
			t.Errorf("climbStairs(%v)=%v, expected %v", c.n, res, c.o)
		}
	}
}
