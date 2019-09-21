package npl

import "testing"

func TestNplRace(t *testing.T) {
	cases := []struct {
		desc         string
		s1, s2       float64
		want1, want2 float64
	}{
		{
			desc:  "low skill section",
			s1:    33,
			s2:    59,
			want1: 7,
			want2: 5,
		}, {
			desc:  "medium skill section",
			s1:    38,
			s2:    88,
			want1: 9,
			want2: 4,
		}, {
			desc:  "medium skill, higher player s1",
			s1:    88,
			s2:    38,
			want1: 9,
			want2: 4,
		}, {
			desc:  "high skill section",
			s1:    80,
			s2:    100,
			want1: 9,
			want2: 6,
		}, {
			desc:  "103 to 94",
			s1:    80,
			s2:    100,
			want1: 9,
			want2: 6,
		}, {
			desc:  "high skill section, higher player s1",
			s1:    100,
			s2:    80,
			want1: 9,
			want2: 6,
		},
	}

	for _, c := range cases {
		g1, g2 := NplRace(c.s1, c.s2)
		if g1 != c.want1 || g2 != c.want2 {
			t.Errorf("want: %f:%f, got: %f:%f", c.want1, c.want2, g1, g2)
		}
	}
}
