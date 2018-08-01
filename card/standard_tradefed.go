package card

import "github.com/valakuzhyk/planetdomains/internal"

/* --------------- Trade Federation ---------------- */

type solarSkiff struct{}

func (e solarSkiff) GetName() string              { return "Solar Skiff" }
func (e solarSkiff) GetCost() int                 { return 1 }
func (e solarSkiff) GetFaction() Faction          { return TRADE_FED }
func (e solarSkiff) PlayEffect(p internal.Player) { p.AddTrade(2) }
func (e solarSkiff) AllyEffect(p internal.Player) { p.DrawCards(1) }

type tradeHauler struct{}

func (e tradeHauler) GetName() string              { return "Trade Hauler" }
func (e tradeHauler) GetCost() int                 { return 2 }
func (e tradeHauler) GetFaction() Faction          { return TRADE_FED }
func (e tradeHauler) PlayEffect(p internal.Player) { p.AddTrade(3) }
func (e tradeHauler) AllyEffect(p internal.Player) { p.AddAuthority(3) }

type storageSilo struct{}

func (e storageSilo) GetName() string              { return "Storage Silo" }
func (e storageSilo) GetCost() int                 { return 2 }
func (e storageSilo) GetFaction() Faction          { return TRADE_FED }
func (e storageSilo) PlayEffect(p internal.Player) { p.AddAuthority(2) }
func (e storageSilo) AllyEffect(p internal.Player) { p.AddTrade(2) }
func (e storageSilo) GetDefense() (int, bool)      { return 3, false }

type patrolCutter struct{}

func (e patrolCutter) GetName() string              { return "Patrol Cutter" }
func (e patrolCutter) GetCost() int                 { return 3 }
func (e patrolCutter) GetFaction() Faction          { return TRADE_FED }
func (e patrolCutter) PlayEffect(p internal.Player) { p.AddTrade(2); p.AddCombat(3) }
func (e patrolCutter) AllyEffect(p internal.Player) { p.AddAuthority(4) }

type peacekeeper struct{}

func (e peacekeeper) GetName() string              { return "Peacekeeper" }
func (e peacekeeper) GetCost() int                 { return 6 }
func (e peacekeeper) GetFaction() Faction          { return TRADE_FED }
func (e peacekeeper) PlayEffect(p internal.Player) { p.AddCombat(6); p.AddAuthority(6) }
func (e peacekeeper) AllyEffect(p internal.Player) { p.DrawCards(1) }

type loyalColony struct{}

func (e loyalColony) GetName() string     { return "Loyal Colony" }
func (e loyalColony) GetCost() int        { return 7 }
func (e loyalColony) GetFaction() Faction { return TRADE_FED }
func (e loyalColony) PlayEffect(p internal.Player) {
	p.AddCombat(3)
	p.AddAuthority(3)
	p.AddTrade(3)
}
func (e loyalColony) GetDefense() (int, bool) { return 6, false }

// CARDS TO IMPLEMENT FULLY
// NOT READY CARDS BELOW

type frontierFerry struct{}

// func (e frontierFerry) ScrapEffect(p internal.Player) { p.DestroyOpponentBase() }
func (e frontierFerry) GetName() string              { return "Frontier Ferry" }
func (e frontierFerry) GetCost() int                 { return 4 }
func (e frontierFerry) GetFaction() Faction          { return TRADE_FED }
func (e frontierFerry) PlayEffect(p internal.Player) { p.AddTrade(3); p.AddAuthority(4) }

type colonySeedShip struct{}

// TODO(add draw effect, if you have played a trade fed card, this goes directly to your hand)
func (e colonySeedShip) GetName() string     { return "Colony Seed Ship" }
func (e colonySeedShip) GetCost() int        { return 5 }
func (e colonySeedShip) GetFaction() Faction { return TRADE_FED }
func (e colonySeedShip) PlayEffect(p internal.Player) {
	p.AddCombat(3)
	p.AddAuthority(3)
	p.AddTrade(3)
}

type federationShipyard struct{}

func (e federationShipyard) GetName() string              { return "Federation Shipyard" }
func (e federationShipyard) GetCost() int                 { return 6 }
func (e federationShipyard) GetFaction() Faction          { return TRADE_FED }
func (e federationShipyard) PlayEffect(p internal.Player) { p.AddTrade(2) }
func (e federationShipyard) AllyEffect(p internal.Player) { /* put card on top */ }
func (e federationShipyard) GetDefense() (int, bool)      { return 6, true }

type factoryWorld struct{}

func (e factoryWorld) GetName() string     { return "Factory World" }
func (e factoryWorld) GetCost() int        { return 8 }
func (e factoryWorld) GetFaction() Faction { return TRADE_FED }
func (e factoryWorld) PlayEffect(p internal.Player) {
	p.AddTrade(3) /* put card on top */
}
func (e factoryWorld) GetDefense() (int, bool) { return 6, true }

type centralStation struct{}

func (e centralStation) GetName() string     { return "Central Station" }
func (e centralStation) GetCost() int        { return 4 }
func (e centralStation) GetFaction() Faction { return TRADE_FED }
func (e centralStation) PlayEffect(p internal.Player) {
	p.AddTrade(2)
	// Add abilities that are triggered once if certain conditions are met.
}
func (e centralStation) GetDefense() (int, bool) { return 5, false }

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
		commandCenter{}:      2,
		imperialPalace{}:     1,
	}

	for object, count := range counts {
		for i := 0; i < count; i++ {
			d.PlaceOnTop(object)
		}
	}

	d.Shuffle()

	return d
}
