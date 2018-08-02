package card

import "github.com/valakuzhyk/planetdomains/internal"

type agingBattleship struct{}

func (e agingBattleship) GetName() string               { return "Aging Battleship" }
func (e agingBattleship) GetCost() int                  { return 5 }
func (e agingBattleship) GetFaction() Faction           { return STAR_EMPIRE }
func (e agingBattleship) PlayEffect(p internal.Player)  { p.AddCombat(5) }
func (e agingBattleship) AllyEffect(p internal.Player)  { p.DrawCards(1) }
func (e agingBattleship) ScrapEffect(p internal.Player) { p.DrawCards(2); p.AddCombat(2) }
func (e agingBattleship) GetDefense() (int, bool)       { return 5, true }

type heavyCruiser struct{}

func (e heavyCruiser) GetName() string              { return "Heavy Cruiser" }
func (e heavyCruiser) GetCost() int                 { return 5 }
func (e heavyCruiser) GetFaction() Faction          { return STAR_EMPIRE }
func (e heavyCruiser) PlayEffect(p internal.Player) { p.AddCombat(4); p.DrawCards(1) }
func (e heavyCruiser) AllyEffect(p internal.Player) { p.DrawCards(1) }

// TO work on

type commandCenter struct{}

func (e commandCenter) GetName() string     { return "Command Center" }
func (e commandCenter) GetCost() int        { return 4 }
func (e commandCenter) GetFaction() Faction { return STAR_EMPIRE }
func (e commandCenter) PlayEffect(p internal.Player) {
	p.AddTrade(2) /* evaluate at end, 2 combat x # of star empire ships */
}
func (e commandCenter) GetDefense() (int, bool) { return 4, true }

type imperialPalace struct{}

func (e imperialPalace) GetName() string     { return "Imperial Palace" }
func (e imperialPalace) GetCost() int        { return 7 }
func (e imperialPalace) GetFaction() Faction { return STAR_EMPIRE }
func (e imperialPalace) PlayEffect(p internal.Player) {
	p.DrawCards(1)
	// Opponent discards card
}
func (e imperialPalace) AllyEffect(p internal.Player) { p.AddCombat(4) }
func (e imperialPalace) GetDefense() (int, bool)      { return 6, true }

type orbitalPlatform struct{}

func (e orbitalPlatform) GetName() string     { return "Orbital Platform" }
func (e orbitalPlatform) GetCost() int        { return 3 }
func (e orbitalPlatform) GetFaction() Faction { return STAR_EMPIRE }
func (e orbitalPlatform) PlayEffect(p internal.Player) {
	// Draw a card if you discard a card
}
func (e orbitalPlatform) AllyEffect(p internal.Player) { p.AddCombat(3) }
func (e orbitalPlatform) GetDefense() (int, bool)      { return 4, false }

type supplyDepot struct{}

func (e supplyDepot) GetName() string     { return "Supply Depot" }
func (e supplyDepot) GetCost() int        { return 6 }
func (e supplyDepot) GetFaction() Faction { return STAR_EMPIRE }
func (e supplyDepot) PlayEffect(p internal.Player) {
	// Discard up to two cards gain 2 money or 2 combat for each card
}
func (e supplyDepot) AllyEffect(p internal.Player) { p.DrawCards(1) }
func (e supplyDepot) GetDefense() (int, bool)      { return 5, true }

type emperorsDreadnaught struct{}

func (e emperorsDreadnaught) GetName() string     { return "Emperor's Dreadnaught" }
func (e emperorsDreadnaught) GetCost() int        { return 8 }
func (e emperorsDreadnaught) GetFaction() Faction { return STAR_EMPIRE }
func (e emperorsDreadnaught) PlayEffect(p internal.Player) {
	p.AddCombat(8)
	p.DrawCards(1)
	// Opponent Discards card
	// If bought and you have played a star empire card, this goes to the hand.
}

type falcon struct{}

func (e falcon) GetName() string               { return "Falcon" }
func (e falcon) GetCost() int                  { return 3 }
func (e falcon) GetFaction() Faction           { return STAR_EMPIRE }
func (e falcon) PlayEffect(p internal.Player)  { p.AddCombat(2); p.DrawCards(1) }
func (e falcon) ScrapEffect(p internal.Player) { /* Target opponent discards a card */ }

type gunship struct{}

func (e gunship) GetName() string               { return "Gunship" }
func (e gunship) GetCost() int                  { return 4 }
func (e gunship) GetFaction() Faction           { return STAR_EMPIRE }
func (e gunship) PlayEffect(p internal.Player)  { p.AddCombat(5) /* opponent discards a card */ }
func (e gunship) ScrapEffect(p internal.Player) { p.AddTrade(4) }

type lancer struct{}

func (e lancer) GetName() string              { return "Lancer" }
func (e lancer) GetCost() int                 { return 2 }
func (e lancer) GetFaction() Faction          { return STAR_EMPIRE }
func (e lancer) PlayEffect(p internal.Player) { p.AddCombat(4) /* If opponent has a base, add 2 */ }
func (e lancer) AllyEffect(p internal.Player) { /* target opponent discards a card */ }

type starBarge struct{}

func (e starBarge) GetName() string              { return "Star Barge" }
func (e starBarge) GetCost() int                 { return 1 }
func (e starBarge) GetFaction() Faction          { return STAR_EMPIRE }
func (e starBarge) PlayEffect(p internal.Player) { p.AddTrade(2) }
func (e starBarge) AllyEffect(p internal.Player) { p.DrawCards(1) /* target opponent discards a card */ }
