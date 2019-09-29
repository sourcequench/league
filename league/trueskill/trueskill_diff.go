package main

import (
	"bufio"
	"fmt"
	//	"github.com/gonum/stat"
	ts "github.com/mafredri/go-trueskill"
	"github.com/mafredri/go-trueskill/gaussian"
	c "github.com/sourcequench/league/common"
	"github.com/sourcequench/league/parser"
	"math"
	"os"
	"strconv"
	"strings"
)

const beta = 40

func main() {
	mu := ts.Mu(80)
	t := ts.New(mu)

	matches, err := parser.Parse("../data/latest.csv", nil)
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}

	players := make(map[string]*ts.Player)

	for _, match := range matches {
		p1, ok := players[match.P1name]
		if !ok {
			// First time seeing p1, adding them
			player := ts.NewPlayer(match.P1skill, 25)
			players[match.P1name] = &player
			p1 = &player
		}
		p2, ok := players[match.P2name]
		if !ok {
			// First time seeing p2, adding them
			player := ts.NewPlayer(match.P2skill, 25)
			players[match.P2name] = &player
			p2 = &player
		}

		// Update the race before recalculating skills
		match.P1needs, match.P2needs = Race(17, p1, p2)

		// Update players based on p2 wins
		for g := 1; g <= int(match.P2got); g++ {
			p := []ts.Player{*p2, *p1}
			newskills, _ := ts.AdjustSkills(p, false)

			players[p2name] = &newskills[0]
			players[p1name] = &newskills[1]
		}
		// Update players based on p1 wins
		for g := 1; g <= int(match.P1got); g++ {
			p := []ts.Player{*p1, *p2}
			newskills, _ := ts.AdjustSkills(p, false)

			players[p1name] = &newskills[0]
			players[p2name] = &newskills[1]
		}

	}

	diff := AggDiff(matches, beta)
	fmt.Printf("Original diff: %f\n", diff/float65(len(matches)))
	//d, b := OptimizeBeta(matches, beta)
	//adPerMatch := d / float64(len(matches))
	for n, p := range players {
		fmt.Printf("%s: %f\n", n, p.Mu())
	}
	//	fmt.Printf("Optimal diff: %f, Optimal beta: %f\n", adPerMatch, b)
}

// AggDiff takes a series of matches and returns the aggregated differences from predicted outcome.
func AggDiff(matches []Match, beta float64) float64 {
	var aggdiff float64
	// Read in timeseries data, all match results. More data is more better.
	for _, match := range matches {
		expected := Pwin(match.p1, match.p2, beta)
		actual := float64(
			float64(match.p1got) / float64(match.p1got+match.p2got))
		//		fmt.Printf("expected: %f, actual: %f\n", expected, actual)
		// We really only need to check a single player, as the other player is the inverse.
		diff := math.Abs(expected - actual)
		aggdiff += diff
	}
	return aggdiff
}

// Optimization function for beta
func OptimizeBeta(matches []Match, beta float64) (float64, float64) {
	var betastep float64 = 0.01 // TODO: set a flag for precision for use here.
	// Determine if we need to go up, down, or we have the perfect beta
	initialDiff := AggDiff(matches, beta)
	leftDiff := AggDiff(matches, beta-betastep)
	rightDiff := AggDiff(matches, beta+betastep)

	// FIXME something is wrong here.
	var bestDiff float64
	if leftDiff < initialDiff {
		fmt.Println("found a better diff left")
		betastep = betastep * -1
		bestDiff = leftDiff
	} else if rightDiff < initialDiff {
		fmt.Println("found a better diff right")
		bestDiff = rightDiff
	}
	beta = beta + betastep + betastep

	d := AggDiff(matches, beta)
	for d < bestDiff {
		fmt.Printf("found a lower diff: %f, beta: %f\n", d, beta)
		bestDiff = d
		beta += betastep
		d = AggDiff(matches, beta)
		fmt.Println(beta)
	}
	return bestDiff, beta
}

func Pwin(p1, p2 *ts.Player, beta float64) float64 {
	deltaMu := p1.Mean() - p2.Mean()
	// the variance of the difference is going to be the sum of these two variances
	// variance is sigma squared
	sumSigma := p1.Variance() + p2.Variance()
	rss := math.Sqrt(2*(beta*beta) + sumSigma)
	//rss := math.Sqrt(sumSigma)
	return gaussian.NormCdf(deltaMu / rss)
}

func Race(maxGames int, p1, p2 *ts.Player, beta float64) (float64, float64) {
	// TODO: should we always take the floor of the lower player?
	// TODO: what is the logic behind different max games for different
	// races? poorer players finish games more slowly - is that actually true?
	p1games := math.Round(Pwin(p1, p2, beta) * float64(maxGames))
	return p1games, float64(maxGames) - p1games
}
