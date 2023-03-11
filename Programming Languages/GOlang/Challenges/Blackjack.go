package main

import (
	"fmt"
	"math/rand"
	"time"
)

var suits = []string{"Hearts", "Clubs", "Diamonds", "Spades"}
var values = []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

type Card struct {
	Suit  string
	Face  string
	Value int
}

func (c Card) String() string {
	return fmt.Sprintf("%v of %v value: %v", c.Face, c.Suit, c.Value)
}

type Container struct {
	Cards []Card
	Name  string
}

type Player struct {
	Hand     Container
	Money    int
	PlayerID int
}

func (p Player) String() string {
	var PlayerText = "Hand:\n"
	for _, v := range p.Hand.Cards {
		PlayerText += v.String()
	}

	return fmt.Sprintf("PlayerID: %v\nMoney: %v\nHand%v", p.PlayerID, p.Money, PlayerText)
}

type Game struct {
	Players         []Player
	Dealer          Player
	Deck            Container
	NumberOfPlayers int
}

type Setup interface {
	SetupDeck()
	SetupPlayers()
	ShuffleDeck()
}

func SetupGame(s Setup) {
	s.SetupDeck()
	s.SetupPlayers()
	s.ShuffleDeck()
}

func main() {

	Blackjack := Game{NumberOfPlayers: 1}
	SetupGame(&Blackjack)
	Blackjack.Deck.PrintContainer()
	Blackjack.PrintPlayers()
}

func (g *Game) SetupDeck() {

	deck := Container{Name: "Deck"}
	g.Deck = deck
	cardSetup := Card{}

	for _, v := range suits {
		for _, w := range values {
			cardSetup.Suit = v
			cardSetup.Face = w
			switch cardSetup.Face {
			case "Jack", "Queen", "King":
				cardSetup.Value = 10
			case "Ace":
				cardSetup.Value = 1
			case "Two":
				cardSetup.Value = 2
			case "Three":
				cardSetup.Value = 3
			case "Four":
				cardSetup.Value = 4
			case "Five":
				cardSetup.Value = 5
			case "Six":
				cardSetup.Value = 6
			case "Seven":
				cardSetup.Value = 7
			case "Eight":
				cardSetup.Value = 8
			case "Nine":
				cardSetup.Value = 9
			case "Ten":
				cardSetup.Value = 10
			default:
				cardSetup.Value = 0
			}
			g.Deck.Cards = append(g.Deck.Cards, cardSetup)
		}
	}
}

func (g *Game) ShuffleDeck() {

	rand.Seed(time.Now().Unix())

	for i := 0; i < 3; i++ {
		for i := 0; i < len(g.Deck.Cards); i++ {
			r := rand.Intn(len(g.Deck.Cards))

			g.Deck.Cards[i], g.Deck.Cards[r] = g.Deck.Cards[r], g.Deck.Cards[i]

		}
	}
}

func (g *Game) SetupPlayers() {

	p := Player{Money: 1000}
	p.Hand.Name = "Player Hand"

	for i := 0; i < g.NumberOfPlayers; i++ {
		p.PlayerID = i + 1
		g.Players = append(g.Players, p)
	}
}

func (c Container) PrintContainer() {
	fmt.Println("Container:", c.Name, "Contents:")
	for _, v := range c.Cards {
		fmt.Println(v)
	}
}

func (g *Game) PrintPlayers() {
	for _, v := range g.Players {
		fmt.Println(v)
	}
}
