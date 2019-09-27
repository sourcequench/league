package main

// go run npl_diff.go -v=2 -logtostderr=true

import (
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
	matches, errs := p.Parse("data/this-time-for-sure.csv", nil)
	if errs != nil {
		log.Fatalf("could not parse file: %v", errs)
	}

	skills := make(map[string]float64)

	for _, match := range matches {
		skills[match.P1name] = match.P1skill
		skills[match.P2name] = match.P2skill
	}
	// Show the final skill ratings
	fmt.Println("FINAL SKILLS:")
	for player, skill := range skills {
		fmt.Printf("%s: %f\n", player, skill)
	}

	diffs := c.Diff(matches)
	mu, sigma := stat.MeanStdDev(diffs, nil)

	fmt.Printf("Got/Needs average match difference: %f games, Sigma: %f\n", mu, sigma)
	s := npl.ThreeTwoOne{}
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
	for user, diffs := range userDiffs {
		mu, sigma = stat.MeanStdDev(diffs, nil)
		fmt.Printf("%s: mean: %f, sigma: %f\n", user, mu, sigma)
	}
	glog.Flush()

}
