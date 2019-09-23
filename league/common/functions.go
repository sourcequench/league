package common

import (
	"fmt"
	"github.com/dylrich/rating/glicko2"
	"github.com/mafredri/go-trueskill/gaussian"
	"github.com/sourcequench/league/npl"
	"math"
	"math/rand"
)

// Diff takes a series of matches and returns the aggregated differences from predicted outcome.
func Diff(matches []Match) []float64 {
	// Read in timeseries data, all match results. More data is more better.
	var diffs []float64
	for _, match := range matches {
		d1 := math.Abs(match.P1needs - match.P1got)
		d2 := math.Abs(match.P2needs - match.P2got)
		diffs = append(diffs, d2)
		diffs = append(diffs, d1)
	}
	return diffs
}

// BetaDiff takes a series of matches and returns the aggregated differences from predicted outcome.
func BetaDiff(matches []Match, beta float64) []float64 {
	var diffs []float64
	players := make(map[string]*glicko2.Player)
	params := glicko2.Parameters{
		InitialDeviation:  27,
		InitialVolatility: .06,
	}
	for _, match := range matches {
		// Add players as we discover them.
		p1, ok := players[match.P1name]
		if !ok {
			params.InitialRating = match.P1skill
			players[match.P1name] = glicko2.NewPlayer(params)
			p1 = players[match.P1name]
		}
		p2, ok := players[match.P2name]
		if !ok {
			params.InitialRating = match.P2skill
			players[match.P2name] = glicko2.NewPlayer(params)
			p2 = players[match.P2name]
		}

		expected := Pwin(p1, p2, beta)
		actual := float64(
			float64(match.P1got) / float64(match.P1got+match.P2got))
		//              fmt.Printf("expected: %f, actual: %f\n", expected, actual)
		// We really only need to check a single player, as the other player is the inverse.
		diff := math.Abs(expected - actual)
		diffs = append(diffs, diff)
	}

	return diffs
}

func Pwin(p1, p2 *glicko2.Player, beta float64) float64 {
	// TODO: do a ts.New and set beta on that?
	deltaMu := p1.Rating - p2.Rating
	sumSigma := math.Pow(p1.Deviation, 2) + math.Pow(p1.Deviation, 2)
	//      sumSigma := p1.Deviation + p1.Deviation
	rss := math.Sqrt(2*(beta*beta) + sumSigma)
	//      rss := math.Sqrt(sumSigma)
	return gaussian.NormCdf(deltaMu / rss)
}

// AdjustedMatches updates historic matchdata by making adjustments to the
// skill updating mechainsm.
func UpdateMatches(matches []Match) []Match {
	// Structure outside the match loop for tracking our adjusted skills.
	skills := make(map[string]float64)

	var adjMatches []Match
	for _, match := range matches {
		adjMatch := Match{}
		// Look up the latest skills - or if seeing the player for the first time add
		// their skill based on the skill in this first match.
		p1skill, ok := skills[match.P1name]
		adjMatch.P1skill = p1skill
		if !ok {
			skills[match.P1name] = match.P1skill
			adjMatch.P1skill = match.P1skill

		}

		p2skill, ok := skills[match.P2name]
		adjMatch.P2skill = p2skill
		if !ok {
			skills[match.P2name] = match.P2skill
			adjMatch.P2skill = match.P2skill
		}

		// Look up the race in the chart. Higher skill player has the higher game need.
		p1higher := match.HigherPlayer(adjMatch.P1skill, adjMatch.P2skill)

		if p1higher {
			// new needs based on the latest skill
			adjMatch.P1needs, adjMatch.P2needs = npl.NplRace(adjMatch.P1skill, adjMatch.P2skill)
			// Model a new "got" games, if historic data can't determine the winner.
			adjMatch.P1got, adjMatch.P2got = UpdateGot(adjMatch.P1needs, adjMatch.P2needs, match.P1got, match.P2got)
		} else {
			// new needs based on the latest skill
			adjMatch.P2needs, adjMatch.P1needs = npl.NplRace(adjMatch.P1skill, adjMatch.P2skill)
			// Model a new "got" games, if historic data can't determine the winner.
			adjMatch.P1got, adjMatch.P2got = UpdateGot(adjMatch.P1needs, adjMatch.P2needs, match.P1got, match.P2got)
		}

		/* FINDING BAD RACE CALCS IN DATA - MOVE THIS

		n1, n2 := npl.NplRace(match.P1skill, match.P2skill)
		p1higher = match.HigherPlayer(match.P1skill, match.P2skill)
		if p1higher {
			if match.P1needs != n1 || match.P2needs != n2 {
				fmt.Printf("RACE MISCALC: %s, %s, %s, %f, %f\n", match.P1name, match.P2name, match.Date, match.P1skill, match.P2skill)
				fmt.Printf("played race %f:%f\n", match.P2needs, match.P1needs)
				fmt.Printf("calc'd race %f:%f\n", n1, n2)
			}
		} else {
			if match.P2needs != n1 || match.P1needs != n2 {
				fmt.Printf("RACE MISCALC: %s, %s, %s, %f, %f\n", match.P1name, match.P2name, match.Date, match.P1skill, match.P2skill)
				fmt.Printf("played race %f:%f\n", match.P2needs, match.P1needs)
				fmt.Printf("calc'd race %f:%f\n", n1, n2)
			}

		}
		*/

		maxGames := adjMatch.P1needs + adjMatch.P2needs - 1
		playedGames := match.P1got + match.P2got

		// Conditionaly adjust skills, based on who won and how close it was.
		if adjMatch.P1got == adjMatch.P1needs {
			w := skills[match.P1name]
			l := skills[match.P2name]
			skills[match.P1name], skills[match.P2name] = npl.AdjustSkills(w, l, maxGames, playedGames)
		} else {
			w := skills[match.P2name]
			l := skills[match.P1name]
			skills[match.P2name], skills[match.P1name] = npl.AdjustSkills(w, l, maxGames, playedGames)
		}
		adjMatch.P2skill, adjMatch.P1skill = skills[match.P1name], skills[match.P2name]
		/*
			fmt.Printf(
				"### %3f, %3f, %2f, %2f\n### %3f, %3f, %2f, %2f\n\n",
				match.P1skill, match.P2skill, match.P1needs, match.P2needs,
				adjMatch.P1skill, adjMatch.P2skill, adjMatch.P1needs, adjMatch.P2needs,
			)
		*/
		adjMatches = append(adjMatches, adjMatch)
	}
	for n, s := range skills {
		fmt.Printf("%s:%f\n", n, s)
	}
	return adjMatches
}

// Fabricate new "got" results if necessary on historic matches.
func UpdateGot(p1Needs, p2Needs, p1Got, p2Got float64) (float64, float64) {
	// We didn't have enough matches for the new race, use math/rand
	// to make something up proportional
	for p1Got < p1Needs && p2Got < p2Needs {
		pwin := p1Needs / (p1Needs + p2Needs)
		r := rand.Float64()
		if r < pwin {
			p1Got += 1
		} else {
			p2Got += 1
		}
	}
	return p1Got, p2Got
}
