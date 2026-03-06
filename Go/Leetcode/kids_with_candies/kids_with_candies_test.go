package kids_with_candies

import (
	"slices"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	cases := []struct {
		candies      []int
		extraCandies int
		res          []bool
	}{
		{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
		{[]int{12, 1, 12}, 10, []bool{true, false, true}},
	}

	for _, c := range cases {
		res := kidsWithCandies(c.candies, c.extraCandies)
		if !slices.Equal(res, c.res) {
			t.Errorf("kidsWithCandies(%v,%v)=%v, but expected: %v", c.candies, c.extraCandies, res, c.res)
		}
	}
}
