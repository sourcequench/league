package main

import (
	"fmt"
	"github.com/dylrich/rating/glicko2"
	"github.com/gonum/stat"
	"github.com/mafredri/go-trueskill/gaussian"
	c "github.com/sourcequench/league/common"
	"github.com/sourcequench/league/parser"
	"log"
	"math"
)

func main() {
	glicko2.SystemConstant = 0.1
	beta := float64(41.03)

	matches, err := parser.Parse("../data/latest.csv", nil)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	// Set the glicko2 params
	params := glicko2.Parameters{
		InitialDeviation:  27,
		InitialVolatility: .06,
	}

	players := make(map[string]*glicko2.Player)

	// hang on to predicted versus actual differences
	var diffs []float64
	for _, match := range matches {
		p1, ok := players[match.P1name]
		if !ok {
			// First time seeing p1, adding them
			params.InitialRating = match.P1skill
			p := glicko2.NewPlayer(params)
			players[match.P1name] = p
			p1 = p
		}
		p2, ok := players[match.P2name]
		if !ok {
			// First time seeing p1, adding them
			params.InitialRating = match.P2skill
			p := glicko2.NewPlayer(params)
			players[match.P2name] = p
			p2 = p
		}

		// Calculate new race before updating with the results.
		match.P1needs, match.P2needs = Race(17, p1, p2, beta)
		match.P1got, match.P2got = c.UpdateGot(match.P1needs, match.P2needs, match.P1got, match.P2got)

		if match.P2needs == 0 || match.P1needs == 0 {
			fmt.Printf("shit divide by zero: %f, %f\n", match.P2needs, match.P1needs)
		}
		// Add each player's percent of got/needs
		diffs = append(diffs, match.P1got/match.P1needs)
		diffs = append(diffs, match.P2got/match.P2needs)

		// Update players based on p1 wins
		for g := 1; g <= int(match.P1got); g++ {
			p1.Win(p2.Rating, p2.Deviation)
			p2.Lose(p1.Rating, p1.Deviation)
		}
		// Update players based on p2 wins
		for g := 1; g <= int(match.P2got); g++ {
			p2.Win(p1.Rating, p1.Deviation)
			p1.Lose(p2.Rating, p2.Deviation)
		}
	}

	for name, p := range players {
		fmt.Printf("%s rating:%f deviation:%f\n", name, p.Rating, p.Deviation)
	}

	fmt.Printf("%v\n", diffs)
	mu, sigma := stat.MeanStdDev(diffs, nil)
	fmt.Printf("Mean: %f, Stddev: %f\n", mu, sigma)

}

func Pwin(p1, p2 *glicko2.Player, beta float64) float64 {
	// TODO: do a ts.New and set beta on that?
	deltaMu := p1.Rating - p2.Rating
	sumSigma := math.Pow(p1.Deviation, 2) + math.Pow(p1.Deviation, 2)
	//	sumSigma := p1.Deviation + p1.Deviation
	rss := math.Sqrt(2*(beta*beta) + sumSigma)
	//	rss := math.Sqrt(sumSigma)
	return gaussian.NormCdf(deltaMu / rss)
}

func Race(maxGames int, p1, p2 *glicko2.Player, b float64) (float64, float64) {
	// TODO: should we always take the floor of the lower player?
	// TODO: what is the logic behind different max games for different
	// races? poorer players finish games more slowly - is that actually true?
	p1games := math.Round(Pwin(p1, p2, b) * float64(maxGames))
	return p1games, float64(maxGames) - p1games
}
