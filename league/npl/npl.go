package npl

import (
	"math"
)

// Return the race from the chart.
// The two game needs are returned in the same order as the arguments.
func NplRace(p1, p2 float64) (float64, float64) {
	highestSkill := math.Max(p1, p2)
	diffSkill := math.Abs(p1 - p2)
	var higher, lower float64
	switch {
	case highestSkill <= 59:
		switch {
		case diffSkill <= 8:
			higher, lower = 6, 6
		case diffSkill <= 17:
			higher, lower = 6, 5
		case diffSkill <= 27:
			higher, lower = 7, 5
		case diffSkill <= 37:
			higher, lower = 7, 4
		case diffSkill <= 48:
			higher, lower = 7, 3
		case diffSkill <= 59:
			higher, lower = 8, 3
		case diffSkill > 59:
			higher, lower = 9, 3
		}
	case highestSkill >= 60 && highestSkill <= 89:
		switch {
		case diffSkill <= 7:
			higher, lower = 7, 7
		case diffSkill <= 15:
			higher, lower = 7, 6
		case diffSkill <= 23:
			higher, lower = 7, 5
		case diffSkill <= 32:
			higher, lower = 8, 5
		case diffSkill <= 41:
			higher, lower = 8, 4
		case diffSkill <= 50:
			higher, lower = 9, 4
		case diffSkill <= 59:
			higher, lower = 9, 3
		case diffSkill <= 69:
			higher, lower = 10, 3
		case diffSkill <= 79:
			higher, lower = 11, 3
		case diffSkill >= 80:
			higher, lower = 12, 3
		}
	case highestSkill >= 90:
		switch {
		case diffSkill <= 5:
			higher, lower = 8, 8
		case diffSkill <= 11:
			higher, lower = 8, 7
		case diffSkill <= 17:
			higher, lower = 8, 6
		case diffSkill <= 23:
			higher, lower = 9, 6
		case diffSkill <= 30:
			higher, lower = 9, 5
		case diffSkill <= 37:
			higher, lower = 10, 5
		case diffSkill <= 44:
			higher, lower = 10, 4
		case diffSkill <= 52:
			higher, lower = 11, 4
		case diffSkill <= 60:
			higher, lower = 12, 4
		case diffSkill <= 68:
			higher, lower = 12, 3
		case diffSkill <= 77:
			higher, lower = 13, 3
		case diffSkill <= 86:
			higher, lower = 14, 3
		case diffSkill <= 95:
			higher, lower = 15, 3
		case diffSkill >= 96:
			higher, lower = 17, 3
		}
	default:
		return 0, 0
	}
	// Cases above always have higher on the left, swap if necessary.
	if p2 > p1 {
		lower, higher = higher, lower
	}
	return higher, lower
}

func Pwin(p1, p2 float64) float64 {
	// TODO: which number of games goes with which player
	g1, g2 := NplRace(p1, p2)
	return g1 / (g1 + g2)
}

type Player struct {
	name  string
	skill float64
}

// Update and higher, lower = skills based on how close the game was.
func AdjustSkills(winner, loser, maxGames, playedGames float64) (float64, float64) {
	percent := playedGames / maxGames
	switch {
	case percent <= 0.75:
		winner += 2
		loser -= 2
	case percent <= 0.9:
		winner += 1
		loser -= 1
	default:
		winner += 0
		loser -= 0
	}
	return winner, loser
}
