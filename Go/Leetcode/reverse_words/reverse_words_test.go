package reverse_words

import "testing"

func TestReverseWords(t *testing.T) {
	cases := []struct {
		input string
		out   string
	}{
		{"a good   example", "example good a"},
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
	}

	for _, c := range cases {
		res := reverseWords(c.input)
		if res != c.out {
			t.Errorf("reverseWords(%v)=%v, but expected %v", c.input, res, c.out)
		}
	}
}
