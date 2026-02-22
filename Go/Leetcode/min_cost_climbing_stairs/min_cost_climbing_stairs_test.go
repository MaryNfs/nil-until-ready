package min_cost_climbing_stairs

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	cases := []struct {
		cost []int
		out  int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}

	for _, c := range cases {
		res := minCostClimbingStairs(c.cost)
		if res != c.out {
			t.Errorf("minCostClimbingStairs(%v)=%v, expected %v", c.cost, res, c.out)
		}
	}
}
