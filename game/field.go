package game

import (
	"fmt"
	"strings"

	"github.com/ryanuber/columnize"

	"github.com/valakuzhyk/planetdomains/card"
	"github.com/valakuzhyk/planetdomains/decks/colonywars"
	"github.com/valakuzhyk/planetdomains/decks/standard"
)

// Field contains cards that are not affiliated with a player yet
type Field struct {
	activePlayer int
	players      []*person
	TradeRow     []card.Card
	Explorers    card.Deck
	ScrapHeap    []card.Card
	TradeDeck    card.Deck
}

func (f Field) GetActivePerson() *person {
	return f.players[f.activePlayer]
}

func (f Field) String() string {
	p1 := f.players[0]
	p2 := f.players[1]
	s := []string{}
	table := []string{
		fmt.Sprintf("----%s---- | ----Field---- | ----%s----", p1.Name, p2.Name),
		fmt.Sprintf("Authority: %d | Explorers: %d | Authority: %d", p1.Authority, f.Explorers.Len(), p2.Authority),
		fmt.Sprintf("Combat:    %d | Deck:      %d | Combat:    %d", p1.Combat, f.TradeDeck.Len(), p2.Combat),
		fmt.Sprintf("Trade:     %d | Scrap:     %d | Trade:     %d", p1.Trade, len(f.ScrapHeap), p2.Trade),
		fmt.Sprintf("Deck:      %d |               | Deck:      %d", p1.Deck.Len(), p2.Deck.Len()),
		fmt.Sprintf("Hand:      %d |               | Hand:      %d", len(p1.Hand), len(p2.Hand)),
		fmt.Sprintf("Discard:   %d |               | Discard:   %d", p1.Discard.Len(), p2.Discard.Len()),
	}
	result := columnize.SimpleFormat(table)
	s = append(s, result)
	s = append(s, fmt.Sprintf("Trade Row: %s", card.List(f.TradeRow...)))
	s = append(s, "")

	activePlayer := f.players[f.activePlayer]
	s = append(s, fmt.Sprintf("%s Turn", activePlayer.Name))
	if len(activePlayer.Bases) != 0 {
		s = append(s, fmt.Sprintf("Bases: %v", card.List(activePlayer.Bases...)))
	}
	if len(activePlayer.PlayedCards) != 0 {
		s = append(s, fmt.Sprintf("Played: %v", card.List(activePlayer.PlayedCards...)))
	}
	s = append(s, fmt.Sprintf("Hand: %v", card.List(activePlayer.Hand...)))

	s = append(s, "")
	// Print cards on trade row
	// Print length of scrap heap
	// Print length of deck
	return strings.Join(s, "\n")
}

// CreateFor2 returns a new game for 2.
func CreateFor2() (*Field, error) {
	tradeDeck := colonywars.Deck()
	tradeRow := []card.Card{}
	for i := 0; i < 5; i++ {
		c := tradeDeck.Draw()
		tradeRow = append(tradeRow, c)
	}

	f := &Field{
		TradeDeck: tradeDeck,
		TradeRow:  tradeRow,
		Explorers: standard.NewExplorerDeck(),
		ScrapHeap: []card.Card{},
	}

	p1, err := newPerson("Player 1", 50, standard.StarterDeck(), f)
	if err != nil {
		return nil, err
	}
	p1.drawToHand(3)

	p2, err := newPerson("Player 2", 50, standard.StarterDeck(), f)
	if err != nil {
		return nil, err
	}
	p2.drawToHand(5)
	f.players = []*person{&p1, &p2}
	f.activePlayer = 0

	return f, nil
}
