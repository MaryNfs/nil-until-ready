package reverse_vowels

import "testing"

func TestReverseVowels(t *testing.T) {
	cases := []struct {
		s string
		o string
	}{
		{"IceCreAm", "AceCreIm"},
		{"leetcode", "leotcede"},
	}

	for _, c := range cases {
		res := reverseVowels(c.s)
		if res != c.o {
			t.Errorf("reverseVowels(%v)=%v but expected %v", c.s, res, c.o)
		}
	}
}
