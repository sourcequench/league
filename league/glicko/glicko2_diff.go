package main

import (
	"fmt"
	"github.com/dylrich/rating/glicko2"
	"github.com/gonum/stat"
	"github.com/mafredri/go-trueskill/gaussian"
	"github.com/sourcequench/league/optimize"
	"github.com/sourcequench/league/parser"
	"log"
	"math"
)

func main() {
	glicko2.SystemConstant = 0.2

	matches, err := parser.Parse("../data/ken-normalized")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	// Set the glicko2 params
	params := glicko2.Parameters{
		InitialDeviation:  27,
		InitialVolatility: .06,
	}

	players := make(map[string]*glicko2.Player)

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
	beta := float64(82.5)
	_, b, dl := optimize.OptimizeBeta(matches, beta)
	mu, sigma := stat.MeanStdDev(dl, nil)

	fmt.Printf("Optimal diff: %f, Optimal beta: %f, Sigma: %f\n", mu, b, sigma)
	player1, player2 := "WinstonW", "KenB"
	fmt.Printf("Percent %s beating %s: %f\n", player1, player2, Pwin(players["BobbyS"], players["KevinH"], b))
	n1, n2 := race(30, players[player1], players[player2], b)
	fmt.Printf("Race: %f %f\n", n1, n2)
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

func race(maxGames int, p1, p2 *glicko2.Player, b float64) (float64, float64) {
	// TODO: should we always take the floor of the lower player?
	// TODO: what is the logic behind different max games for different
	// races? poorer players finish games more slowly - is that actually true?
	p1games := math.Round(Pwin(p1, p2, b) * float64(maxGames))
	return p1games, float64(maxGames) - p1games
}
