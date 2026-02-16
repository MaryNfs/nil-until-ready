package word_pattern

import "testing"

func TestWordPattern(t *testing.T) {
	expected := []struct {
		Pattern string
		S       string
		Output  bool
	}{
		{"abba", "dog cat cat dog", true},
		{"abba", "dog cat cat fish", false},
		{"aaaa", "dog cat cat dog", false},
		{"ab", "dog cat cat dog", false},
	}
	for _, e := range expected {
		res := wordPattern(e.Pattern, e.S)
		if e.Output != res {
			t.Errorf("wordPattern(%s, %s)=%v but expected %v", e.Pattern, e.S, res, e.Output)
		}
	}
}
