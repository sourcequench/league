package main

// go run npl_diff.go -v=2 -logtostderr=true

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/gonum/stat"
	c "github.com/sourcequench/league/common"
	"github.com/sourcequench/league/npl"
	p "github.com/sourcequench/league/parser"
	"log"
)

func init() {
	flag.Parse()
}

func main() {
	matches, errs := p.Parse("data/latest.csv", nil)
	if errs != nil {
		log.Fatalf("could not parse file: %v", errs)
	}

	// Show the final skill ratings
	fmt.Println("ENDING SKILLS:")
	finalMap := c.FinalSkill(matches)
	for player, skill := range finalMap {
		fmt.Printf("%s: %f\n", player, skill)
	}

	diffs := c.Diff(matches)
	mu, sigma := stat.MeanStdDev(diffs, nil)

	fmt.Printf("Got/Needs average match difference: %f games, Sigma: %f\n", mu, sigma)
	s := npl.Two{}
	adjMatches := c.UpdateMatches(matches, s)
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
	normalMap := c.PlayerNormal(userDiffs)
	for user, stats := range normalMap {
		fmt.Printf("%s: mean: %f, sigma: %f\n", user, stats[0], stats[1])
	}
	glog.Flush()

	// Look for result swaps.
	for i, match := range matches {
		if match.P1needs == match.P1got && adjMatches[i].P1needs != adjMatches[i].P1got {
			o, _ := json.MarshalIndent(match, "", "  ")
			fmt.Printf("ORIGINAL MATCH: %s\n\n", string(o))
			a, _ := json.MarshalIndent(adjMatches[i], "", "  ")
			fmt.Printf("ADJUSTED MATCH: %s\n\n", string(a))
		}
	}

}
