package common

import (
	"reflect"
	"testing"
)

func TestMiscalc(t *testing.T) {
	cases := []struct {
		desc    string
		matches []Match
		want    []Match
	}{
		{
			desc: "correct calc",
			matches: []Match{
				{
					Date:    "2019-01-01",
					P1name:  "Dude",
					P2name:  "Walter",
					P1needs: 5,
					P2needs: 7,
					P1got:   2,
					P2got:   7,
					P1skill: 33,
					P2skill: 59,
				}, {
					Date:    "2019-02-01",
					P1name:  "Donny",
					P2name:  "Bunny",
					P1needs: 9,
					P2needs: 4,
					P1got:   9,
					P2got:   3,
					P1skill: 88,
					P2skill: 38,
				},
			},
			want: []Match{},
		}, {
			desc: "miscalc",
			matches: []Match{
				{
					Date:    "2019-01-01",
					P1name:  "Dude",
					P2name:  "Walter",
					P1needs: 5,
					P2needs: 10,
					P1got:   2,
					P2got:   7,
					P1skill: 33,
					P2skill: 59,
				}, {
					Date:    "2019-02-01",
					P1name:  "Donny",
					P2name:  "Bunny",
					P1needs: 9,
					P2needs: 4,
					P1got:   9,
					P2got:   3,
					P1skill: 88,
					P2skill: 38,
				},
			},
			want: []Match{
				{
					Date:    "2019-01-01",
					P1name:  "Dude",
					P2name:  "Walter",
					P1needs: 5,
					P2needs: 10,
					P1got:   2,
					P2got:   7,
					P1skill: 33,
					P2skill: 59,
				},
			},
		},
	}

	for _, c := range cases {
		got := Miscalc(c.matches)
		if !reflect.DeepEqual(got, c.want) && (len(got) != 0 && len(c.want) != 0) {
			t.Errorf("desc: %s, got %v does not match want %v", c.desc, got, c.want)
		}
	}
}
