package common

import (
	"math"
)

type Match struct {
	Date, P1name, P2name                             string
	P1needs, P1got, P2needs, P2got, P1skill, P2skill float64
	P1won, P2won                                     bool
}

// Race charts give a race, so we need to know which player is higher to know
// which number of games goes with which player. True if p1 is higher. Equal
// skills won't matter as the race will be the same.
func (p *Match) HigherPlayer(p1skill, p2skill float64) bool {
	higher := math.Max(p2skill, p1skill)
	if higher == p1skill {
		return true
	} else {
		return false
	}
}
