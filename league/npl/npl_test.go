package npl

import (
	"math"
	"testing"
)

func TestNplRace(t *testing.T) {
	cases := []struct {
		desc         string
		s1, s2       float64
		want1, want2 float64
	}{
		{
			desc:  "low skill section, lower player first",
			s1:    33,
			s2:    59,
			want1: 5,
			want2: 7,
		}, {
			desc:  "medium skill section, lower player first",
			s1:    38,
			s2:    88,
			want1: 4,
			want2: 9,
		}, {
			desc:  "medium skill, higher player first",
			s1:    88,
			s2:    38,
			want1: 9,
			want2: 4,
		}, {
			desc:  "high skill section",
			s1:    80,
			s2:    100,
			want1: 6,
			want2: 9,
		}, {
			desc:  "103 to 94",
			s1:    80,
			s2:    100,
			want1: 6,
			want2: 9,
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

func TestTwo(t *testing.T) {
	cases := []struct {
		desc         string
		s1, s2       float64
		want1, want2 float64
	}{
		{
			desc:  "p1 wins",
			s1:    33,
			s2:    59,
			want1: 35,
			want2: 57,
		},
	}

	for _, c := range cases {
		s := Two{}
		g1, g2 := s.Update(c.s1, c.s2, 0, 0)
		if g1 != c.want1 || g2 != c.want2 {
			t.Errorf("want: %f:%f, got: %f:%f", c.want1, c.want2, g1, g2)
		}
	}
}

func TestThreeTwoOne(t *testing.T) {
	cases := []struct {
		desc                  string
		s1, s2                float64
		want1, want2          float64
		maxGames, playedGames float64
	}{
		{
			desc:        "p1 beats the piss out of p2, adjust by 3",
			s1:          100,
			s2:          100,
			maxGames:    15,
			playedGames: 9,
			want1:       103,
			want2:       97,
		}, {
			desc:        "p1 won by 3 games, adjust by 2",
			s1:          100,
			s2:          100,
			maxGames:    15,
			playedGames: 13,
			want1:       102,
			want2:       98,
		},
	}

	for _, c := range cases {
		s := ThreeTwoOne{}
		g1, g2 := s.Update(c.s1, c.s2, c.maxGames, c.playedGames)
		if g1 != c.want1 || g2 != c.want2 {
			t.Errorf("want: %f:%f, got: %f:%f", c.want1, c.want2, g1, g2)
		}
	}
}

func TestNplPwin(t *testing.T) {
	cases := []struct {
		desc   string
		s1, s2 float64
		want   float64
	}{
		{
			desc: "big skill difference, 64, lower player first",
			s1:   30,
			s2:   94,
			want: 0.822222,
		}, {
			desc: "same skill",
			s1:   50,
			s2:   50,
			want: 0.5,
		}, {
			desc: "exactly 30 points different",
			s1:   80,
			s2:   50,
			want: 2 / 3.0, // This is the same as saying 2:1 ratio
		},
	}

	const tolerance = .000001
	for _, c := range cases {
		pwin := NplPwin(c.s1, c.s2)
		if math.Abs(pwin-c.want) > tolerance {
			t.Errorf("want: %f, got: %f", c.want, pwin)
		}
	}
}
