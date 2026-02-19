package single_number

import "testing"

func TestSingleNumber(t *testing.T) {
	cases := []struct {
		numbers []int
		output  int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
	}

	for _, c := range cases {
		res := singleNumber(c.numbers)
		if res != c.output {
			t.Errorf("singleNumber(%v)=%v, expected %v", c.numbers, res, c.output)
		}
	}
}
