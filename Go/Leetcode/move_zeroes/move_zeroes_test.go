package move_zeroes

import (
	"slices"
	"testing"
)

var nums []int

func TestMoveZeroes(t *testing.T) {
	cases := []struct {
		in []int
		ex []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
	}

	for _, c := range cases {
		moveZeroes(c.in)
		if !slices.Equal(c.in, c.ex) {
			t.Errorf("moveZeroes=%v, but it should be %v", c.in, c.ex)
		}
	}
}
