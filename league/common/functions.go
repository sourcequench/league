package common

import (
	"github.com/dylrich/rating/glicko2"
	"github.com/golang/glog"
	"github.com/mafredri/go-trueskill/gaussian"
	"github.com/sourcequench/league/interfaces"
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
		diff := math.Abs(expected - actual)
		diffs = append(diffs, diff)
	}

	return diffs
}

// PercentDiff takes a series of matches and returns the percent differences from predicted outcome.
func PercentDiff(matches []Match) []float64 {
	// Read in timeseries data, all match results. More data is more better.
	var diffs []float64
	for _, match := range matches {
		d1 := match.P1got / match.P1needs
		d2 := match.P2got / match.P2needs
		diffs = append(diffs, d2)
		diffs = append(diffs, d1)
	}
	return diffs
}

// PerUserPercentDiff takes a series of matches and returns the series of
// percentage differences from predicted outcome broken down by user.
func PerUserPercentDiff(matches []Match) map[string][]float64 {
	// Read in timeseries data, all match results. More data is more better.
	userDiffs := make(map[string][]float64)
	for _, match := range matches {
		d1 := match.P1got / match.P1needs
		d2 := match.P2got / match.P2needs
		userDiffs[match.P1name] = append(userDiffs[match.P1name], d1)
		userDiffs[match.P2name] = append(userDiffs[match.P2name], d2)
	}
	return userDiffs
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

// Provides a map of initial skills based on a set of matches.
func InitialSkill(matches []Match) map[string]float64 {
	skills := make(map[string]float64)
	for _, match := range matches {
		_, ok := skills[match.P1name]
		if !ok {
			skills[match.P1name] = match.P1skill
		}
		_, ok := skills[match.P2name]
		if !ok {
			skills[match.P2name] = match.P2skill
		}
	}
	return skills
}

// UpdateMatches updates historic matchdata by making adjustments to the
// skill updating mechainsm.
func UpdateMatches(matches []Match, upskill interfaces.Skill) []Match {
	// Structure outside the match loop for tracking our adjusted skills.
	skills := make(map[string]float64)

	var adjMatches []Match
	glog.V(2).Infoln("INITIAL SKILL")
	for _, match := range matches {
		adjMatch := Match{}

		// DELETE THESE TWO LINES
		adjMatch.P1needs, adjMatch.P2needs = match.P1needs, match.P2needs
		adjMatch.P1got, adjMatch.P2got = match.P1got, match.P2got

		// Copy names over to the adjusted match
		adjMatch.P1name = match.P1name
		adjMatch.P2name = match.P2name

		// Look up the latest skills - or if seeing the player for the first time add
		// their skill based on the skill in this first match.
		p1skill, ok := skills[match.P1name]
		adjMatch.P1skill = p1skill
		if !ok {
			skills[match.P1name] = match.P1skill
			glog.V(2).Infof("%s: %f\n", match.P1name, match.P1skill)
			adjMatch.P1skill = match.P1skill
			p1skill = match.P1skill
		}
		p2skill, ok := skills[match.P2name]
		if !ok {
			skills[match.P2name] = match.P2skill
			glog.V(2).Infof("%s: %f\n", match.P2name, match.P2skill)
			adjMatch.P2skill = match.P2skill
			p2skill = match.P2skill
		}
		adjMatch.P2skill = p2skill

		// Look up and adjust needs from the race chart.
		adjMatch.P1needs, adjMatch.P2needs = npl.NplRace(adjMatch.P1skill, adjMatch.P2skill)

		// Model a new "got" games, if historic data can't determine the winner.

		adjMatch.P1got, adjMatch.P2got = UpdateGot(adjMatch.P1needs, adjMatch.P2needs, match.P1got, match.P2got)

		maxGames := adjMatch.P1needs + adjMatch.P2needs - 1
		playedGames := match.P1got + match.P2got

		// Conditionaly adjust skills, based on who won and how close it was.
		if adjMatch.P1got == adjMatch.P1needs {
			w := skills[match.P1name]
			l := skills[match.P2name]
			skills[match.P1name], skills[match.P2name] = upskill.Update(w, l, maxGames, playedGames)
		} else {
			w := skills[match.P2name]
			l := skills[match.P1name]
			skills[match.P2name], skills[match.P1name] = upskill.Update(w, l, maxGames, playedGames)
		}
		adjMatch.P1skill, adjMatch.P2skill = skills[match.P1name], skills[match.P2name]
		/*
			glog.V(2).Infof(
				"### %3f, %3f, %2f, %2f\n### %3f, %3f, %2f, %2f\n\n",
				match.P1skill, match.P2skill, match.P1needs, match.P2needs,
				adjMatch.P1skill, adjMatch.P2skill, adjMatch.P1needs, adjMatch.P2needs,
			)
		*/

		adjMatches = append(adjMatches, adjMatch)
	}
	return adjMatches
}

// Fabricate new "got" results if necessary on historic matches.
// TODO: change this to take a skill update function as an argument for improved unit testing.
func UpdateGot(p1Needs, p2Needs, p1Got, p2Got float64) (float64, float64) {
	// We have too many wins for both players based on the new race, roll back wins proportionally.
	for p1Got > p1Needs && p2Got > p2Needs {
		pwin := p1Needs / (p1Needs + p2Needs)
		r := rand.Float64()
		glog.V(2).Infof("both players have too many games needs:%f|%f got:%f|%f ", p1Needs, p2Needs, p1Got, p2Got)
		if r < pwin {
			glog.V(2).Infof("p1 docked a game\n")
			p1Got -= 1
		} else {
			glog.V(2).Infof("p2 docked a game\n")
			p2Got -= 1
		}

	}
	// We didn't have enough matches for the new race, use math/rand
	// to make something up proportional
	for p1Got < p1Needs && p2Got < p2Needs {
		pwin := p1Needs / (p1Needs + p2Needs)
		r := rand.Float64()
		glog.V(2).Infof("neither player can achieve a win: needs:%f|%f got:%f|%f ", p1Needs, p2Needs, p1Got, p2Got)
		if r < pwin {
			glog.V(2).Infof("p1 given a game\n")
			p1Got += 1
		} else {
			glog.V(2).Infof("p1 given a game\n")
			p2Got += 1
		}
	}
	// Only one player had too many games - they win.
	if p1Got > p1Needs {
		glog.V(2).Infof("p1 had enough games to win: needs:%f|%f got:%f|%f \n", p1Needs, p2Needs, p1Got, p2Got)
		p1Got = p1Needs
	}
	if p2Got > p2Needs {
		glog.V(2).Infof("p2 had enough games to win: needs:%f|%f got:%f|%f \n", p1Needs, p2Needs, p1Got, p2Got)
		p2Got = p2Needs
	}

	return p1Got, p2Got
}
