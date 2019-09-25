package main

import (
	"fmt"
	c "github.com/sourcequench/league/common"
	p "github.com/sourcequench/league/parser"
	"log"
)

func main() {
	matches, err := p.Parse("data/ken-normalized")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	badGot := c.GotMistake(matches)
	for _, match := range badGot {
		fmt.Printf("%v\n", match)
	}
}
