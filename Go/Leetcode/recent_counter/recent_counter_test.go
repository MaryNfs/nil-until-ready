package recent_counter

import "testing"

func TestRecentCounter(t *testing.T) {
	cases := []struct {
		op   string
		arg  int
		want int
	}{
		{"RecentCounter", 0, 0},
		{"ping", 1, 1},
		{"ping", 100, 2},
		{"ping", 3001, 3},
		{"ping", 3002, 3},
	}

	var rc RecentCounter

	for i, c := range cases {
		switch c.op {
		case "RecentCounter":
			rc = Constructor()

		case "ping":
			got := rc.Ping(c.arg)
			if got != c.want {
				t.Errorf("step %d: Ping(%d) = %d, want %d",
					i, c.arg, got, c.want)
			}
		}
	}
}
