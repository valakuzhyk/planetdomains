package card

import "github.com/valakuzhyk/planetdomains/internal"

type frontierStation struct{}

func (e frontierStation) GetName() string     { return "Frontier Station" }
func (e frontierStation) GetCost() int        { return 6 }
func (e frontierStation) GetFaction() Faction { return MACHINE_CULT }
func (e frontierStation) PlayEffect(p internal.Player) {
	// 2 trade or 3 combat
}
func (e frontierStation) GetDefense() (int, bool) { return 6, true }

type stealthTower struct{}

func (e stealthTower) GetName() string     { return "Stealth Tower" }
func (e stealthTower) GetCost() int        { return 5 }
func (e stealthTower) GetFaction() Faction { return MACHINE_CULT }
func (e stealthTower) PlayEffect(p internal.Player) {
	// Copy another base...
}
func (e stealthTower) GetDefense() (int, bool) { return 5, true }

type theIncinerator struct{}

func (e theIncinerator) GetName() string     { return "The Incinerator" }
func (e theIncinerator) GetCost() int        { return 8 }
func (e theIncinerator) GetFaction() Faction { return MACHINE_CULT }
func (e theIncinerator) PlayEffect(p internal.Player) {
	// Scrap up to two cards in hand or discard pile
}
func (e theIncinerator) ScrapEffect(p internal.Player) {
	// 2 damage for each card scrapped from hand or discard pile this turn
}
func (e theIncinerator) GetDefense() (int, bool) { return 6, true }

type theOracle struct{}

func (e theOracle) GetName() string     { return "The Oracle" }
func (e theOracle) GetCost() int        { return 4 }
func (e theOracle) GetFaction() Faction { return MACHINE_CULT }
func (e theOracle) PlayEffect(p internal.Player) {
	// Scrap a card from your hand
}
func (e theOracle) AllyEffect(p internal.Player) { p.AddCombat(3) }
func (e theOracle) GetDefense() (int, bool)      { return 5, true }

type warningBeacon struct{}

func (e warningBeacon) GetName() string     { return "Warning Beacon" }
func (e warningBeacon) GetCost() int        { return 2 }
func (e warningBeacon) GetFaction() Faction { return MACHINE_CULT }
func (e warningBeacon) PlayEffect(p internal.Player) {
	p.AddCombat(2)
	// when bought, if machine cult was played, this goes into your hand.
}
func (e warningBeacon) GetDefense() (int, bool) { return 2, true }

type battleBot struct{}

func (e battleBot) GetName() string     { return "Battle Bot" }
func (e battleBot) GetCost() int        { return 1 }
func (e battleBot) GetFaction() Faction { return MACHINE_CULT }
func (e battleBot) PlayEffect(p internal.Player) {
	p.AddCombat(2)
	// You may scrap a card in your hand
}
func (e battleBot) AllyEffect(p internal.Player) { p.AddCombat(2) }

type convoyBot struct{}

func (e convoyBot) GetName() string     { return "Convoy Bot" }
func (e convoyBot) GetCost() int        { return 3 }
func (e convoyBot) GetFaction() Faction { return MACHINE_CULT }
func (e convoyBot) PlayEffect(p internal.Player) {
	p.AddCombat(4)
	// You may scrap a card in your hand or discard
}
func (e convoyBot) AllyEffect(p internal.Player) { p.AddCombat(2) }

type mechCruiser struct{}

func (e mechCruiser) GetName() string     { return "Mech Cruiser" }
func (e mechCruiser) GetCost() int        { return 5 }
func (e mechCruiser) GetFaction() Faction { return MACHINE_CULT }
func (e mechCruiser) PlayEffect(p internal.Player) {
	p.AddCombat(6)
	// You may scrap a card in your hand or discard
}
func (e mechCruiser) AllyEffect(p internal.Player) { /* destroy a target base */ }

type miningMech struct{}

func (e miningMech) GetName() string     { return "Mining Mech" }
func (e miningMech) GetCost() int        { return 4 }
func (e miningMech) GetFaction() Faction { return MACHINE_CULT }
func (e miningMech) PlayEffect(p internal.Player) {
	p.AddTrade(3)
	// You may scrap a card in your hand or discard
}
func (e miningMech) AllyEffect(p internal.Player) { p.AddCombat(3) }

type repairBot struct{}

func (e repairBot) GetName() string     { return "Repair Bot" }
func (e repairBot) GetCost() int        { return 2 }
func (e repairBot) GetFaction() Faction { return MACHINE_CULT }
func (e repairBot) PlayEffect(p internal.Player) {
	p.AddTrade(2)
	// You may scrap a card in your discard
}
func (e repairBot) ScrapEffect(p internal.Player) { p.AddCombat(2) }

type theWrecker struct{}

func (e theWrecker) GetName() string     { return "The Wrecker" }
func (e theWrecker) GetCost() int        { return 7 }
func (e theWrecker) GetFaction() Faction { return MACHINE_CULT }
func (e theWrecker) PlayEffect(p internal.Player) {
	p.AddCombat(6)
	// You may scrap up to 2 cards in your discard and or hand
}
func (e theWrecker) ScrapEffect(p internal.Player) { p.DrawCards(1) }
