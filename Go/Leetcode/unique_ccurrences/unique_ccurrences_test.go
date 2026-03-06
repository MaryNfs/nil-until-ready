package unique_ccurrences

import "testing"

func TestUniqueOccurrences(t *testing.T) {
	cases := []struct {
		arr []int
		res bool
	}{
		{[]int{1, 2, 2, 1, 1, 3}, true},
		{[]int{1, 2}, false},
		{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
		{[]int{1}, true},
	}
	for _, c := range cases {
		res := uniqueOccurrences(c.arr)
		if res != c.res {
			t.Errorf("uniqueOccurrences(%v)=%v but expected: %v", c.arr, res, c.res)
		}
	}
}
