package common

import "testing"

func TestNplRace(t *testing.T) {
	cases := []struct {
		desc   string
		s1, s2 float64
		want   bool
	}{
		{
			desc: "p1 is higher",
			s1:   93,
			s2:   59,
			want: true,
		}, {
			desc: "p2 is higher",
			s1:   38,
			s2:   88,
			want: false,
		}, {
			desc: "same skill",
			s1:   100,
			s2:   100,
			want: true,
		},
	}

	for _, c := range cases {
		// instantiate a Match to exercise the code.
		m := Match{P1skill: c.s1, P2skill: c.s2}
		got := m.HigherPlayer(c.s1, c.s2)
		if got != c.want {
			t.Errorf("desc: %s, want: %v, got: %v", c.desc, c.want, got)
		}
	}
}
