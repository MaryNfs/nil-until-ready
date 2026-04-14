package predict_party_victory

import "testing"

func TestPredictPartyVictory(t *testing.T) {
	cases := []struct {
		input  string
		output string
	}{
		{"RD", "Radiant"},
		{"RDD", "Dire"},
		{"DDRRR", "Dire"},
	}
	for _, c := range cases {
		res := predictPartyVictory(c.input)
		if res != c.output{
			t.Errorf("predictPartyVictory(%v)=%v, but expected %v",c.input,res,c.output)
		}
	}
}
