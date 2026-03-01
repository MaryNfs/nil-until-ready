package is_subsequence

import (
	"testing"
)

func TestIsSubsequence(t *testing.T) {
	cases := []struct {
		s   string
		t   string
		out bool
	}{
		{"a", "v", false},
		{"abc", "aefbettgc", true},
		{"", "ff", true},
		{"er", "sdfsdfsd", false},
	}

	for _, c := range cases {
		res := isSubsequence(c.s, c.t)
		if res != c.out {
			t.Errorf("isSubsequence(%v, %v)=%v, but expected %v", c.s, c.t, res, c.out)
		}
	}
}
