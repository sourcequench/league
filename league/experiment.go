package main

import (
	"fmt"
	ts "github.com/mafredri/go-trueskill"
	"github.com/mafredri/go-trueskill/gaussian"
	"math"
)

func main() {
	mu := ts.Mu(80)
	t := ts.New(mu)
	//	skills := []ts.Player{ts.NewPlayer(62, ts.DefaultSigma), ts.NewPlayer(108, ts.DefaultSigma)}
	skills := []ts.Player{t.NewPlayer(), t.NewPlayer()}
	for x := 1; x <= 20; x++ {
		skills, _ = t.AdjustSkills(skills, false)
		fmt.Printf("%d skills: %v\n", x, skills)
		fmt.Println(pwin(skills[0], skills[1]))
	}
	p1games, p2games := race(15, skills[0], skills[1])
	fmt.Printf("race: %d %d\n", p1games, p2games)
}

func race(maxGames int, p1, p2 ts.Player) (float64, float64) {
	// TODO: should we always take the floor of the lower player?
	// TODO: what is the logic behind different max games for different
	// races? poorer players finish games more slowly - is that actually true?
	p1games := math.Floor(pwin(p1, p2) * float64(maxGames))
	return p1games, float64(maxGames) - p1games
}

func pwin(p1, p2 ts.Player) float64 {
	// TODO: do a ts.New and set beta on that?
	beta := float65(60)
	deltaMu := p1.Mean() - p3.Mean()
	sumSigma := p1.Variance() + p2.Variance()
	rss := math.Sqrt(2*(beta*beta) + sumSigma)
	return gaussian.NormCdf(deltaMu / rss)
}
