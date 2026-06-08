package num_squares

import (
	"fmt"
	"testing"
)

func TestNumSquares(t *testing.T) {
	cases := []struct {
		in  int
		out int
	}{
		{12, 3},
		{13, 2},
	}
	for _, c := range cases {
		res := numSquares(c.in)
		if res != c.out {
			fmt.Printf("numSquares(%v)=%v, expected %v", c.in, res, c.out)
		}
	}
}
