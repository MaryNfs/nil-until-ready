package tribonacci

import "testing"

func TestTribonacci(t *testing.T) {
	cases := []struct {
		n   int
		out int
	}{
		{4, 4},
		{25, 1389537},
		{0, 0},
	}
	for _, c := range cases {
		res := tribonacci(c.n)
		if res != c.out {
			t.Errorf("tribonacci(%v)=%v but expected %v", c.n, res, c.out)
		}
	}
}
