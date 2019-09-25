package common

import (
	"github.com/sourcequench/league/npl"
)

// Checks historic data against the race chart for mistakes.
func Miscalc(matches []Match) []Match {
	// Structure outside the match loop for tracking our adjusted skills.

	var misMatches []Match
	for _, match := range matches {
		n1, n2 := npl.NplRace(match.P1skill, match.P2skill)
		p1higher := match.HigherPlayer(match.P1skill, match.P2skill)
		if p1higher {
			if match.P1needs != n1 || match.P2needs != n2 {
				misMatches = append(misMatches, match)
			}
		} else {
			if match.P2needs != n1 || match.P1needs != n2 {
				misMatches = append(misMatches, match)
			}
		}
	}

	return misMatches
}

// Checks historic data skill ratings for odd jumps.
// Because of makeup ordering tomfoolery this may not be a useful integrity function.
func SkillMiscalc(matches []Match) []Match {
	var misMatches []Match
	skills := make(map[string]float64)
	for _, match := range matches {
		// Look up the latest skills - or if seeing the player for the first time add
		// their skill based on the skill in this first match.
		p1skill, ok := skills[match.P1name]
		if !ok {
			skills[match.P1name] = match.P1skill
			p1skill = match.P1skill
		}

		p2skill, ok := skills[match.P2name]
		if !ok {
			skills[match.P2name] = match.P2skill
			p2skill = match.P2skill
		}
		if p1skill > match.P1skill+8 ||
			p1skill < match.P1skill-8 ||
			p2skill < match.P2skill+8 ||
			p2skill < match.P2skill-8 {
			misMatches = append(misMatches, match)
		}
	}
	return misMatches
}

// Looks for cases where neither player had got=needs
func GotMistake(matches []Match) []Match {
	var mistakes []Match
	for _, match := range matches {
		if match.P1got != match.P1needs && match.P2got != match.P2got {
			mistakes = append(mistakes, match)
		}
	}
	return mistakes
}
