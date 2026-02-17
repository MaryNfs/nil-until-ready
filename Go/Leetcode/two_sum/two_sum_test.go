package two_sum

import (
	"slices"
	"testing"
)

func TestTosum(t *testing.T) {
	cases := []struct {
		numbers []int
		targert int
		output  []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}
	for _, c := range cases {
		res := twoSum(c.numbers, c.targert)
		if !slices.Equal(res, c.output) {
			t.Errorf("twoSum(%v, %v) = %v, expected %v", c.numbers, c.targert, res, c.output)
		}
	}
}
