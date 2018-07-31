package card

import "github.com/valakuzhyk/planetdomains/internal"

// Scout

type scout struct{}

func (e scout) GetName() string              { return "Scout" }
func (e scout) GetCost() int                 { return 0 }
func (e scout) GetFaction() Faction          { return UNALIGNED }
func (e scout) PlayEffect(p internal.Player) { p.AddTrade(1) }

// Viper
type viper struct{}

func (e viper) GetName() string              { return "Viper" }
func (e viper) GetCost() int                 { return 0 }
func (e viper) GetFaction() Faction          { return UNALIGNED }
func (e viper) PlayEffect(p internal.Player) { p.AddCombat(1) }

// DefaultStarterDeck is what a player normally starts with at the beginning of the game
func DefaultStarterDeck() Deck {
	d := &deckImpl{}
	for i := 0; i < 8; i++ {
		d.PlaceOnTop(scout{})
	}
	for i := 0; i < 2; i++ {
		d.PlaceOnTop(viper{})
	}
	d.Shuffle()
	return d
}

// Explorers

type explorer struct{}

func (e explorer) GetName() string              { return "Explorer" }
func (e explorer) GetCost() int                 { return 2 }
func (e explorer) GetFaction() Faction          { return UNALIGNED }
func (e explorer) PlayEffect(p internal.Player) { p.AddTrade(2) }

// NewExplorerDeck returns a new deck containing 10 explorers
func NewExplorerDeck() Deck {
	d := &deckImpl{}
	for i := 0; i < 10; i++ {
		d.PlaceOnTop(&explorer{})
	}
	return d
}
