package npl

import (
	"fmt"
	"math"
)

type Player struct {
	name  string
	skill float64
}

// Implements the interfaces.Skill interface - updating skill by 3, 2, or 1 points.
type FargoSkill struct{}

func (s FargoSkill) Update(winner, loser, maxGames, playedGames float64) (float64, float64) {
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

/*
In the Fargo system, two players who are 100 points apart are expected to win games
in a 2:1 ratio on average. 200 points and it’s 4:1, and so on.
The skill difference formula 100*log_2(p1/p2) where p1 is the
chance player 1 will win a single game and p2 is the probability that player 2
will win a game. (p1+p2=1)
*/

func Difference(p1chance, p2chance float64) (float64, error) {
	if p1chance+p2chance != 1 {
		return 0, fmt.Errorf("chances do not add to 1: %f, %f", p1chance, p2chance)
	}
	return 100 * math.Log2(p1chance/p2chance), nil
}

// TODO: implement this
func FargoRace(p1skill, p2skill float64) (float64, float64) {
	return 0, 0
}

// FitRace will return a race fitted to the percentage win chance of the players
// given their skill rating. Fitted means the closest whole number race for which
// one of the players would win given the provided "got" games.

func FitRace(p1skill, p2skill, p1got, p2got float64) (float64, float64) {
	// Get the chance of each player winning.
	higherChance := FargoPwin(p1skill, p2skill)
	lowerChance := 1 - higherChance
	played := p1got + p2got

	var p1Higher bool
	var higherGot, lowerGot float64
	// Which player is higher?
	if p1skill > p2skill {
		higherGot, lowerGot = p1got, p2got
		p1Higher = true
	} else {
		lowerGot, higherGot = p1got, p2got
	}

	higherNeeds, lowerNeeds := math.Round(higherChance*played), math.Round(lowerChance*played)
	// Starting with the number of games actually played, round to the number of games needed for player1 and player2.

	// Both players do not have enough to win, walk played matches down until we have a winner.
	for higherGot < higherNeeds && lowerGot < lowerNeeds {
		played -= 1
		higherNeeds, lowerNeeds = math.Round(higherChance*played), math.Round(lowerChance*played)
	}
	// One or more players has too many, walk needs up until we have a winner.
	for higherGot > higherNeeds || lowerGot > lowerNeeds {
		played += 1
		higherNeeds, lowerNeeds = math.Round(higherChance*played), math.Round(lowerChance*played)
	}
	// Return race needs in the same position as the player skills provided.
	if p1Higher {
		return higherNeeds, lowerNeeds
	} else {
		return lowerNeeds, higherNeeds
	}
}

// FargoPwin gives us the percentage chance that the higher skill player would win.
func FargoPwin(p1skill, p2skill float64) float64 {
	diff := math.Abs(p1skill - p2skill)
	// 100*log2(p1chance/p2chance) = difference, use algebra to solve for p1chance/p2chance
	ratio := math.Pow(math.Pow(2, diff), 1/100.0)
	// so a ratio of 2:1 is 2 out of 3 games, so 0.66666 percent chance the better player will win.
	return ratio / (1 + ratio)
}
