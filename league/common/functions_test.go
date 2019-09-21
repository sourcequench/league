package common

import (
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
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

func TestUpdateMatches(t *testing.T) {
	cases := []struct {
		desc    string
		matches []Match
		want    []Match
	}{
		{
			desc: "season 1",
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
			want: []Match{
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
		},
	}

	for _, c := range cases {
		got := UpdateMatches(c.matches)
		for i, match := range c.matches {
			if match.P1needs != got[i].P1needs {
				t.Errorf("desc: %s, P2needs, want: %f, got: %f", c.desc, match.P1needs, got[i].P1needs)
			}
			if match.P2needs != got[i].P2needs {
				t.Errorf("desc: %s, P2needs, want: %f, got: %f", c.desc, match.P2needs, got[i].P2needs)
			}
			if match.P1got != got[i].P1got {
				t.Errorf("desc: %s, P1got, want: %f, got: %f", c.desc, match.P1got, got[i].P1got)
			}
			if match.P2got != got[i].P2got {
				t.Errorf("desc: %s, P2got, want: %f, got: %f", c.desc, match.P2got, got[i].P2got)
			}
		}
	}
}

func TestUpdateGot(t *testing.T) {
	cases := []struct {
		desc               string
		p1n, p2n, p1g, p2g float64
		p1want, p2want     float64
	}{
		{
			desc: "Missing a game",
			p1n:  5,
			p2n:  4,
			p1g:  2,
			p2g:  3,
		}, {
			desc: "Missing a few games",
			p1n:  10,
			p2n:  7,
			p1g:  1,
			p2g:  3,
		},
	}

	for _, c := range cases {
		p1g, p2g := UpdateGot(c.p1n, c.p2n, c.p1g, c.p2g)
		if p1g != c.p1n && p2g != c.p2n {
			t.Errorf("desc: %s, got: %f, %f", c.desc, p1g, p2g)
		}
	}
}
