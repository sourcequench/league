package optimize

import (
	"github.com/gonum/stat"
	c "github.com/sourcequench/league/common"
	//	"github.com/sourcequench/league/interfaces"
)

// Optimization function for beta
func OptimizeBeta(matches []c.Match, beta float64) (float64, float64, []float64) {
	var betastep float64 = 0.01 // TODO: set a flag for precision for use here.
	// Determine if we need to go up, down, or we have the perfect beta
	id := c.BetaDiff(matches, beta)
	ld := c.BetaDiff(matches, beta-betastep)
	rd := c.BetaDiff(matches, beta+betastep)

	initialDiff, _ := stat.MeanStdDev(id, nil)
	leftDiff, _ := stat.MeanStdDev(ld, nil)
	rightDiff, _ := stat.MeanStdDev(rd, nil)

	var bestDiff float64

	if leftDiff < initialDiff {
		betastep = betastep * -1
		bestDiff = leftDiff
	} else if rightDiff < initialDiff {
		bestDiff = rightDiff
	}
	beta = beta + betastep + betastep

	// TODO: I calculate diffs elsewhere - think about this further.
	diffList := c.BetaDiff(matches, beta)
	d, _ := stat.MeanStdDev(diffList, nil)
	for d < bestDiff {
		bestDiff = d
		beta += betastep
		diffList = c.BetaDiff(matches, beta)
		d, _ = stat.MeanStdDev(diffList, nil)
	}
	return bestDiff, beta, diffList
}

/*
// A generic optimization function for a set of matches.
// A bit ugly here, but we take in matches, then a skill adjustment interface,
// then a function which takes in p1skill, p2skill, p1got, p2got and returns a
// calculated race.
// From there we tweak the factor until we get an optimal result.
func Optimize(matches []c.Match, factor float64, s interfaces.Skill n func(float64, float64, float64, float64)) (float64, float64) {
	betastep := 0.1
	// Determine if we need to go up, down, or we have the perfect beta
	id := c.Diff(matches)
	// slighty smaller.
	// TODO: change the interface of updatematches to take function n from arguments.
	// man, ok this requires some big refactoring.
	adjMatches := c.UpdateMatches(matches, s)
	ld := c.Diff(matches)
	rd := c.Diff(matches)

	initialDiff, _ := stat.MeanStdDev(id, nil)
	leftDiff, _ := stat.MeanStdDev(ld, nil)
	rightDiff, _ := stat.MeanStdDev(rd, nil)

	var bestDiff float64

	if leftDiff < initialDiff {
		betastep = betastep * -1
		bestDiff = leftDiff
	} else if rightDiff < initialDiff {
		bestDiff = rightDiff
	}
	beta = beta + betastep + betastep

	// TODO: I calculate diffs elsewhere - think about this further.
	diffList := c.BetaDiff(matches, beta)
	d, _ := stat.MeanStdDev(diffList, nil)
	for d < bestDiff {
		bestDiff = d
		beta += betastep
		diffList = c.BetaDiff(matches, beta)
		d, _ = stat.MeanStdDev(diffList, nil)
	}
	return bestDiff, beta, diffList
}
*/
