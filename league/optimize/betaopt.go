package optimize

import (
	"github.com/gonum/stat"
	c "github.com/sourcequench/league/common"
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
