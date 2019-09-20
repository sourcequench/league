package common

type Season struct {
	sdate, edate, name string
	// Player numbers
	players []string
	// Date to player name mapping
	schedule map[string][]string
}
