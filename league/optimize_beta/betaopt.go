package main

import (
	"fmt"
	ts "github.com/mafredri/go-trueskill"
	"github.com/mafredri/go-trueskill/gaussian"
	"math"
)

type match struct {
	// Information about skill and variance of each player at match time.
	p1m, p1s, p2m, p2s float64
	// Player 1 and 2 games won
	p1g, p2g int32
}

func main() {

	matches := []match{
		{
			p1m: 101.01, p1s: 5.5, p2m: 41.24, p2s: 6.2,
			p1g: 12, p2g: 4,
		}, {
			p1m: 81.45, p1s: 4.1, p2m: 61.08, p2s: 5.2,
			p1g: 8, p2g: 6,
		}, {
			p1m: 55.5, p1s: 7.1, p2m: 45.54, p2s: 5.9,
			p1g: 7, p2g: 6,
		}, {
			p1m: 106.5, p1s: 3.2, p2m: 90.3, p2s: 5.9,
			p1g: 6, p2g: 9,
		},
	}
	beta := float64(62)
	d, b := OptimizeBeta(matches, beta)
	fmt.Printf("Optimal diff: %f, Optimal beta: %f", d, b)
}

// AggDiff takes a series of matches and returns the aggregated differences from predicted outcome.
func AggDiff(matches []match, beta float64) float64 {
	var aggdiff float64
	// Read in timeseries data, all match results. More data is more better.
	for _, match := range matches {
		p1 := ts.NewPlayer(match.p1m, match.p1s)
		p2 := ts.NewPlayer(match.p2m, match.p2s)
		expected := Pwin(p1, p2, beta)
		actual := float64(
			float64(match.p1g) / float64(match.p1g+match.p2g))
		//		fmt.Printf("expected: %f, actual: %f\n", expected, actual)
		// We really only need to check a single player, as the other player is the inverse.
		diff := math.Abs(expected - actual)
		aggdiff += diff
	}
	return aggdiff
}

// Optimization function for beta
func OptimizeBeta(matches []match, beta float64) (float64, float64) {
	var betastep float64 = 0.01 // TODO: set a flag for precision for use here.
	// Determine if we need to go up, down, or we have the perfect beta
	initialDiff := AggDiff(matches, beta)
	leftDiff := AggDiff(matches, beta-betastep)
	rightDiff := AggDiff(matches, beta+betastep)

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

func Pwin(p1, p2 ts.Player, beta float64) float64 {
	// TODO: do a ts.New and set beta on that?
	deltaMu := p1.Mean() - p2.Mean()
	sumSigma := p1.Variance() + p2.Variance()
	rss := math.Sqrt(2*(beta*beta) + sumSigma)
	return gaussian.NormCdf(deltaMu / rss)
}
