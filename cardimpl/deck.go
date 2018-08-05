package cardimpl

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/valakuzhyk/planetdomains/card"
)

// NewDeck returns a new deck
func NewDeck() card.Deck {
	return &deckImpl{}
}

type deckImpl struct {
	cards []card.Card
}

func (d *deckImpl) Len() int {
	return len(d.cards)
}

func (d *deckImpl) IsEmpty() bool {
	return d.Len() == 0
}

func (d *deckImpl) Peek() card.Card {
	if d.IsEmpty() {
		panic("Can't peek with no cards left!")
	}
	return d.cards[0]
}

func (d *deckImpl) Draw() card.Card {
	if d.IsEmpty() {
		panic("Don't draw when there aren't any cards left!")
	}
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card
}

func (d *deckImpl) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := len(d.cards) - 1; i > 0; i-- { // Fisher–Yates shuffle
		j := rand.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

func (d *deckImpl) DrawAll() []card.Card {
	allCards := d.cards
	d.cards = []card.Card{}
	return allCards
}

func (d *deckImpl) PlaceOnBottom(c ...card.Card) {
	d.cards = append(d.cards, c...)
}

func (d *deckImpl) PlaceOnTop(c ...card.Card) {
	d.cards = append(c, d.cards...)
}

func (d *deckImpl) String() string {
	s := []string{"Cards in deck are: "}
	for _, c := range d.cards {
		s = append(s, fmt.Sprintf("    %s", c.GetName()))
	}
	return strings.Join(s, "\n")
}
