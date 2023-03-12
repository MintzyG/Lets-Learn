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
	Bet      int
}

func (p Player) String() string {
	var PlayerText = ":\n"
	for _, v := range p.Hand.Cards {
		PlayerText += v.String() + " "
	}

	return fmt.Sprintf("PlayerID: %v\nMoney: %v\nHand%v", p.PlayerID, p.Money, PlayerText)
}

type Game struct {
	Players         []Player
	Dealer          Player
	Deck            Container
	NumberOfPlayers int
	CurrentPlayer   int
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

type GameLoop interface {
	FirstDeal()
	DealCard(p *Player)
	TakeBets()
	CheckBj() bool
}

func PlayRound(gl GameLoop) bool {
	gl.TakeBets()
	gl.FirstDeal()
	if gl.CheckBj() {
		return true
	}

	return false
}

func main() {

	Blackjack := Game{NumberOfPlayers: 1}
	SetupGame(&Blackjack)

	// Game Loop
	for {
		if PlayRound(&Blackjack) {
			continue
		}
	}

}

func (g *Game) SetupDeck() {

	deck := Container{Name: "Deck"}
	g.Deck = deck
	cardSetup := Card{}
	for i := 0; i < 2; i++ {
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

	g.Dealer.PlayerID = 0

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

// DealCard Retrieves last Card from Deck, appends it to the given Player, then erases said Card from Deck
func (g *Game) DealCard(p *Player) {
	lastCard := g.Deck.Cards[len(g.Deck.Cards)-1]
	p.Hand.Cards = append(p.Hand.Cards, lastCard)
	g.Deck.Cards = g.Deck.Cards[:len(g.Deck.Cards)-1]
}

func (g *Game) FirstDeal() {

	if len(g.Deck.Cards) < len(g.Players)*2+2 {
		g.CleanDeck()
		g.SetupDeck()
		g.ShuffleDeck()
	}

	for i := 0; i < 2; i++ {
		for i := 0; i < len(g.Players); i++ {
			g.DealCard(&g.Players[i])
		}
		g.DealCard(&g.Dealer)
	}
}

func (g *Game) CleanDeck() {
	g.Deck.Cards = g.Deck.Cards[:0]
}

func (g *Game) TakeBets() {

	var betAmount int

	for i, v := range g.Players {
		fmt.Printf("Player %v: How much would you like to bet, you have %v\n", v.PlayerID, v.Money)
		for {
			_, err := fmt.Scanln(&betAmount)
			if err != nil {
				fmt.Println("ERROR: You made a invalid bet, please bet again")
				continue
			}

			if betAmount > v.Money || betAmount <= 0 {
				fmt.Println("You made a invalid bet, please bet again")
				continue
			}

			g.Players[i].Bet = betAmount
			g.Players[i].Money -= betAmount

			fmt.Println("You've bet", g.Players[i].Bet, "now you have", v.Money)

			break

		}
	}
}

//TODO: code the ReturnBets(), TableWins() and PayoutWinners() functions

func (g *Game) CheckBj() bool {

	var TotalValue int
	var Blackjacks []Player
	var Lost []Player

	for _, v := range g.Dealer.Hand.Cards {
		TotalValue += v.Value
	}

	if TotalValue == 11 {
		Blackjacks = append(Blackjacks, g.Dealer)
	}

	for i := range g.Players {
		TotalValue = 0
		for _, w := range g.Players[i].Hand.Cards {
			TotalValue += w.Value
		}
		if TotalValue == 11 {
			Blackjacks = append(Blackjacks, g.Players[i])
		} else {
			Lost = append(Lost, g.Players[i])
		}
	}

	if len(Blackjacks) > 1 && Blackjacks[0].PlayerID == 0 {
		ReturnBets(g, Blackjacks...)
		// LostBet(g, Lost...)
		// return true
	} else if len(Blackjacks) == 1 && Blackjacks[0].PlayerID == 0 {
		// TableWins(g)
		// LostBet(g, Lost...)
		// return true
	} else {
		// PayoutWinners(g, Blackjacks...)
		// LostBet(g, Lost...)
		// return true
	}

	return false
}

func ReturnBets(g *Game, p ...Player) {

	for i := range p {
		if g.Players[i].PlayerID != 0 {
			g.Players[i].Money += g.Players[i].Bet
			g.Players[i].Bet = 0
		}
	}

}
