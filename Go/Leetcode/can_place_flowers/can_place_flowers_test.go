package can_place_flowers

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	cases := []struct {
		flowerbed []int
		n         int
		output    bool
	}{
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 2, false},
		{[]int{1, 0, 0, 0, 0, 1}, 2, false},
		{[]int{0}, 1, true},
		{[]int{1}, 0, true},
	}
	for _, c := range cases {
		res := canPlaceFlowers(c.flowerbed, c.n)
		if res != c.output {
			t.Errorf("canPlaceFlowers(%v,%v)=%v, expected %v", c.flowerbed, c.n, res, c.output)
		}
	}
}
