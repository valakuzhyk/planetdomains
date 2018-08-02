package game

import (
	log "github.com/sirupsen/logrus"
	"github.com/valakuzhyk/planetdomains/card"
)

// TODO: If we want to handle the ability to choose what abilities activate when,
// we will need some kind of registration system for the abilities (I think)

func (p1 *person) PlayHand(p2 *person) {
	log.Infof("******Printing Turn for %s********", p1.Name)

	// Discard cards from your hand if needed
	p1.DiscardCards(p1.ToDiscard)

	// Apply effects from bases
	for _, b := range p1.Bases {
		b.PlayEffect(p1, p2)
	}

	// Apply all effects from cards
	for len(p1.Hand) != 0 {
		// Put the cards that are in the resolved pile back in the hand
		// to see if any synergies come about
		resolveCards(p1, p2)
		// Loop on buying cards and applying effects if those cards come into play
		purchaseCards(p1)
	}

	// Commit the effects after
	commitTurnState(p1, p2)

	p1.discardHandAndResolved()
	p1.turnState = turnState{}
	p1.drawToHand(5)
}

func resolveCards(p1, p2 *person) {
	for len(p1.Hand) != 0 {
		c := p1.Hand[0]
		log.Debugf("Playing %s", c.GetName())
		p1.Hand = p1.Hand[1:]

		c.PlayEffect(p1, p2)
		if base, ok := c.(card.Base); ok {
			p1.Bases = append(p1.Bases, base)
		} else {
			p1.ResolvedCards = append(p1.ResolvedCards, c)
		}
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
	// TODO
}
