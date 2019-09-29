package glicko

import (
	"fmt"
	"github.com/dylrich/rating/glicko2"
	"github.com/gonum/stat"
	"github.com/mafredri/go-trueskill/gaussian"
	c "github.com/sourcequench/league/common"
	"github.com/sourcequench/league/optimize"
	"github.com/sourcequench/league/parser"
	"log"
	"math"
)

const beta = 41.03

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

// Diff takes a matche and returns the aggregated differences from predicted outcome.
func Diff(match c.Match, beta float64) []float64 {
	var diffs []float64
	players := make(map[string]*glicko2.Player)
	params := glicko2.Parameters{
		InitialDeviation:  27,
		InitialVolatility: .06,
	}
	for _, match := range matches {
		// Add players as we discover them.
		p1, ok := players[match.P1name]
		if !ok {
			params.InitialRating = match.P1skill
			players[match.P1name] = glicko2.NewPlayer(params)
			p1 = players[match.P1name]
		}
		p2, ok := players[match.P2name]
		if !ok {
			params.InitialRating = match.P2skill
			players[match.P2name] = glicko2.NewPlayer(params)
			p2 = players[match.P2name]
		}

		expected := Pwin(p1, p2, beta)
		actual := float64(
			float64(match.P1got) / float64(match.P1got+match.P2got))
		diff := math.Abs(expected - actual)
		diffs = append(diffs, diff)
	}

	return diffs
}
