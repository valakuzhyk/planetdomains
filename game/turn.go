package game

import (
	log "github.com/sirupsen/logrus"

	"github.com/valakuzhyk/planetdomains/card"
)

/* External Person interface */
func (p1 *person) PlayHand(p2 *person) {
	log.Infof("******Printing Turn for %s********", p1.Name)

	// Discard cards from your hand if needed
	p1.DiscardCards(p1.NumToDiscard)

	// Apply all effects from cards
	resolveCards(p1, p2, p1.Hand...)
	// Loop on buying cards and applying effects if those cards come into play
	purchaseCards(p1)

	// Commit the effects after
	commitTurnState(p1, p2)

	p1.turnState = turnState{}
	p1.discardHand()
	p1.drawToHand(5)
}

func purchaseCards(p1 *person) {
	// If possible try and purchase some cards
}

func resolveCards(p1, p2 *person, cards ...card.Card) {
	for _, c := range cards {
		c.PlayEffect(p1)
	}
}

func commitTurnState(p1, p2 *person) {
	// Waste whatever money wasn't spent
	log.Printf("Wasted %d trade", p1.turnState.Trade)
	log.Printf("Applied %d damage to opponent", p1.turnState.Combat)
	p2.AddAuthority(-p1.turnState.Combat)
	log.Printf("Recovered %d authority", p1.turnState.AdditionalAuthority)
	p1.AddAuthority(p1.turnState.AdditionalAuthority)
	return
}

func (p *person) DiscardCards(numToDiscard int) {
	if numToDiscard > 0 {
		log.Println("Discarding at the beginning of your turn is not implemented yet")
	}
}
