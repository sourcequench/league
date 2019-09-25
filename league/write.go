package main

import (
	p "github.com/sourcequench/league/parser"
	"log"
)

func main() {
	matches, err := p.Parse("data/ken-normalized")
	if err != nil {
		log.Fatalf("could not open file: %v", err)
	}
	p.ProtoOut(matches)
}
