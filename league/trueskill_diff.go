package main

import (
	"bufio"
	"fmt"
	"github.com/gonum/stat"
	ts "github.com/mafredri/go-trueskill"
	"github.com/mafredri/go-trueskill/gaussian"
	"math"
	"os"
	"strconv"
	"strings"
)

type Match struct {
	// Information about skill and variance of each player at match time.
	p1, p2 *ts.Player
	// Player 1 and 2 games won
	p1got, p2got float64
}

func main() {
	mu := ts.Mu(80)
	t := ts.New(mu)

	f, err := os.Open("data/alldata")
	if err != nil {
		fmt.Printf("shit: %v", err)
		return
	}
	defer f.Close()

	players := make(map[string]*ts.Player)

	scanner := bufio.NewScanner(f)
	var matches []Match
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")

		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("That line was wacky!: %v", fields)
			}
		}()

		p1name := fields[0]
		p2name := fields[7]

		p1skill, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			fmt.Printf("p1skill error: %v %v\n", fields, err)
		}
		p2skill, err := strconv.ParseFloat(fields[9], 64)
		if err != nil {
			fmt.Printf("p2skill error: %v %v\n", fields, err)
		}

		p1, ok := players[p1name]
		if !ok {
			// First time seeing p1, adding them
			player := ts.NewPlayer(p1skill, 8.333)
			players[p1name] = &player
			p1 = &player
		}
		p2, ok := players[p2name]
		if !ok {
			// First time seeing p2, adding them
			player := ts.NewPlayer(p2skill, 8.333)
			players[p2name] = &player
			p2 = &player
		}

		p1got, err := strconv.ParseFloat(fields[4], 64)
		if err != nil {
			fmt.Printf("p1got: %v\n", err)
		}
		p2got, err := strconv.ParseFloat(fields[11], 64)
		if err != nil {
			fmt.Printf("p2got: %v\n", err)
		}

		// Update players based on p2 wins
		for g := 1; g <= int(p2got); g++ {
			// FIXME - this is wrong
			p := []ts.Player{*p2, *p1}
			t.AdjustSkills(p, false)
			adjusted, _ := t.AdjustSkills(p, false)
			players[p2name] = &adjusted[0]
			players[p1name] = &adjusted[1]
		}
		// Update players based on p1 wins
		for g := 1; g <= int(p1got); g++ {
			p := []ts.Player{*p1, *p2}
			adjusted, _ := t.AdjustSkills(p, false)
			players[p1name] = &adjusted[0]
			players[p2name] = &adjusted[1]
		}

		m := Match{
			p1got: p1got,
			p2got: p2got,
			p1:    p1,
			p2:    p2,
		}
		matches = append(matches, m)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("shit2: %v", err)
	}

	beta := float64(82.5)
	diff := AggDiff(matches, beta)
	fmt.Printf("Original diff: %f\n", diff/float64(len(matches)))
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
	// TODO: do a ts.New and set beta on that?
	// https://www.khanacademy.org/math/ap-statistics/random-variables-ap/combining-random-variables/v/analyzing-the-difference-in-distributions
	deltaMu := p1.Mean() - p2.Mean()
	// the variance of the difference is going to be the sum of these two variances
	// variance is sigma squared
	sumSigma := p1.Variance() + p2.Variance()
	rss := math.Sqrt(2*(beta*beta) + sumSigma)
	//rss := math.Sqrt(sumSigma)
	return gaussian.NormCdf(deltaMu / rss)
}
