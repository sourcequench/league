package main

import (
	"fmt"
	"github.com/gonum/stat"
	c "github.com/sourcequench/league/common"
	p "github.com/sourcequench/league/parser"
	"log"
)

func main() {
	matches, err := p.Parse("data/ken-normalized")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}

	diffs := c.Diff(matches)
	mu, sigma := stat.MeanStdDev(diffs, nil)

	fmt.Printf("Got/Needs average match difference: %f games, Sigma: %f\n", mu, sigma)
	adjMatches := c.UpdateMatches(matches)

	diffs = c.Diff(adjMatches)
	mu, sigma = stat.MeanStdDev(diffs, nil)

	fmt.Printf("Adjusted got/needs average match difference: %f games, Sigma: %f\n", mu, sigma)
}
