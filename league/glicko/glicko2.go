package glicko

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
