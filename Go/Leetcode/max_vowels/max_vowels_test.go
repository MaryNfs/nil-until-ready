package max_vowels

import "testing"

func TestMaxVowels(t *testing.T) {
	cases := []struct {
		s   string
		k   int
		out int
	}{
		{"abciiidef", 3, 3},
		{"aeiou", 2, 2},
		{"leetcode", 3, 2},
		{"weallloveyou", 7, 4},
	}
	for _, c := range cases {
		res := maxVowels(c.s, c.k)
		if res != c.out {
			t.Errorf("maxVowels(%v,%v)=%v but expected %v", c.s, c.k, res, c.out)
		}
	}
}
