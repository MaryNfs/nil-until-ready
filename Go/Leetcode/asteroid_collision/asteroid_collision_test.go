package asteroid_collision

import (
	"slices"
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	cases := []struct {
		asts []int
		out  []int
	}{
		{[]int{5, 10, -5}, []int{5, 10}},
		{[]int{8, -8}, []int{}},
		{[]int{10, 2, -5}, []int{10}},
		{[]int{3, 5, -6, 2, -1, 4}, []int{-6, 2, 4}},
	}
	for _, c := range cases {
		res := asteroidCollision(c.asts)
		if !slices.Equal(res, c.out) {
			t.Errorf("asteroidCollision(%v)=%v but expected %v", c.asts, res, c.out)
		}
	}
}
