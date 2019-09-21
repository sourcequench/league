package parser

import (
	"bufio"
	"fmt"
	s "github.com/sourcequench/league/common"
	"os"
	"strconv"
	"strings"
)

func Parse(f string) ([]s.Match, error) {
	var matches []s.Match
	fh, err := os.Open(f)
	if err != nil {
		return nil, fmt.Errorf("Parse: file open error %v", err)
	}
	defer fh.Close()

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		p1skill, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			fmt.Printf("error getting p1skill: %v\n", fields)
		}
		p1needs, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			fmt.Printf("p1needs: %v\n", err)
		}
		p1got, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			fmt.Printf("error getting p1got: %v\n", fields)
		}
		p2skill, err := strconv.ParseFloat(fields[6], 64)
		if err != nil {
			fmt.Printf("error getting p2skill: %v\n", fields)
		}
		p2needs, err := strconv.ParseFloat(fields[7], 64)
		if err != nil {
			fmt.Printf("error getting p2needs: %v\n", err)
		}
		p2got, err := strconv.ParseFloat(fields[8], 64)
		if err != nil {
			fmt.Printf("error getting p2got: %v\n", fields)
		}

		p1name := fields[0]
		p2name := fields[5]

		m := s.Match{
			P1name:  p1name,
			P2name:  p2name,
			Date:    fields[10],
			P1needs: p1needs,
			P1got:   p1got,
			P2needs: p2needs,
			P2got:   p2got,
			P1skill: p1skill,
			P2skill: p2skill,
		}
		matches = append(matches, m)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner error: %v", err)
	}

	return matches, nil
}
