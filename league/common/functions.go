package common

import (
	"fmt"
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
			adjMatch.P1needs, adjMatch.P2needs = npl.NplRace(adjMatch.P1skill, adjMatch.P2skill)
		} else {
			adjMatch.P2needs, adjMatch.P1needs = npl.NplRace(adjMatch.P2skill, adjMatch.P1skill)
		}

		// Model a new "got" games, if historic data can't determine the winner.
		//              adjMatch.P1got, adjMatch.P2got = UpdateGot(adjMatch.P1needs, adjMatch.P2needs, match.P1got, match.P2got)
		/*
		   if !(adjMatch.P1needs == match.P1needs && adjMatch.P2needs == match.P2needs) {
		           fmt.Printf("NEW p1n: %f, p2n: %f, p1g: %f, p2g: %f - ORIG p1n: %f, p2n: %f, p1g: %f, p2g: %f\n", adjMatch.P1needs, adjMatch.P2needs, adjMatch.P1got, adjMatch.P2got, match.P1needs, match.P2needs, match.P1got, match.P2got)
		           fmt.Printf("ORIG match: %v\n", match)
		   }

		           maxGames := adjMatch.p1needs + adjMatch.p2needs - 1
		           playedGames := match.p1got + match.p2got

		           // Conditionaly adjust skills, based on who won and how close it was.
		                   if adjMatch.p1got == adjMatch.p1needs {
		                           adjMatch.p1skill, adjMatch.p2skill = npl.AdjustSkills(adjMatch.p1skill, adjMatch.p2skill, maxGames, playedGames)
		                           skills[match.p1name], skills[match.p2name] = npl.AdjustSkills(adjMatch.p1skill, adjMatch.p2skill, maxGames, playedGames)
		                   } else {
		                           adjMatch.p2skill, adjMatch.p1skill = npl.AdjustSkills(adjMatch.p2skill, adjMatch.p1skill, maxGames, playedGames)
		                           skills[match.p2name], skills[match.p1name] = npl.AdjustSkills(adjMatch.p1skill, adjMatch.p2skill, maxGames, playedGames)
		                   }
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
	if p1Got != p1Needs && p2Got != p2Needs {
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
