package game

import (
	"github.com/valakuzhyk/planetdomains/card"
)

// Field contains cards that are not affiliated with a player yet
type Field struct {
	TradeRow  []card.Card
	Explorers card.Deck
	ScrapHeap []card.Card
	TradeDeck card.Deck
}

// CreateFor2 returns a new game for 2.
func CreateFor2() (*person, *person, error) {
	f := &Field{}
	p1, err := newPerson("Player 1", 50, card.DefaultStarterDeck(), f)
	if err != nil {
		return nil, nil, err
	}
	p1.drawToHand(3)

	p2, err := newPerson("Player 2", 50, card.DefaultStarterDeck(), f)
	if err != nil {
		return nil, nil, err
	}
	p2.drawToHand(5)

	return &p1, &p2, nil
}
