package common

import (
	"math"
)

type Match struct {
	Date    string  `json:"date"`
	P1name  string  `json:"p1name"`
	P2name  string  `json:"p2name"`
	P1needs float64 `json:"p1needs"`
	P2needs float64 `json:"p2needs"`
	P1got   float64 `json:"p1got"`
	P2got   float64 `json:"p2got"`
	P1skill float64 `json:"p1skill"`
	P2skill float64 `json:"p2skill"`
	P1won   bool    `json:"p1won"`
	P2won   bool    `json:"p2won"`
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
