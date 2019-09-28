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
	g1, g2 := NplRace(p1, p2)
	return g1 / (g1 + g2)
}

type Player struct {
	name  string
	skill float64
}

// Implements the interfaces.Skill interface - updating skill by 3, 2, or 1 points.
type ThreeTwoOne struct{}

func (s ThreeTwoOne) Update(winner, loser, maxGames, playedGames float64) (float64, float64) {
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

// Implements the interfaces.Skill interface - updating skill by 2, 1, or 0 points.
type TwoOneZero struct{}

func (s TwoOneZero) Update(winner, loser, maxGames, playedGames float64) (float64, float64) {
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

// Implements the interfaces.Skill interface - updating based on the original +-2.
type Two struct{}

func (s Two) Update(winner, loser, maxGames, playedGames float64) (float64, float64) {
	winner += 2
	loser -= 2
	return winner, loser
}

// Implements the interfaces.Skill interface - doing nothing for control
type NoChange struct{}

func (s NoChange) Update(winner, loser, maxGames, playedGames float64) (float64, float64) {
	return winner, loser
}

/*
In the NPL system, two players who are 30 points apart are expected to win games
in a 2:1 ratio on average. 60 points and it’s 4:1, and so on. This is identical
to the Fargo probability scale except that the two systems scales are different
by a ratio of 30:100. In Fargo, if two players are 100 points apart they can be
expected to win in a 2:1 ratio. In the NPL the points difference is given by the
formula 100*log_10(p1/p2) and in Fargo it is 100*log_2(p1/p2) where p1 is the
chance player 1 will win a single game and p2 is the probability that player 2
will win a game. (p1+p2=1)
*/

func Difference(p1chance, p2chance float64) float64 {
	return 100 * math.Log10(p1chance/p2chance)
}

// NplPwin gives us the percentage chance that the higher skill player would win.
func NplPwin(p1skill, p2skill float64) float64 {
	diff := math.Abs(p1skill - p2skill)
	// At a 30 point difference the higher player should win in a ratio of 2:1
	// also known as 0.666... 30 more points is 4:1, or .8
	// test case, difference of 0 should return .5

	// How many 30 point differences do we have?
	thirties := math.Floor(diff / 30)
	numerator := math.Pow(2, thirties)
	denominator := numerator + 1
	basePwin := numerator / denominator

	// How much do we adjust percentage win chance in between 30 point differences?
	// Split up the difference between 0.5 and 0.6666... into 30 steps.
	step := (2/3.0 - 0.5) / 30
	// modulo gets us the remainder for proportional adjustment
	mod := int(diff) % 30
	pwin := basePwin + (float64(mod) * step)
	return pwin
}
