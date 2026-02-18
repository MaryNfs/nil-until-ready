package coin_change

import "testing"

func TestCoinChange(t *testing.T) {
	cases := []struct {
		coins  []int
		amount int
		output int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
		{[]int{186, 419, 83, 408}, 6249, 20},
	}

	for _, c := range cases {
		res := coinChange(c.coins,c.amount)
		if res != c.output{
			t.Errorf("coinChange(%v,%v)=%v, expected: %v",c.coins,c.amount,res,c.output)
		}
	}
}
