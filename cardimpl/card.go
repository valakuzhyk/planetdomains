package cardimpl

import (
	"github.com/fatih/color"
	"github.com/valakuzhyk/planetdomains/card"
)

type Ship struct {
	Name     string
	Cost     int
	Factions []card.Faction

	PlayEffects          []card.Effect
	AllyEffects          []card.Effect
	ScrapEffects         []card.Effect
	playAbilityActivated bool
	allyAbilityActivated bool
}

func (c *Ship) GetName() string { return c.Name }
func (c *Ship) GetCost() int    { return c.Cost }
func (c *Ship) GetFaction() card.Faction {
	if len(c.Factions) > 0 {
		return c.Factions[0]
	}
	return card.UNALIGNED
}

func (c *Ship) Reset() {
	c.playAbilityActivated = false
	c.allyAbilityActivated = false
}

// String returns a string representing the card.
func (c Ship) String() string {
	switch c.GetFaction() {
	case card.TRADE_FED:
		return color.HiBlueString(c.GetName())
	case card.BLOB:
		return color.HiGreenString(c.GetName())
	case card.MACHINE_CULT:
		return color.HiRedString(c.GetName())
	case card.STAR_EMPIRE:
		return color.HiYellowString(c.GetName())
	}
	return c.GetName()
}

func (c *Ship) GetPlayEffects() []card.Effect {
	return c.PlayEffects
}

func (c *Ship) GetScrapEffects() []card.Effect {
	return c.ScrapEffects
}

func (c *Ship) GetAllyEffects() []card.Effect {
	return c.AllyEffects
}

type BaseImpl struct {
	Ship
	Defense       int
	IsBaseOutpost bool
}

func (b *BaseImpl) GetDefense() int {
	return b.Defense
}

func (b *BaseImpl) IsOutpost() bool {
	return b.IsBaseOutpost
}
