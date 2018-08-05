package game

import (
	"fmt"

	"github.com/manifoldco/promptui"

	"github.com/buger/goterm"
	log "github.com/sirupsen/logrus"
	"github.com/valakuzhyk/planetdomains/card"
	"github.com/valakuzhyk/planetdomains/utils"
)

// TODO: If we want to handle the ability to choose what abilities activate when,
// we will need some kind of registration system for the abilities (I think)

func (f *Field) PlayHand() {
	playCardsOption := "Play Cards"
	activateAbilitiesOption := "Activate Abilities"
	buyCardsOption := "Buy Cards"
	endTurnOption := "End Turn"

	p1 := f.players[f.activePlayer]
	p2 := f.players[(f.activePlayer+1)%len(f.players)]

	// Discard cards from your hand if needed
	p1.DiscardCards(p1.ToDiscard)
	for _, b := range p1.Bases.Cards {
		b.Reset()
	}

	// Apply all effects from cards
	for true {
		goterm.Clear()
		goterm.Flush()
		fmt.Println(p1.field)

		options := []string{endTurnOption, activateAbilitiesOption}
		if p1.Hand.Len() != 0 {
			options = append(options, playCardsOption)
		}
		if p1.Trade != 0 {
			options = append(options, buyCardsOption)
		}

		prompt := promptui.Select{
			Label: "Which would you like to do?:",
			Items: options,
			Size:  5,
		}
		_, chosenOption, err := prompt.Run()

		if err != nil {
			log.Fatal("Error choosing option: ", err)
		}
		if chosenOption == endTurnOption {
			break
		}

		switch chosenOption {
		case buyCardsOption:
			purchaseCards(p1)
		case activateAbilitiesOption:
			resolveCards(p1, p2)
		case playCardsOption:
			resolveCards(p1, p2)
		}
	}

	// Commit the effects after
	commitTurnState(p1, p2)

	p1.discardHandAndResolved()
	p1.turnState = turnState{}
	p1.drawToHand(5)
	f.activePlayer = (f.activePlayer + 1) % len(f.players)
}

func resolveCards(p1, p2 *person) {
	for p1.Hand.Len() != 0 {
		c := p1.Hand.Take(0)
		log.Debugf("Playing %s", c.GetName())

		for _, e := range c.GetPlayEffects() {
			e.Activate(c, p1, p2)
		}
		if base, ok := c.(card.Base); ok {
			p1.Bases.Add(base)
		} else {
			p1.PlayedCards.Add(c)
		}
	}
}

func commitTurnState(p1, p2 *person) {
	// Waste whatever money wasn't spent
	fmt.Printf("Wasted %d trade\n", p1.turnState.Trade)
	fmt.Printf("Applied %d damage to opponent\n", p1.turnState.Combat)
	p2.AddAuthority(-p1.turnState.Combat)
	fmt.Printf("Recovered %d authority\n", p1.turnState.AdditionalAuthority)
	p1.AddAuthority(p1.turnState.AdditionalAuthority)
	return
}

func (p *person) DiscardCards(numToDiscard int) {
	for numToDiscard > 0 && p.Hand.Len() > 0 {
		_ = utils.PickCard("What card would you like to discard?", p.Hand.Cards, true /* required */)

		// get rid of the

	}
}
