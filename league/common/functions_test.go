package common

import (
	"github.com/sourcequench/league/npl"
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

func TestPercentDiff(t *testing.T) {
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
			want: []float64{0.4, 1, 0.75, 1},
		},
	}

	for _, c := range cases {
		got := PercentDiff(c.matches)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("desc: %s, want: %v, got: %v", c.desc, c.want, got)
		}
	}
}

func TestWinRecord(t *testing.T) {
	cases := []struct {
		desc    string
		matches []Match
		want    map[string]float64
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
					P1name:  "Walter",
					P2name:  "Bunny",
					P1needs: 9,
					P2needs: 4,
					P1got:   9,
					P2got:   3,
					P1skill: 88,
					P2skill: 38,
				},
			},
			want: map[string]float64{"Dude": 1.0, "Walter": 0.5, "Bunny": 0},
		},
	}

	for _, c := range cases {
		got := WinRecord(c.matches)
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
					P1got:   5,
					P2got:   5,
					P1skill: 33,
					P2skill: 59,
				}, {
					Date:    "2019-02-01",
					P1name:  "Dude",
					P2name:  "Bunny",
					P1needs: 9,
					P2needs: 4,
					P1got:   9,
					P2got:   3,
					P1skill: 35,
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
					P1name:  "Dude",
					P2name:  "Bunny",
					P1needs: 9,
					P2needs: 4,
					P1got:   9,
					P2got:   3,
					P1skill: 35,
					P2skill: 38,
				},
			},
		},
	}

	for _, c := range cases {
		s := npl.Two{}
		got := UpdateMatches(c.matches, s)
		for i, match := range c.matches {
			if match.P1needs != got[i].P1needs {
				t.Errorf("desc: %s, P1needs, want: %f, got: %f", c.desc, match.P1needs, got[i].P1needs)
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
			if match.P1skill != got[i].P1skill {
				t.Errorf("desc: %s, P1skill, want: %f, got: %f", c.desc, match.P1skill, got[i].P1skill)
			}
			if match.P2skill != got[i].P2skill {
				t.Errorf("desc: %s, P2skill, want: %f, got: %f", c.desc, match.P2skill, got[i].P2skill)
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
		}, {
			desc: "Both players too many games",
			p1n:  10,
			p2n:  7,
			p1g:  11,
			p2g:  9,
		}, {
			desc: "P1 too many games",
			p1n:  10,
			p2n:  7,
			p1g:  11,
			p2g:  3,
		}, {
			desc: "P2 too many games",
			p1n:  10,
			p2n:  7,
			p1g:  1,
			p2g:  9,
		},
	}

	for _, c := range cases {
		p1g, p2g := UpdateGot(c.p1n, c.p2n, c.p1g, c.p2g)
		if p1g != c.p1n && p2g != c.p2n {
			t.Errorf("desc: %s, got: %f, %f", c.desc, p1g, p2g)
		}
	}
}

func TestStatUpdateGot(t *testing.T) {
	cases := []struct {
		desc               string
		p1n, p2n, p1g, p2g float64
		p1on, p2on         float64
		p1want, p2want     float64
	}{
		{
			desc:   "one player has too many, one not enough",
			p1n:    4,
			p2n:    10,
			p1on:   6,
			p2on:   9,
			p1g:    6,
			p2g:    6,
			p1want: 4,
			p2want: 7,
		}, {
			desc:   "same got/needs ratio",
			p1n:    6,
			p2n:    10,
			p1on:   3,
			p2on:   7,
			p1g:    3,
			p2g:    5,
			p1want: 6,
			p2want: 7,
		}, {
			desc:   "another too many, not enough",
			p1n:    6,
			p2n:    5,
			p1on:   9,
			p2on:   4,
			p1g:    8,
			p2g:    4,
			p1want: 4,
			p2want: 7,
		}, {
			desc:   "both players have too many",
			p1n:    4,
			p2n:    6,
			p1on:   6,
			p2on:   9,
			p1g:    8,
			p2g:    4,
			p1want: 4,
			p2want: 5,
		},
	}

	for _, c := range cases {
		p1g, p2g := StatUpdateGot(c.p1n, c.p2n, c.p1g, c.p2g, c.p1on, c.p2on)
		if p1g != c.p1n && p2g != c.p2n {
			t.Errorf("desc: %s, got: %f, %f", c.desc, p1g, p2g)
		}
	}
}
