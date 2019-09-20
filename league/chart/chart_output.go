package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Match struct {
	date, p1name, p2name                             string
	p1needs, p1got, p2needs, p2got, p1skill, p2skill float64
	p1won, p2won                                     bool
}

func main() {
	// open, read, and parse the data into a chart structure
	f, err := os.Open("data/spring2018")
	if err != nil {
		fmt.Printf("shit: %v", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var matches []Match
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

	// Hang on to recorded dates of any matches to build out the graph data.
	uniqDates := make(map[string]bool)

	type player struct {
		// date to skill mapping
		skills map[string]int
	}

	players := make(map[string]player)
	for _, match := range matches {
		// player1
		_, ok := players[match.p1name]
		// player1 is new
		if !ok {
			players[match.p1name] = player{make(map[string]int)}
		}
		// player2
		_, ok = players[match.p2name]
		// player1 is new
		if !ok {
			players[match.p2name] = player{make(map[string]int)}
		}

		players[match.p1name].skills[match.date] = int(match.p1skill)
		players[match.p2name].skills[match.date] = int(match.p2skill)
		uniqDates[match.date] = true
	}

	// Sort the dates, because we used unordered map keys to uniq
	var dates []string
	for d := range uniqDates {
		dates = append(dates, d)
	}
	sort.Strings(dates)

	/*
	     datasets: [{
	         data: [86,114,106,106,107,111,133,221,783,2478],
	         label: "Africa",
	         borderColor: "#3e95cd",
	         fill: false
	       }, {
	         data: [282,350,411,502,635,809,947,1402,3700,5267],
	         label: "Asia",
	         borderColor: "#8e5ea2",
	         fill: false
	       }, {
	         data: [168,170,178,190,203,276,408,547,675,734],
	         label: "Europe",
	         borderColor: "#3cba9f",
	         fill: false
	       }, {
	         data: [40,20,10,16,24,38,74,167,508,784],
	         label: "Latin America",
	         borderColor: "#e8c3b9",
	         fill: false
	       }, {
	         data: [6,3,2,2,7,26,82,172,312,433],
	         label: "North America",
	         borderColor: "#c45850",
	         fill: false
	       }
	     ]
	   },
	*/

	// Print labels
	label := strings.Join(dates, ",")
	fmt.Printf("labels: [%s],\n", label)

	for name, p := range players {
		data := "\tdata: ["
		for _, d := range dates {
			s, ok := p.skills[d]
			// No match this week, assuming the previous week.
			if !ok {
				data += ("null, ")
				continue
			}
			data += fmt.Sprintf("%d, ", s)
		}
		fmt.Printf(data)
		fmt.Printf("]\n\tlabel: %q,\n", name)
		fmt.Printf("\tborderColor: #3e95cd,\n\tfill: false\n}, {\n")
	}

}
