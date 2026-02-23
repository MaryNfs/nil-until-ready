package count_bits

import (
	"slices"
	"testing"
)

func TestCountBits(t *testing.T) {
	cases := []struct {
		n   int
		out []int
	}{
		{2, []int{0, 1, 1}},
		{5, []int{0, 1, 1, 2, 1, 2}},
	}
	for _, c := range cases {
		res := countBits(c.n)
		if !slices.Equal(res, c.out) {
			t.Errorf("countBits(%v)=%v but expected %v", c.n, res, c.out)
		}
	}
}
