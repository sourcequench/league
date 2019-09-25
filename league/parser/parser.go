package parser

import (
	"bufio"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	c "github.com/sourcequench/league/common"
	pb "github.com/sourcequench/league/proto"
	"log"
	"os"
	"strconv"
	"strings"
)

func Parse(f string) ([]c.Match, error) {
	var matches []c.Match
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

		m := c.Match{
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

func ProtoOut(matches []c.Match) {
	season := &pb.Season{}
	for _, match := range matches {
		m := pb.Match{
			P1Name:  match.P1name,
			P2Name:  match.P2name,
			P1Needs: match.P1needs,
			P2Needs: match.P2needs,
			P1Got:   match.P1got,
			P2Got:   match.P2got,
			P1Skill: match.P1skill,
			P2Skill: match.P2skill,
			Date:    match.Date,
		}
		season.Matches = append(season.Matches, &m)
	}

	marshaler := &jsonpb.Marshaler{Indent: "  ", EmitDefaults: false, EnumsAsInts: false, OrigName: true}

	f, err := os.Create("matches.json")
	if err != nil {
		log.Fatalln("Failed to create file:", err)
	}
	w := bufio.NewWriter(f)

	err = marshaler.Marshal(w, season)
	if err != nil {
		log.Fatalln("Failed to write season:", err)
	}
}
