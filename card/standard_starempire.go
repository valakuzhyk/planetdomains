package card

import "github.com/valakuzhyk/planetdomains/internal"

// TO work on

type commandCenter struct{}

func (e commandCenter) GetName() string     { return "Command Center" }
func (e commandCenter) GetCost() int        { return 4 }
func (e commandCenter) GetFaction() Faction { return TRADE_FED }
func (e commandCenter) PlayEffect(p internal.Player) {
	p.AddTrade(2) /* evaluate at end, 2 combat x # of star empire ships */
}
func (e commandCenter) GetDefense() (int, bool) { return 4, true }

type imperialPalace struct{}

func (e imperialPalace) GetName() string     { return "Command Center" }
func (e imperialPalace) GetCost() int        { return 7 }
func (e imperialPalace) GetFaction() Faction { return TRADE_FED }
func (e imperialPalace) PlayEffect(p internal.Player) {
	p.DrawCards(1)
	// Opponent discards card
}
func (e imperialPalace) AllyEffect(p internal.Player) { p.AddCombat(4) }
func (e imperialPalace) GetDefense() (int, bool)      { return 6, true }
