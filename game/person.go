package game

import (
	"fmt"
	"log"
	"strings"

	"github.com/valakuzhyk/planetdomains/card"
)

// person is a person playing the game
type person struct {
	Name string

	field     *Field
	Authority int
	Deck      card.Deck
	Discard   card.Deck
	Bases     []card.Card

	Hand []card.Card
	turnState
}

type turnState struct {
	// Present at the beginning of the turn, must be resolved before doing anything else.
	ToDiscard int

	ResolvedCards       []card.Card
	Trade               int
	Combat              int
	AdditionalAuthority int
}

// newPerson returns a new player for the game
func newPerson(name string, startingAuthority int, startingDeck card.Deck, field *Field) (person, error) {
	if startingAuthority <= 0 {
		return person{}, fmt.Errorf("Unable to create player with < 0 authority")
	}
	if startingDeck.IsEmpty() {
		return person{}, fmt.Errorf("Player needs to have cards for a deck")
	}

	return person{
		Name:      name,
		Authority: startingAuthority,
		Deck:      startingDeck,
		Discard:   card.NewDeck(),
		field:     field,
	}, nil
}

func (p *person) String() string {
	s := []string{
		fmt.Sprintf("*********PRINTING PLAYER STATE ***************"),
		fmt.Sprintf("%s, Authority: %d", p.Name, p.Authority)}
	s = append(s, fmt.Sprintf("Deck len: %d, Discard len: %d", p.Deck.Len(), p.Discard.Len()))
	s = append(s, "")
	cards := []string{}
	for _, c := range p.Hand {
		cards = append(cards, card.String(c))
	}

	s = append(s, "Hand: "+strings.Join(cards, ", "))

	return strings.Join(s, "\n")
}

// places the hand in the discard pile
func (p *person) discardHandAndResolved() {
	for _, card := range p.Hand {
		p.Discard.PlaceOnTop(card)
	}
	for _, card := range p.ResolvedCards {
		p.Discard.PlaceOnTop(card)
	}
	p.Hand = []card.Card{}
	p.ResolvedCards = []card.Card{}
}

// draws cards from the deck to the hand
func (p *person) drawToHand(n uint) {
	p.Hand = append(p.Hand, p.draw(n)...)
}

func (p *person) draw(n uint) []card.Card {
	// Draw until either
	//  * you have reached n cards
	//  * you have drawn your deck, at which point draw as many cards as you
	//    need out of discard
	// You may not draw enough cards, at which point further calls to draw
	// return an empty array
	drawnCards := []card.Card{}
	for !p.Deck.IsEmpty() {
		if uint(len(drawnCards)) == n {
			return drawnCards
		}
		drawnCards = append(drawnCards, p.Deck.Draw())
	}

	if p.Discard.IsEmpty() {
		return drawnCards
	}
	p.Deck.PlaceOnBottom(p.Discard.DrawAll()...)
	p.Deck.Shuffle()

	if n < uint(len(drawnCards)) {
		log.Fatal("Somehow, we have gotten more cards than we should. Exiting before ridiculousness")
	}
	cardsLeftToDraw := n - uint(len(drawnCards))

	return append(drawnCards, p.draw(cardsLeftToDraw)...)
}

func (p *person) AddTrade(trade int) error {
	if p.Trade+trade < 0 {
		return fmt.Errorf("Existing trade: %d, adding %d results in trade below 0", p.Trade, trade)
	}
	p.Trade += trade
	return nil
}

func (p *person) GetTrade() int {
	return p.Trade
}

func (p *person) AddCombat(combat int) error {
	if p.Combat+combat < 0 {
		return fmt.Errorf("Existing Combat: %d, adding %d results in Combat below 0", p.Combat, combat)
	}
	p.Combat += combat
	return nil
}

func (p *person) AddAuthority(authority int) {
	if p.Authority+authority < 0 {
		log.Fatal("You have died.")
	}
	p.Authority += authority
}

func (p *person) DrawCards(n uint) {
	p.drawToHand(n)
}

func (p *person) DiscardCard() {
	p.ToDiscard++
}

func (p *person) DestroyBase() {

}

func (p *person) ScrapFromHand() {

}

func (p *person) ScrapFromDiscard() {

}

func (p *person) ScrapFromTradeRow(n uint) {

}

func (p *person) ScrapFromHandOrDiscard(n uint) {

}

func (p *person) AcquireCardLessThan(maxCost uint) {

}
