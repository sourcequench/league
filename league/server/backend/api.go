// Package main implements a simple gRPC server
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Some structures
type Season struct {
	startDate string  `json:"startDate"`
	endDate   string  `json:"endDate"`
	timeSpec  string  `json:"timeSpec"`
	matches   []Match `json:"matches"`
}

type Match struct {
	date    string `json:"date"`
	p1Name  string `json:"p1Name"`
	p1Needs string `json:"p1Needs"`
	p1Got   string `json:"p1Got"`
	p1Skill int    `json:"p1Skill"`
	p2Name  string `json:"p2Name"`
	p2Needs string `json:"p2Needs"`
	p2Got   string `json:"p2Got"`
	p2Skill int    `json:"p2Skill"`
}

type Player struct {
	firstName string  `json:"firstName"`
	lastName  string  `json:"lastName"`
	nickName  string  `json:"nickName"`
	email     string  `json:"email"`
	phone     string  `json:"phone"`
	mu        float64 `json:"mu"`
	sigma     float64 `json:"sigma"`
	waitList  bool    `json:"waitList"`
	active    bool    `json:"active"`
}

type Roster []Player

func GetRoster(w http.ResponseWriter, r *http.Request) {
	roster := Roster{
		Player{firstName: "Ryan", lastName: "Shea"},
		Player{firstName: "Brian", lastName: "Begnoche"},
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	j, err := json.Marshal(roster)
	log.Printf("encoding roster: %v %v", j, err)
	if err := json.NewEncoder(w).Encode(roster); err != nil {
		panic(err)
	}
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", GetRoster)
	http.HandleFunc("/api", GetRoster)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
