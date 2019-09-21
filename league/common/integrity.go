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
