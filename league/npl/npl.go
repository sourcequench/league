package npl

import (
	"math"
)

func NplRace(p1, p2 float64) (float64, float64) {
	highestSkill := math.Max(p1, p2)
	diffSkill := math.Abs(p1 - p2)
	switch {
	case highestSkill <= 59:
		switch {
		case diffSkill <= 8:
			return 6, 6
		case diffSkill <= 17:
			return 6, 5
		case diffSkill <= 27:
			return 7, 5
		case diffSkill <= 37:
			return 7, 4
		case diffSkill <= 48:
			return 7, 3
		case diffSkill <= 59:
			return 8, 3
		case diffSkill > 59:
			return 9, 3
		}
	case highestSkill >= 60 && highestSkill <= 89:
		switch {
		case diffSkill <= 7:
			return 7, 7
		case diffSkill <= 15:
			return 7, 6
		case diffSkill <= 23:
			return 7, 5
		case diffSkill <= 32:
			return 8, 5
		case diffSkill <= 41:
			return 8, 4
		case diffSkill <= 50:
			return 9, 4
		case diffSkill <= 59:
			return 9, 3
		case diffSkill <= 69:
			return 10, 3
		case diffSkill <= 79:
			return 11, 3
		case diffSkill >= 80:
			return 12, 3
		}
	case highestSkill >= 90:
		switch {
		case diffSkill <= 5:
			return 8, 8
		case diffSkill <= 11:
			return 8, 7
		case diffSkill <= 17:
			return 8, 6
		case diffSkill <= 23:
			return 9, 6
		case diffSkill <= 30:
			return 9, 5
		case diffSkill <= 37:
			return 10, 5
		case diffSkill <= 44:
			return 10, 4
		case diffSkill <= 52:
			return 11, 4
		case diffSkill <= 60:
			return 12, 4
		case diffSkill <= 68:
			return 12, 4
		case diffSkill <= 77:
			return 12, 3
		case diffSkill <= 86:
			return 13, 3
		case diffSkill <= 95:
			return 14, 3
		case diffSkill >= 96:
			return 17, 3
		}
	}
	return 0, 0
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

// Update and return skills based on how close the game was.
func AdjustSkills(winner, loser, maxGames, playedGames float64) (float64, float64) {
	return winner, loser
	percent := playedGames / maxGames
	switch {
	case percent <= 0.75:
		winner += 3
		loser -= 3
	case percent <= 0.9:
		winner += 2
		loser -= 2
	default:
		winner += 1
		loser -= 1
	}
	return winner, loser
}
