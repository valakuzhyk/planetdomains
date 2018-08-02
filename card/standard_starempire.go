package card

import "github.com/valakuzhyk/planetdomains/internal"

type agingBattleship struct{}

func (e agingBattleship) GetName() string                    { return "Aging Battleship" }
func (e agingBattleship) GetCost() int                       { return 5 }
func (e agingBattleship) GetFaction() Faction                { return STAR_EMPIRE }
func (e agingBattleship) PlayEffect(p1, p2 internal.Player)  { p1.AddCombat(5) }
func (e agingBattleship) AllyEffect(p1, p2 internal.Player)  { p1.DrawCards(1) }
func (e agingBattleship) ScrapEffect(p1, p2 internal.Player) { p1.DrawCards(2); p1.AddCombat(2) }
func (e agingBattleship) GetDefense() (int, bool)            { return 5, true }

type heavyCruiser struct{}

func (e heavyCruiser) GetName() string                   { return "Heavy Cruiser" }
func (e heavyCruiser) GetCost() int                      { return 5 }
func (e heavyCruiser) GetFaction() Faction               { return STAR_EMPIRE }
func (e heavyCruiser) PlayEffect(p1, p2 internal.Player) { p1.AddCombat(4); p1.DrawCards(1) }
func (e heavyCruiser) AllyEffect(p1, p2 internal.Player) { p1.DrawCards(1) }

// TO work on

type commandCenter struct{}

func (e commandCenter) GetName() string     { return "Command Center" }
func (e commandCenter) GetCost() int        { return 4 }
func (e commandCenter) GetFaction() Faction { return STAR_EMPIRE }
func (e commandCenter) PlayEffect(p1, p2 internal.Player) {
	p1.AddTrade(2) /* evaluate at end, 2 combat x # of star empire ships */
}
func (e commandCenter) GetDefense() (int, bool) { return 4, true }

type imperialPalace struct{}

func (e imperialPalace) GetName() string     { return "Imperial Palace" }
func (e imperialPalace) GetCost() int        { return 7 }
func (e imperialPalace) GetFaction() Faction { return STAR_EMPIRE }
func (e imperialPalace) PlayEffect(p1, p2 internal.Player) {
	p1.DrawCards(1)
	p2.DiscardCard()
}
func (e imperialPalace) AllyEffect(p1, p2 internal.Player) { p1.AddCombat(4) }
func (e imperialPalace) GetDefense() (int, bool)           { return 6, true }

type orbitalPlatform struct{}

func (e orbitalPlatform) GetName() string     { return "Orbital Platform" }
func (e orbitalPlatform) GetCost() int        { return 3 }
func (e orbitalPlatform) GetFaction() Faction { return STAR_EMPIRE }
func (e orbitalPlatform) PlayEffect(p1, p2 internal.Player) {
	// Draw a card if you discard a card
}
func (e orbitalPlatform) AllyEffect(p1, p2 internal.Player) { p1.AddCombat(3) }
func (e orbitalPlatform) GetDefense() (int, bool)           { return 4, false }

type supplyDepot struct{}

func (e supplyDepot) GetName() string     { return "Supply Depot" }
func (e supplyDepot) GetCost() int        { return 6 }
func (e supplyDepot) GetFaction() Faction { return STAR_EMPIRE }
func (e supplyDepot) PlayEffect(p1, p2 internal.Player) {
	// Discard up to two cards gain 2 money or 2 combat for each card
}
func (e supplyDepot) AllyEffect(p1, p2 internal.Player) { p1.DrawCards(1) }
func (e supplyDepot) GetDefense() (int, bool)           { return 5, true }

type emperorsDreadnaught struct{}

func (e emperorsDreadnaught) GetName() string     { return "Emperor's Dreadnaught" }
func (e emperorsDreadnaught) GetCost() int        { return 8 }
func (e emperorsDreadnaught) GetFaction() Faction { return STAR_EMPIRE }
func (e emperorsDreadnaught) PlayEffect(p1, p2 internal.Player) {
	p1.AddCombat(8)
	p1.DrawCards(1)
	p2.DiscardCard()
	// If bought and you have played a star empire card, this goes to the hand.
}

type falcon struct{}

func (e falcon) GetName() string                    { return "Falcon" }
func (e falcon) GetCost() int                       { return 3 }
func (e falcon) GetFaction() Faction                { return STAR_EMPIRE }
func (e falcon) PlayEffect(p1, p2 internal.Player)  { p1.AddCombat(2); p1.DrawCards(1) }
func (e falcon) ScrapEffect(p1, p2 internal.Player) { p2.DiscardCard() }

type gunship struct{}

func (e gunship) GetName() string                    { return "Gunship" }
func (e gunship) GetCost() int                       { return 4 }
func (e gunship) GetFaction() Faction                { return STAR_EMPIRE }
func (e gunship) PlayEffect(p1, p2 internal.Player)  { p1.AddCombat(5); p2.DiscardCard() }
func (e gunship) ScrapEffect(p1, p2 internal.Player) { p1.AddTrade(4) }

type lancer struct{}

func (e lancer) GetName() string                   { return "Lancer" }
func (e lancer) GetCost() int                      { return 2 }
func (e lancer) GetFaction() Faction               { return STAR_EMPIRE }
func (e lancer) PlayEffect(p1, p2 internal.Player) { p1.AddCombat(4) /* If opponent has a base, add 2 */ }
func (e lancer) AllyEffect(p1, p2 internal.Player) { p2.DiscardCard() }

type starBarge struct{}

func (e starBarge) GetName() string                   { return "Star Barge" }
func (e starBarge) GetCost() int                      { return 1 }
func (e starBarge) GetFaction() Faction               { return STAR_EMPIRE }
func (e starBarge) PlayEffect(p1, p2 internal.Player) { p1.AddTrade(2) }
func (e starBarge) AllyEffect(p1, p2 internal.Player) { p1.DrawCards(1); p2.DiscardCard() }
