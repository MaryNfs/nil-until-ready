package hamming_weight

import "testing"

func TestHammingWeight(t *testing.T) {
	cases := []struct {
		number int
		output int
	}{
		{11, 3},
		{128, 1},
		{2147483645, 30},
		{1, 1},
		{2, 1},
	}
	for _,c := range cases{
		res := hammingWeight(c.number)
		if res != c.output {
			t.Errorf("hammingWeight(%v)=%v, expected %v",c.number,c.output,res)
		}
	}
}
