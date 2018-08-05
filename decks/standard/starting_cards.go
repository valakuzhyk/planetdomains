package standard

import (
	"github.com/valakuzhyk/planetdomains/card"
	"github.com/valakuzhyk/planetdomains/cardimpl"
	"github.com/valakuzhyk/planetdomains/effect"
)

func scout() card.Card {
	return &cardimpl.Ship{
		Name:        "Scout",
		Cost:        0,
		Factions:    []card.Faction{},
		PlayEffects: []card.Effect{effect.AddTrade{1}},
	}
}

func viper() card.Card {
	return &cardimpl.Ship{
		Name:        "Viper",
		Cost:        0,
		Factions:    []card.Faction{},
		PlayEffects: []card.Effect{effect.AddCombat{1}},
	}
}

// StarterDeck is what a player normally starts with at the beginning of the game
func StarterDeck() card.Deck {
	d := cardimpl.NewDeck()
	for i := 0; i < 8; i++ {
		d.PlaceOnTop(scout())
	}
	for i := 0; i < 2; i++ {
		d.PlaceOnTop(viper())
	}
	d.Shuffle()
	return d
}

func explorer() card.Card {
	return &cardimpl.Ship{
		Name:         "Explorer",
		Cost:         2,
		Factions:     []card.Faction{},
		PlayEffects:  []card.Effect{effect.AddTrade{2}},
		ScrapEffects: []card.Effect{effect.AddCombat{2}},
	}
}

// NewExplorerDeck returns a new deck containing 10 explorers
func NewExplorerDeck() card.Deck {
	d := cardimpl.NewDeck()
	for i := 0; i < 10; i++ {
		d.PlaceOnTop(explorer())
	}
	return d
}
