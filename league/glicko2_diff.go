package main

import (
	"bufio"
	"fmt"
	"github.com/dylrich/rating/glicko2"
	"github.com/gonum/stat"
	"github.com/mafredri/go-trueskill/gaussian"
	"math"
	"os"
	"strconv"
	"strings"
)

type Match struct {
	// Information about skill and variance of each player at match time.
	p1, p2 *glicko2.Player
	// Player 1 and 2 games won
	p1got, p2got float64
}

func main() {
	// Set the glicko2 params
	params := glicko2.Parameters{
		InitialDeviation:  27,
		InitialVolatility: .06,
	}
	glicko2.SystemConstant = 0.2
	f, err := os.Open("data/alldata")
	if err != nil {
		fmt.Printf("shit: %v", err)
		return
	}
	defer f.Close()

	players := make(map[string]*glicko2.Player)

	scanner := bufio.NewScanner(f)
	var matches []Match
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")

		p1name := fields[0]
		p2name := fields[7]

		p1skill, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			fmt.Printf("ERR: p1skill: %v\n", err)
		}
		p2skill, err := strconv.ParseFloat(fields[9], 64)
		if err != nil {
			fmt.Printf("ERR: p2skill: %v\n", err)
		}
		// TODO This seems like a strange place to create players - why not put this at the end?
		p1, ok := players[p1name]
		if !ok {
			// First time seeing p1, adding them
			fmt.Printf("%s initial rating:%f\n", p1name, p1skill)
			params.InitialRating = p1skill
			p1 = glicko2.NewPlayer(params)
			players[p1name] = p1
		}
		p2, ok := players[p2name]
		if !ok {
			// First time seeing p2, adding them
			fmt.Printf("%s initial rating:%f\n", p2name, p2skill)
			params.InitialRating = p2skill
			p2 = glicko2.NewPlayer(params)
			players[p2name] = p2
		}

		p1got, err := strconv.ParseFloat(fields[4], 64)
		if err != nil {
			fmt.Printf("ERR: p1got: %v\n", err)
		}
		p2got, err := strconv.ParseFloat(fields[11], 64)
		if err != nil {
			fmt.Printf("ERR: p2got: %v\n", err)
		}

		// Update players based on p1 wins
		for g := 1; g <= int(p1got); g++ {
			p1.Win(p2.Rating, p2.Deviation)
			p2.Lose(p1.Rating, p1.Deviation)
		}
		// Update players based on p2 wins
		for g := 1; g <= int(p2got); g++ {
			p2.Win(p1.Rating, p1.Deviation)
			p1.Lose(p2.Rating, p2.Deviation)
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

	for name, p := range players {
		fmt.Printf("%s rating:%f deviation:%f\n", name, p.Rating, p.Deviation)
	}
	beta := float64(82.5)
	_, b, dl := OptimizeBeta(matches, beta)
	mu, sigma := stat.MeanStdDev(dl, nil)

	fmt.Printf("Optimal diff: %f, Optimal beta: %f, Sigma: %f\n", mu, b, sigma)
	player1, player2 := "BobbyS", "KevinH"
	fmt.Printf("Percent %s beating %s: %f\n", player1, player2, Pwin(players["BobbyS"], players["KevinH"], b))
	n1, n2 := race(40, players[player1], players[player2], b)
	fmt.Printf("Race: %f %f\n", n1, n2)
}

// AggDiff takes a series of matches and returns the aggregated differences from predicted outcome.
func AggDiff(matches []Match, beta float64) (float64, []float64) {
	var aggdiff float64
	// Read in timeseries data, all match results. More data is more better.
	var diffs []float64
	for _, match := range matches {
		expected := Pwin(match.p1, match.p2, beta)
		actual := float64(
			float64(match.p1got) / float64(match.p1got+match.p2got))
		//		fmt.Printf("expected: %f, actual: %f\n", expected, actual)
		// We really only need to check a single player, as the other player is the inverse.
		diff := math.Abs(expected - actual)
		diffs = append(diffs, diff)
		aggdiff += diff
	}
	return aggdiff, diffs
}

// Optimization function for beta
func OptimizeBeta(matches []Match, beta float64) (float64, float64, []float64) {
	var betastep float64 = 0.01 // TODO: set a flag for precision for use here.
	// Determine if we need to go up, down, or we have the perfect beta
	initialDiff, _ := AggDiff(matches, beta)
	leftDiff, _ := AggDiff(matches, beta-betastep)
	rightDiff, _ := AggDiff(matches, beta+betastep)

	var bestDiff float64
	if leftDiff < initialDiff {
		betastep = betastep * -1
		bestDiff = leftDiff
	} else if rightDiff < initialDiff {
		bestDiff = rightDiff
	}
	beta = beta + betastep + betastep

	d, diffList := AggDiff(matches, beta)
	for d < bestDiff {
		bestDiff = d
		beta += betastep
		d, diffList = AggDiff(matches, beta)
	}
	return bestDiff, beta, diffList
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
