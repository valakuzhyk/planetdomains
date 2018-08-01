package game

import (
	"fmt"
	"strings"

	"github.com/valakuzhyk/planetdomains/card"
)

// Field contains cards that are not affiliated with a player yet
type Field struct {
	TradeRow  []card.Card
	Explorers card.Deck
	ScrapHeap []card.Card
	TradeDeck card.Deck
}

func (f Field) String() string {
	s := []string{"******* PRINTING FIELD *******"}
	s = append(s, fmt.Sprintf("Explorers: %d, Deck: %d, Scrap Heap: %d", f.Explorers.Len(), f.TradeDeck.Len(), len(f.ScrapHeap)))
	// Print cards on trade row
	// Print length of scrap heap
	// Print length of deck
	return strings.Join(s, "\n")
}

// CreateFor2 returns a new game for 2.
func CreateFor2() (*person, *person, error) {
	tradeDeck := card.StandardDeck()
	tradeRow := []card.Card{}
	for i := 0; i < 5; i++ {
		c := tradeDeck.Draw()
		tradeRow = append(tradeRow, c)
	}

	f := &Field{
		TradeDeck: tradeDeck,
		TradeRow:  tradeRow,
		Explorers: card.NewExplorerDeck(),
		ScrapHeap: []card.Card{},
	}

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
