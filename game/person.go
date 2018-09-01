package game

import (
	"fmt"
	"log"
	"strings"

	"github.com/valakuzhyk/planetdomains/internal"
	"github.com/valakuzhyk/planetdomains/utils"

	"github.com/valakuzhyk/planetdomains/card"
)

// person is a person playing the game
type person struct {
	Name string

	field     *Field
	Authority int
	Deck      card.Deck
	Discard   card.Group
	Bases     card.Group

	Hand card.Group
	turnState
}

type turnState struct {
	// Present at the beginning of the turn, must be resolved before doing anything else.
	ToDiscard int

	PlayedCards         card.Group
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
		Discard:   card.Group{},
		field:     field,
	}, nil
}

func (p *person) String() string {
	s := []string{
		fmt.Sprintf("*********PRINTING PLAYER STATE ***************"),
		fmt.Sprintf("%s, Authority: %d", p.Name, p.Authority)}
	s = append(s, fmt.Sprintf("Deck len: %d, Discard len: %d", p.Deck.Len(), p.Discard.Len()))
	s = append(s, "")

	if p.Bases.Len() > 0 {
		s = append(s, fmt.Sprintf("Bases: %v", p.Bases))
	}

	s = append(s, fmt.Sprintf("Hand: %v", p.Hand))

	return strings.Join(s, "\n")
}

// places the hand in the discard pile
func (p *person) discardHandAndResolved() {
	for p.Hand.Len() != 0 {
		p.Discard.Add(p.Hand.Take(0))
	}
	for p.PlayedCards.Len() != 0 {
		p.Discard.Add(p.PlayedCards.Take(0))
	}
}

func (p *person) DrawCards(n uint) {
	p.drawToHand(n)
}

// draws cards from the deck to the hand
func (p *person) drawToHand(n uint) {
	p.Hand.Add(p.draw(n)...)
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
		c := p.Deck.Draw()
		c.Reset()
		drawnCards = append(drawnCards, c)
	}

	if p.Discard.Len() == 0 {
		return drawnCards
	}
	p.Deck.PlaceOnBottom(p.Discard.TakeAll()...)
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
		fmt.Printf("%s has died\n", p.Name)
		log.Fatal()
	}
	p.Authority += authority
}

func (p *person) MustDiscard() {
	p.ToDiscard++
}

func (p *person) DiscardCards(numToDiscard int) {
	for numToDiscard > 0 && p.Hand.Len() > 0 {
		prompt := fmt.Sprintf("You must discard %d card(s), what would you like to discard next?", numToDiscard)
		i := utils.PickCard(prompt, p.Hand.Cards, true /* required */)
		if i < 0 {
			continue
		}
		p.Discard.Add(p.Hand.Take(i))
		numToDiscard--
	}
}

func (p *person) DestroyBase(opponent internal.Player) {

}

// ScrapFromHand allows you to scrap any other card from your hand.
func (p *person) ScrapFromHand() {
	cardIdx := utils.PickCard("What card would you like to scrap from your hand?", p.Hand.Cards, true)
	_ = p.Hand.Take(cardIdx)
}

func (p *person) ScrapFromDiscard() {
	cardIdx := utils.PickCard("What card would you like to scrap from your discard pile?", p.Discard.Cards, true)
	_ = p.Discard.Take(cardIdx)
}

func (p *person) ScrapFromTradeRow(n uint) {

}

func (p *person) ScrapFromHandOrDiscard(n uint) {

}
