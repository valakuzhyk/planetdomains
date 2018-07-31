package card

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

// Deck allows for interacting with a deck of cards
type Deck interface {
	Shuffle()
	Len() int
	IsEmpty() bool
	Draw() Card
	DrawAll() []Card
	PlaceOnTop(...Card)
	PlaceOnBottom(...Card)
	fmt.Stringer
}

// NewDeck returns a new deck
func NewDeck() Deck {
	return &deckImpl{}
}

type deckImpl struct {
	cards []Card
}

func (d *deckImpl) Len() int {
	return len(d.cards)
}

func (d *deckImpl) IsEmpty() bool {
	return d.Len() == 0
}

func (d *deckImpl) Draw() Card {
	if d.IsEmpty() {
		log.Fatal("Don't draw when there aren't any cards left!")
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

func (d *deckImpl) DrawAll() []Card {
	allCards := d.cards
	d.cards = []Card{}
	return allCards
}

func (d *deckImpl) PlaceOnBottom(c ...Card) {
	d.cards = append(d.cards, c...)
}

func (d *deckImpl) PlaceOnTop(c ...Card) {
	d.cards = append(c, d.cards...)
}

func (d *deckImpl) String() string {
	s := []string{"Cards in deck are: "}
	for _, c := range d.cards {
		s = append(s, fmt.Sprintf("    %s", c.GetName()))
	}
	return strings.Join(s, "\n")
}
