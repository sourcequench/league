package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	ts "github.com/mafredri/go-trueskill"
	"github.com/mafredri/go-trueskill/gaussian"
	"html/template"
	"io/ioutil"
	"log"
	"math"
	"net/http"

	pb "github.com/sourcequench/trueskill/proto"
)

var templates = template.Must(
	template.ParseFiles(
		"/home/ryan/go/src/github.com/sourcequench/trueskill/index.html",
	),
)

// Match type enum
type matchType int32

const (
	season matchType = iota
	playoffs
)

// Offset enum
type offset int32

const (
	makeup matchType = iota
	playahead
)

type Player struct {
	playerId  int32
	firstName string
	lastName  string
	nickName  string
	email     string
	phone     string
	mu        float32
	sigma     float32
	//	matches   []*Match
	waitList bool
	active   bool
}

func (p *Player) toProto() *pb.Player {
	player := &pb.Player{
		PlayerId:  p.playerId,
		FirstName: p.firstName,
		LastName:  p.lastName,
		NickName:  p.nickName,
		Email:     p.email,
		Phone:     p.phone,
		Mu:        p.mu,
		Sigma:     p.sigma,
		WaitList:  p.waitList,
		Active:    p.active,
	}
	return player
}

func (p *Player) save() error {
	roster, err := readRoster()
	if err != nil {
		fmt.Errorf("failed to read roster")
	}
	// TODO: check for player already existing
	roster.Players = append(roster.Players, p.toProto())
	out := proto.MarshalTextString(roster)
	log.Print("writing out roster file")
	if err := ioutil.WriteFile("roster.pb", []byte(out), 0644); err != nil {
		log.Fatalln("Failed to write player log:", err)
	}
	return nil
}

func readRoster() (*pb.Roster, error) {
	filename := "roster.pb"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	roster := &pb.Roster{}
	if err := proto.UnmarshalText(string(data), roster); err != nil {
		return nil, err
	}
	log.Print("read roster file successfully")
	return roster, nil
}

type Roster struct {
	players []Player
}

type Match struct {
	p1Id, p2Id      int32
	win             bool
	myGames         int32
	yourGgames      int32
	myGamesNeeded   int32
	yourGamesNeeded int32
	// what will this do?
	matchType matchType
	date      string
}

type Season struct {
	startDate string
	endDate   string
	timeSpec  string
	matches   []*Match
}

func race(maxGames int, p1, p2 ts.Player) (float64, float64) {
	// TODO: should we always take the floor of the lower player?
	// TODO: what is the logic behind different max games for different
	// races? poorer players finish games more slowly - is that actually true?
	p1games := math.Floor(pwin(p1, p2) * float64(maxGames))
	return p1games, float64(maxGames) - p1games
}

func pwin(p1, p2 ts.Player) float64 {
	// TODO: do a ts.New and set beta on that?
	beta := float64(60)
	deltaMu := p1.Mean() - p2.Mean()
	sumSigma := p1.Variance() + p2.Variance()
	rss := math.Sqrt(2*(beta*beta) + sumSigma)
	return gaussian.NormCdf(deltaMu / rss)
}

// Handlers
func indexHandler(w http.ResponseWriter, r *http.Request) {
	rpb, err := readRoster()
	if err != nil {
		log.Fatal("got error reading roster")
	}
	roster := Roster{}
	for _, p := range rpb.GetPlayers() {
		player := Player{
			playerId: p.GetPlayerId(),
		}
		roster.players = append(roster.players, player)
	}
	renderTemplate(w, "index", r)
}

func renderTemplate(w http.ResponseWriter, tmpl string, i interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	log.Fatal(http.ListenAndServe(":9999", nil))
}
