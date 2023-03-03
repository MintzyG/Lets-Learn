package main

import (
	"fmt"
	"math/rand"
	"time"
)

var names []string
var suits []string

type Icards interface {
	SetupValues()
}

type Ideck interface {
	SetupDeck()
	PrintDeck()
	ShuffleDeck()
}

type Iplayer interface {
}

type Igame interface {
	SetupGame()
}

type Deck struct {
	cards []Cards
}

type Cards struct {
	values int
	suits  string
	names  string
}

type Player struct {
	hand  []Cards
	money int
}

type Game struct {
	players []Player
	dealer  Player
	deck    Deck
}

func main() {

	names = []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	suits = []string{"Clubs", "Spades", "Hearts", "Diamonds"}

	var game Game
	game.SetupGame()

}

func (d *Deck) SetupDeck() {
	var card Cards
	for _, suit := range suits {
		for _, name := range names {
			card = Cards{suits: suit, names: name}
			card.values = card.SetupValues()
			d.cards = append(d.cards, card)
		}
	}
}

func (d *Deck) PrintDeck() {
	for _, card := range d.cards {
		fmt.Println(card)
	}
}

func (d *Deck) ShuffleDeck() {
	var shuffleDeck []Cards

	rand.Seed(time.Now().Unix())

}

func (c *Cards) SetupValues() int {
	switch c.names {
	case "Ace":
		return 0
	case "Two":
		return 2
	case "Three":
		return 3
	case "Four":
		return 4
	case "Five":
		return 5
	case "Six":
		return 6
	case "Seven":
		return 7
	case "Eight":
		return 8
	case "Nine":
		return 9
	case "Ten", "Jack", "Queen", "King":
		return 10
	}
	panic("SHIT")
}

func (g *Game) SetupGame() {
	g.deck.SetupDeck()
	g.deck.ShuffleDeck()
}
