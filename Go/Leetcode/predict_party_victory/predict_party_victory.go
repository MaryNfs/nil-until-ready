package predict_party_victory

func predictPartyVictory(senate string) string {
	dq := make([]int, 0)
	rq := make([]int, 0)
	for i, c := range senate {
		if c == 'R' {
			rq = append(rq, i)
		} else {
			dq = append(dq, i)
		}
	}
	for len(dq) > 0 && len(rq) > 0 {
		dTurn := dq[0]
		dq = dq[1:]
		rTurn := rq[0]
		rq = rq[1:]
		if dTurn > rTurn {
			rq = append(rq, rTurn+len(senate))
		} else {
			dq = append(dq, dTurn+len(senate))
		}
	}

	if len(dq) > 0 {
		return "Dire"
	} else {
		return "Radiant"
	}
}