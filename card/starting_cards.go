package card

func scout() Card {
	return &Ship{
		Name:        "Scout",
		Cost:        0,
		Factions:    []Faction{},
		PlayEffects: []Effect{AddTrade{1}},
	}
}

func viper() Card {
	return &Ship{
		Name:        "Viper",
		Cost:        0,
		Factions:    []Faction{},
		PlayEffects: []Effect{AddCombat{1}},
	}
}

// DefaultStarterDeck is what a player normally starts with at the beginning of the game
func DefaultStarterDeck() Deck {
	d := &deckImpl{}
	for i := 0; i < 8; i++ {
		d.PlaceOnTop(scout())
	}
	for i := 0; i < 2; i++ {
		d.PlaceOnTop(viper())
	}
	d.Shuffle()
	return d
}

func explorer() Card {
	return &Ship{
		Name:         "Explorer",
		Cost:         2,
		Factions:     []Faction{},
		PlayEffects:  []Effect{AddTrade{2}},
		ScrapEffects: []Effect{AddCombat{2}},
	}
}

// NewExplorerDeck returns a new deck containing 10 explorers
func NewExplorerDeck() Deck {
	d := &deckImpl{}
	for i := 0; i < 10; i++ {
		d.PlaceOnTop(explorer())
	}
	return d
}

// StandardDeck constructs the cards for a standardDeck
func StandardDeck() Deck {
	d := NewDeck()

	counts := map[Card]int{
		// Trade Federation Cards
		centralStation():     2,
		factoryWorld():       1,
		federationShipyard(): 1,
		loyalColony():        1,
		storageSilo():        2,
		colonySeedShip():     1,
		frontierFerry():      2,
		patrolCutter():       3,
		peacekeeper():        1,
		solarSkiff():         3,
		tradeHauler():        3,

		// Star Empire Cards
		commandCenter():       2,
		imperialPalace():      1,
		orbitalPlatform():     3,
		supplyDepot():         1,
		agingBattleship():     1,
		emperorsDreadnaught(): 1,
		falcon():              2,
		gunship():             2,
		heavyCruiser():        1,
		lancer():              3,
		starBarge():           3,

		// Machine Cult Cards
		frontierStation(): 1,
		stealthTower():    1,
		theIncinerator():  1,
		theOracle():       1,
		warningBeacon():   3,
		battleBot():       3,
		convoyBot():       3,
		mechCruiser():     1,
		miningMech():      2,
		repairBot():       3,
		theWrecker():      1,

		// Blob cards
		bioformer():   2,
		plasmaVent():  1,
		stellarReef(): 3,
		cargoPod():    3,
		leviathan():   1,
		moonwurm():    1,
		parasite():    1,
		ravager():     2,
		predator():    3,
		swarmer():     3,
	}

	for object, count := range counts {
		for i := 0; i < count; i++ {
			d.PlaceOnTop(object)
		}
	}

	d.Shuffle()

	return d
}
