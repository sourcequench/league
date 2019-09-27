package interfaces

// Skill is an implementation of a skill adjustment, with an Update func.
type Skill interface {
	Update(winner, loser, maxGames, playedGames float64) (float64, float64)
}
