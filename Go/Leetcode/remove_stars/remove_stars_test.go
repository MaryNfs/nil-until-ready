package remove_stars

import "testing"

func TestRemoveStars(t *testing.T) {
	cases := []struct {
		s   string
		out string
	}{
		{"leet**cod*e", "lecoe"},
		{"erase*****", ""},
		{"l*", ""},
	}
	for _, c := range cases {
		res := removeStars(c.s)
		if res != c.out {
			t.Errorf("removeStars(%v)=%v but expected %v",c.s,res,c.out)
		}
	}
}
