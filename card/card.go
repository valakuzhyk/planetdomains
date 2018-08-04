package card

import (
	"strings"

	"github.com/fatih/color"
)

// Faction represents what group the card is aligned with, and how it gets
// faction bonuses.
type Faction int

const (
	UNKNOWN Faction = iota
	UNALIGNED
	MACHINE_CULT
	BLOB
	STAR_EMPIRE
	TRADE_FED
)

type Ship struct {
	Name     string
	Cost     int
	Factions []Faction

	PlayEffects          []Effect
	AllyEffects          []Effect
	ScrapEffects         []Effect
	playAbilityActivated bool
	allyAbilityActivated bool
}

func (c *Ship) GetName() string { return c.Name }
func (c *Ship) GetCost() int    { return c.Cost }
func (c *Ship) GetFaction() Faction {
	if len(c.Factions) > 0 {
		return c.Factions[0]
	}
	return UNALIGNED
}

func (c *Ship) Reset() {
	c.playAbilityActivated = false
	c.allyAbilityActivated = false
}

// Card represents a card in the game
// TODO: this needs to include the field and other players
type Card interface {
	GetName() string
	GetCost() int
	GetFaction() Faction
	Reset()
	GetPlayEffects() []Effect
	GetAllyEffects() []Effect
	GetScrapEffects() []Effect
	String() string
}

// String returns a string representing the card.
func (c Ship) String() string {
	switch c.GetFaction() {
	case TRADE_FED:
		return color.HiBlueString(c.GetName())
	case BLOB:
		return color.HiGreenString(c.GetName())
	case MACHINE_CULT:
		return color.HiRedString(c.GetName())
	case STAR_EMPIRE:
		return color.HiYellowString(c.GetName())
	}
	return c.GetName()
}

// StringList returns a list of strings representing the cards given.
func StringList(cards ...Card) []string {
	cardlist := make([]string, len(cards))
	for i, c := range cards {
		cardlist[i] = c.String()
	}
	return cardlist
}

func List(cards ...Card) string {
	return strings.Join(StringList(cards...), ", ")
}

func (c *Ship) GetPlayEffects() []Effect {
	return c.PlayEffects
}

func (c *Ship) GetScrapEffects() []Effect {
	return c.ScrapEffects
}

func (c *Ship) GetAllyEffects() []Effect {
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

// Base cards exist until they are destroyed
type Base interface {
	Card
	GetDefense() int
	IsOutpost() bool
}
