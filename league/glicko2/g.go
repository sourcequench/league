package main

import (
	"fmt"
	"github.com/dylrich/rating/glicko2"
)

func main() {
	params := glicko2.Parameters{
		InitialRating:     80.0,
		InitialDeviation:  8.3,
		InitialVolatility: .006,
	}

	p1 := glicko2.NewPlayer(params)
	p2 := glicko2.NewPlayer(params)

	p1Rating := p1.Rating
	//	p1Deviation := p1.Deviation
	p2Rating := p2.Rating
	p2Deviation := p2.Deviation

	p1Outcome := p1.Win(p2Rating, p2Deviation)
	p2Outcome := p2.Lose(p1Rating, p2Deviation)

	fmt.Printf("Player 1's rating is now %v (%v) with a deviation of %v (%v) and volatility of %v (%v)\n", p1Outcome.Rating, p1Outcome.RatingDelta, p1Outcome.Deviation, p1Outcome.DeviationDelta, p1Outcome.Volatility, p1Outcome.VolatilityDelta)
	fmt.Printf("Player 2's rating is now %v (%v) with a deviation of %v (%v) and volatility of %v (%v)\n", p2Outcome.Rating, p2Outcome.RatingDelta, p2Outcome.Deviation, p2Outcome.DeviationDelta, p2Outcome.Volatility, p2Outcome.VolatilityDelta)
}
