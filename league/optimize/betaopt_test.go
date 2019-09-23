package optimize

import (
	c "github.com/sourcequench/league/common"
	"reflect"
	"testing"
)

func TestOptimizeBeta(t *testing.T) {
	cases := []struct {
		desc               string
		matches            []c.Match
		beta               float64
		wantDiff, wantBeta float64
		wantDiffList       []float64
	}{
		{
			desc: "test1",
			matches: []c.Match{
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
			beta:     80,
			wantDiff: 0.0703378894240613,
			wantBeta: 44.91999999999562,
			wantDiffList: []float64{
				0.14065160026510037,
				3.239010097333406e-05,
			},
		}, {
			desc: "test2",
			matches: []c.Match{
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
			beta:     80,
			wantDiff: 0.0703378894240613,
			wantBeta: 44.91999999999562,
			wantDiffList: []float64{
				0.14065160026510037,
				3.239010097333406e-05,
			},
		},
	}

	for _, c := range cases {
		got1, got2, got3 := OptimizeBeta(c.matches, c.beta)
		if got1 != c.wantDiff {
			t.Errorf("desc: %s, got1 %v does not match wantDiff %v", c.desc, got1, c.wantDiff)
		}
		if got2 != c.wantBeta {
			t.Errorf("desc: %s, got2 %v does not match wantBeta %v", c.desc, got2, c.wantBeta)
		}
		if !reflect.DeepEqual(got3, c.wantDiffList) {
			t.Errorf("desc: %s, got3 %v does not match wantDiffList %v", c.desc, got3, c.wantDiffList)
		}
	}
}
