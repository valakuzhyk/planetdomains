package card

import "github.com/valakuzhyk/planetdomains/internal"

/* --------------- Trade Federation ---------------- */

type solarSkiff struct{}

func (e solarSkiff) GetName() string                   { return "Solar Skiff" }
func (e solarSkiff) GetCost() int                      { return 1 }
func (e solarSkiff) GetFaction() Faction               { return TRADE_FED }
func (e solarSkiff) PlayEffect(p1, p2 internal.Player) { p1.AddTrade(2) }
func (e solarSkiff) AllyEffect(p1, p2 internal.Player) { p1.DrawCards(1) }

type tradeHauler struct{}

func (e tradeHauler) GetName() string                   { return "Trade Hauler" }
func (e tradeHauler) GetCost() int                      { return 2 }
func (e tradeHauler) GetFaction() Faction               { return TRADE_FED }
func (e tradeHauler) PlayEffect(p1, p2 internal.Player) { p1.AddTrade(3) }
func (e tradeHauler) AllyEffect(p1, p2 internal.Player) { p1.AddAuthority(3) }

type storageSilo struct{}

func (e storageSilo) GetName() string                   { return "Storage Silo" }
func (e storageSilo) GetCost() int                      { return 2 }
func (e storageSilo) GetFaction() Faction               { return TRADE_FED }
func (e storageSilo) PlayEffect(p1, p2 internal.Player) { p1.AddAuthority(2) }
func (e storageSilo) AllyEffect(p1, p2 internal.Player) { p1.AddTrade(2) }
func (e storageSilo) GetDefense() (int, bool)           { return 3, false }

type patrolCutter struct{}

func (e patrolCutter) GetName() string                   { return "Patrol Cutter" }
func (e patrolCutter) GetCost() int                      { return 3 }
func (e patrolCutter) GetFaction() Faction               { return TRADE_FED }
func (e patrolCutter) PlayEffect(p1, p2 internal.Player) { p1.AddTrade(2); p1.AddCombat(3) }
func (e patrolCutter) AllyEffect(p1, p2 internal.Player) { p1.AddAuthority(4) }

type peacekeeper struct{}

func (e peacekeeper) GetName() string                   { return "Peacekeeper" }
func (e peacekeeper) GetCost() int                      { return 6 }
func (e peacekeeper) GetFaction() Faction               { return TRADE_FED }
func (e peacekeeper) PlayEffect(p1, p2 internal.Player) { p1.AddCombat(6); p1.AddAuthority(6) }
func (e peacekeeper) AllyEffect(p1, p2 internal.Player) { p1.DrawCards(1) }

type loyalColony struct{}

func (e loyalColony) GetName() string     { return "Loyal Colony" }
func (e loyalColony) GetCost() int        { return 7 }
func (e loyalColony) GetFaction() Faction { return TRADE_FED }
func (e loyalColony) PlayEffect(p1, p2 internal.Player) {
	p1.AddCombat(3)
	p1.AddAuthority(3)
	p1.AddTrade(3)
}
func (e loyalColony) GetDefense() (int, bool) { return 6, false }

type frontierFerry struct{}

func (e frontierFerry) ScrapEffect(p1, p2 internal.Player) { p1.DestroyBase() }
func (e frontierFerry) GetName() string                    { return "Frontier Ferry" }
func (e frontierFerry) GetCost() int                       { return 4 }
func (e frontierFerry) GetFaction() Faction                { return TRADE_FED }
func (e frontierFerry) PlayEffect(p1, p2 internal.Player)  { p1.AddTrade(3); p1.AddAuthority(4) }

// CARDS TO IMPLEMENT FULLY
// NOT READY CARDS BELOW

type colonySeedShip struct{}

// TODO(add draw effect, if you have played a trade fed card, this goes directly to your hand)
func (e colonySeedShip) GetName() string     { return "Colony Seed Ship" }
func (e colonySeedShip) GetCost() int        { return 5 }
func (e colonySeedShip) GetFaction() Faction { return TRADE_FED }
func (e colonySeedShip) PlayEffect(p1, p2 internal.Player) {
	p1.AddCombat(3)
	p1.AddAuthority(3)
	p1.AddTrade(3)
}

type federationShipyard struct{}

func (e federationShipyard) GetName() string                   { return "Federation Shipyard" }
func (e federationShipyard) GetCost() int                      { return 6 }
func (e federationShipyard) GetFaction() Faction               { return TRADE_FED }
func (e federationShipyard) PlayEffect(p1, p2 internal.Player) { p1.AddTrade(2) }
func (e federationShipyard) AllyEffect(p1, p2 internal.Player) { /* put card on top */ }
func (e federationShipyard) GetDefense() (int, bool)           { return 6, true }

type factoryWorld struct{}

func (e factoryWorld) GetName() string     { return "Factory World" }
func (e factoryWorld) GetCost() int        { return 8 }
func (e factoryWorld) GetFaction() Faction { return TRADE_FED }
func (e factoryWorld) PlayEffect(p1, p2 internal.Player) {
	p1.AddTrade(3) /* put card on top */
}
func (e factoryWorld) GetDefense() (int, bool) { return 6, true }

type centralStation struct{}

func (e centralStation) GetName() string     { return "Central Station" }
func (e centralStation) GetCost() int        { return 4 }
func (e centralStation) GetFaction() Faction { return TRADE_FED }
func (e centralStation) PlayEffect(p1, p2 internal.Player) {
	p1.AddTrade(2)
	// Add abilities that are triggered once if certain conditions are met.
}
func (e centralStation) GetDefense() (int, bool) { return 5, false }
