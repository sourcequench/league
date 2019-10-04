package npl

import (
	"fmt"
	"math"
	"testing"
)

func TestFargoPwin(t *testing.T) {
	cases := []struct {
		desc   string
		s1, s2 float64
		want   float64
	}{
		{
			desc: "lower player first",
			s1:   330,
			s2:   590,
			want: 0.858414,
		}, {
			desc: "bigger difference, lower player first",
			s1:   380,
			s2:   880,
			want: 0.969697,
		}, {
			desc: "higher player first",
			s1:   880,
			s2:   380,
			want: 0.969697,
		},
	}

	for _, c := range cases {
		tolerance := 0.0001
		got := FargoPwin(c.s1, c.s2)
		if math.Abs(got-c.want) > tolerance {
			t.Errorf("want: %f, got: %f", c.want, got)
		}
	}
}

func TestDifference(t *testing.T) {
	cases := []struct {
		desc   string
		c1, c2 float64
		want   float64
		err    error
	}{
		{
			desc: "higher player first, 60/40",
			c1:   0.60,
			c2:   0.40,
			want: 58.496250,
			err:  nil,
		}, {
			desc: "same odds, zero difference",
			c1:   0.5,
			c2:   0.5,
			want: 0,
			err:  nil,
		}, {
			desc: "25/75",
			c1:   0.25,
			c2:   0.75,
			// TODO: Should the function return negative or not?
			want: -158.496250,
			err:  nil,
		}, {
			desc: "doesn't add to 1",
			c1:   0.35,
			c2:   0.75,
			want: 9,
			err:  fmt.Errorf("an error"),
		},
	}

	for _, c := range cases {
		tolerance := 0.0001
		got, e := Difference(c.c1, c.c2)
		// For non error cases, was our result close enough to expectations?
		if math.Abs(got-c.want) > tolerance && c.err == nil {
			t.Errorf("%s: want: %f, got: %f", c.desc, c.want, got)
		}
		if e != nil && c.err == nil {
			t.Errorf("got an errror, but didn't want one %v", e)
		}
	}
}

func TestFargoRace(t *testing.T) {
	cases := []struct {
		desc         string
		s1, s2       float64
		want1, want2 float64
	}{
		{
			desc:  "placeholder",
			s1:    10,
			s2:    10,
			want1: 0,
			want2: 0,
		},
	}

	for _, c := range cases {
		g1, g2 := FargoRace(c.s1, c.s2)
		if g1 != c.want1 || g2 != c.want2 {
			t.Errorf("want: %f:%f, got: %f:%f", c.want1, c.want2, g1, g2)
		}
	}
}

func TestFargoSkill(t *testing.T) {
	cases := []struct {
		desc         string
		s1, s2       float64
		want1, want2 float64
	}{
		{
			desc:  "p1 wins",
			s1:    33,
			s2:    59,
			want1: 34,
			want2: 58,
		},
	}

	for _, c := range cases {
		s := FargoSkill{}
		g1, g2 := s.Update(c.s1, c.s2, 0, 0)
		if g1 != c.want1 || g2 != c.want2 {
			t.Errorf("want: %f:%f, got: %f:%f", c.want1, c.want2, g1, g2)
		}
	}
}

func TestFitRace(t *testing.T) {
	cases := []struct {
		desc         string
		s1, s2       float64
		g1, g2       float64
		want1, want2 float64
	}{
		{
			desc:  "difference of 100, lower player first",
			s1:    500,
			s2:    600,
			g1:    8,
			g2:    8,
			want1: 8,
			want2: 15,
		}, {
			desc:  "no difference",
			s1:    500,
			s2:    500,
			g1:    6,
			g2:    5,
			want1: 6,
			want2: 6,
		},
	}

	for _, c := range cases {
		n1, n2 := FitRace(c.s1, c.s2, c.g1, c.g2)
		if n1 != c.want1 || n2 != c.want2 {
			t.Errorf("want: %f, %f  got: %f, %f", c.want1, c.want2, n1, n2)
		}
	}
}
