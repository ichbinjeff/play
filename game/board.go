package game

import (
	"fmt"
	"sort"
)

type Board struct {
	players []Player
	Cards   Cards
}

func (board *Board) Join(player Player) {
	board.players = append(board.players, player)
}

func (board *Board) Start() {
	board.Cards.StartGame()
}

func (board *Board) Play() {
	if len(board.Cards.CardsToServe) == 0 {
		fmt.Println("game over %s is the winer", board.GetWiner().Name)
	}
	if len(board.players) <= 1 {
		fmt.Println("require at least 2 players")
	}
	for _, p := range board.players {
		card := board.Cards.Serve()
		p.AtHand = card
	}
	sort.Sort(byCards(board.players))

}

func (board *Board) GetWiner() Player {

}

func sortPlayerByCardsNum(players []Player) []Player {
	return
}

type byCards []Player

func (s byCards) Len() int {
	return len(s)
}

func (s byCards) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byCards) Less(i, j int) bool {
	return s[i].AtHand.Number < s[j].AtHand.Number
}
