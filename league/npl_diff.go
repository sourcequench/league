package main

import (
	"bufio"
	"fmt"
	"github.com/gonum/stat"
	"github.com/sourcequench/league/npl"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Match struct {
	date, p1name, p2name                             string
	p1needs, p1got, p2needs, p2got, p1skill, p2skill float64
	p1won, p2won                                     bool
}

func main() {
	// open, read, and parse the data into a match structure
	// open data file
	// read the bytes
	// loop over, reading each line
	// strings.Spit each line into a slice of strings
	// extract the fields we need, for now we just need:
	//   date, p1needs, p1got, p2needs, p2got, p1skill, p2skill
	// TODO: sort by date so date entry mistakes don't fuck us up
	// calculate a float of the expected percent win of p1 (g1 / (g1 + g2))
	//
	//
	f, err := os.Open("data/alldata")
	if err != nil {
		fmt.Printf("shit: %v", err)
		return
	}
	defer f.Close()

	// player to skill mapping
	skills := make(map[string]*float64)

	scanner := bufio.NewScanner(f)
	var matches []Match
	var adjMatches []Match
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		p1needs, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			fmt.Printf("p1needs: %v\n", err)
			fmt.Println(fields)
		}
		p1got, err := strconv.ParseFloat(fields[4], 64)
		if err != nil {
			fmt.Printf("error getting p1got: %v\n", fields)
		}
		p2needs, err := strconv.ParseFloat(fields[10], 64)
		if err != nil {
			fmt.Printf("error getting p2needs: %v\n", err)
		}
		p2got, err := strconv.ParseFloat(fields[11], 64)
		if err != nil {
			fmt.Printf("error getting p2got: %v\n", fields)
		}
		p1skill, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			fmt.Printf("error getting p1skill: %v\n", fields)
		}
		p2skill, err := strconv.ParseFloat(fields[9], 64)
		if err != nil {
			fmt.Printf("error getting p2skill: %v\n", fields)
		}

		p1name := fields[0]
		p2name := fields[7]

		skills[p1name] = &p1skill
		skills[p2name] = &p2skill

		m := Match{
			p1name:  p1name,
			p2name:  p2name,
			date:    fields[6],
			p1needs: p1needs,
			p1got:   p1got,
			p2needs: p2needs,
			p2got:   p2got,
			p1skill: p1skill,
			p2skill: p2skill,
		}
		matches = append(matches, m)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("shit2: %v", err)
	}
	diffs := AggDiff(matches)
	mu, sigma := stat.MeanStdDev(diffs, nil)

	fmt.Printf("Got/Needs average match difference: %f games, Sigma: %f\n", mu, sigma)

	for _, match := range matches {
		p1skill := skills[match.p1name]
		p2skill := skills[match.p2name]

		higher := math.Max(*p2skill, *p1skill)
		// Look up the race from the chart.
		// TODO this function always returns in higher, lower order - so we have to figure out which is which, yuck.
		if higher == match.p1skill {
			match.p1needs, match.p2needs = npl.NplRace(*skills[match.p1name], *skills[match.p2name])
		} else {
			match.p2needs, match.p1needs = npl.NplRace(*skills[match.p1name], *skills[match.p2name])
		}
		// We didn't have enough matches for the new race, use math/rand to make something up proportional
		if match.p1got != match.p1needs && match.p2got != match.p2needs {
			pwin := match.p1needs / (match.p1needs + match.p2needs)
			r := rand.Float64()
			if r < pwin {
				match.p1got += 1
			}
		}

		maxGames := match.p1needs + match.p2needs - 1
		playedGames := match.p1got + match.p2got
		if match.p1needs == match.p1got {
			npl.AdjustSkills(skills[match.p1name], skills[match.p2name], maxGames, playedGames)
		} else {
			npl.AdjustSkills(skills[match.p2name], skills[match.p1name], maxGames, playedGames)
		}
		adjMatches = append(adjMatches, match)
	}
	for name, skill := range skills {
		fmt.Printf("%s: %f\n", name, *skill)
	}
	diffs = AggDiff(adjMatches)
	mu, sigma = stat.MeanStdDev(diffs, nil)

	fmt.Printf("Adjusted got/needs average match difference: %f games, Sigma: %f\n", mu, sigma)
}

// AggDiff takes a series of matches and returns the aggregated differences from predicted outcome.
func AggDiff(matches []Match) []float64 {
	// Read in timeseries data, all match results. More data is more better.
	var diffs []float64
	for _, match := range matches {
		d1 := math.Abs(match.p1needs - match.p1got)
		d2 := math.Abs(match.p2needs - match.p2got)
		diffs = append(diffs, d1)
		diffs = append(diffs, d2)
	}
	return diffs
}
