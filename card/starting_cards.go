package card

import "github.com/valakuzhyk/planetdomains/internal"

type scout struct{}

func (e scout) GetName() string                   { return "Scout" }
func (e scout) GetCost() int                      { return 0 }
func (e scout) GetFaction() Faction               { return UNALIGNED }
func (e scout) PlayEffect(p1, p2 internal.Player) { p1.AddTrade(1) }

type viper struct{}

func (e viper) GetName() string                   { return "Viper" }
func (e viper) GetCost() int                      { return 0 }
func (e viper) GetFaction() Faction               { return UNALIGNED }
func (e viper) PlayEffect(p1, p2 internal.Player) { p1.AddCombat(1) }

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

type explorer struct{}

func (e explorer) GetName() string                   { return "Explorer" }
func (e explorer) GetCost() int                      { return 2 }
func (e explorer) GetFaction() Faction               { return UNALIGNED }
func (e explorer) PlayEffect(p1, p2 internal.Player) { p1.AddTrade(2) }

// NewExplorerDeck returns a new deck containing 10 explorers
func NewExplorerDeck() Deck {
	d := &deckImpl{}
	for i := 0; i < 10; i++ {
		d.PlaceOnTop(&explorer{})
	}
	return d
}

// StandardDeck constructs the cards for a standardDeck
func StandardDeck() Deck {
	d := NewDeck()

	counts := map[Card]int{
		// Trade Federation Cards
		centralStation{}:     2,
		factoryWorld{}:       1,
		federationShipyard{}: 1,
		loyalColony{}:        1,
		storageSilo{}:        2,
		colonySeedShip{}:     1,
		frontierFerry{}:      2,
		patrolCutter{}:       3,
		peacekeeper{}:        1,
		solarSkiff{}:         3,
		tradeHauler{}:        3,

		// Star Empire Cards
		commandCenter{}:       2,
		imperialPalace{}:      1,
		orbitalPlatform{}:     3,
		supplyDepot{}:         1,
		agingBattleship{}:     1,
		emperorsDreadnaught{}: 1,
		falcon{}:              2,
		gunship{}:             2,
		heavyCruiser{}:        1,
		lancer{}:              3,
		starBarge{}:           3,

		// Machine Cult Cards
		frontierStation{}: 1,
		stealthTower{}:    1,
		theIncinerator{}:  1,
		theOracle{}:       1,
		warningBeacon{}:   3,
		battleBot{}:       3,
		convoyBot{}:       3,
		mechCruiser{}:     1,
		miningMech{}:      2,
		repairBot{}:       3,
		theWrecker{}:      1,

		// Blob cards
		bioformer{}:   2,
		plasmaVent{}:  1,
		stellarReef{}: 3,
		cargoPod{}:    3,
		leviathan{}:   1,
		moonwurm{}:    1,
		parasite{}:    1,
		ravager{}:     2,
		predator{}:    3,
		swarmer{}:     3,
	}

	for object, count := range counts {
		for i := 0; i < count; i++ {
			d.PlaceOnTop(object)
		}
	}

	d.Shuffle()

	return d
}
