package common

import (
	"reflect"
	"testing"
)

func TestAggDiff(t *testing.T) {
	cases := []struct {
		desc    string
		matches []Match
		want    []float64
	}{
		{
			desc: "season 1",
			matches: []Match{
				{
					Date:    "2019-01-01",
					P1name:  "Dude",
					P2name:  "Walter",
					P1needs: 7,
					P2needs: 5,
					P1got:   7,
					P2got:   2,
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
			want: []float64{3, 0, 1, 0},
		},
	}

	for _, c := range cases {
		got := Diff(c.matches)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("desc: %s, want: %v, got: %v", c.desc, c.want, got)
		}
	}
}
