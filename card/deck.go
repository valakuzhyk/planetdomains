package card

import "fmt"

// Deck allows for interacting with a deck of cards
type Deck interface {
	Shuffle()
	Len() int
	IsEmpty() bool
	Peek() Card
	Draw() Card
	DrawAll() []Card
	PlaceOnTop(...Card)
	PlaceOnBottom(...Card)
	fmt.Stringer
}
