package main

import (
	"fmt"
	"github.com/gonum/stat"
	c "github.com/sourcequench/league/common"
	p "github.com/sourcequench/league/parser"
	"log"
)

func main() {
	matches, err := p.Parse("data/nobobby")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}

	skills := make(map[string]float64)

	for _, match := range matches {
		skills[match.P1name] = match.P1skill
		skills[match.P2name] = match.P2skill
	}
	/*
		// Show the final skill ratings
		for player, skill := range skills {
			fmt.Printf("%s: %f\n", player, skill)
		}
	*/

	diffs := c.Diff(matches)
	mu, sigma := stat.MeanStdDev(diffs, nil)

	fmt.Printf("Got/Needs average match difference: %f games, Sigma: %f\n", mu, sigma)
	adjMatches := c.UpdateMatches(matches)
	/*
		for i, _ := range matches {
			fmt.Printf("%v\n", matches[i])
			fmt.Printf("%v\n\n", adjMatches[i])
		}
	*/
	diffs = c.PercentDiff(adjMatches)
	mu, sigma = stat.MeanStdDev(diffs, nil)

	fmt.Printf("Adjusted got/needs average match difference: %f games, Sigma: %f\n", mu, sigma)

	userDiffs := c.PerUserPercentDiff(adjMatches)
	for user, diffs := range userDiffs {
		mu, sigma = stat.MeanStdDev(diffs, nil)
		fmt.Printf("%s: mean: %f, sigma: %f\n", user, mu, sigma)
	}

}
