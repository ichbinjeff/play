package game

import "math/rand"

type Card struct {
	Number int      `json:"number"`
	Label  string   `json:"label"`
	Front  bool     `json:"front"`
	Type   CardType `json:"type"`
}

type Cards struct {
	CardsToServe []Card
	CardsToTake  []Card
}

type CardType int

const (
	Club    CardType = 0
	Spade   CardType = 1
	Heart   CardType = 2
	Diamond CardType = 3
)

var cards Cards

func (cards *Cards) StartGame() {
	cardsToServe := []Card{}
	for i := 1; i <= 13; i++ {
		club := Card{Number: i, Front: false, Type: Club}
		spade := Card{Number: i, Type: Spade}
		heart := Card{Number: i, Type: Heart}
		diamond := Card{Number: i, Type: Diamond}
		cardsToServe = append(cardsToServe, []Card{club, spade, heart, diamond}...)
	}
	cards.CardsToServe = cardsToServe
	cards.CardsToTake = []Card{}
}

// Serve serves 1 card to 1 player
func (cards *Cards) Serve() Card {
	index := rand.Intn(len(cards.CardsToServe))
	card := cards.CardsToServe[index]
	cards.CardsToServe = append(cards.CardsToServe[:index], cards.CardsToServe[index+1:]...)
	cards.CardsToTake = append(cards.CardsToTake, card)
	return card
}

func (cards *Cards) ClearTake() {
	cards.CardsToTake = []Card{}
}

func Compare(card1 Card, card2 Card) int {
	return card1.Number - card2.Number
}
