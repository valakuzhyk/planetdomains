package card

import (
	"github.com/fatih/color"
	"github.com/valakuzhyk/planetdomains/internal"
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

// Card represents a card in the game
// TODO: this needs to include the field and other players
type Card interface {
	GetName() string
	GetCost() int
	GetFaction() Faction
	PlayEffect(internal.Player)
}

// String returns a string representing the card.
func String(c Card) string {
	switch c.GetFaction() {
	case TRADE_FED:
		return color.BlueString(c.GetName())
	case BLOB:
		return color.GreenString(c.GetName())
	case MACHINE_CULT:
		return color.RedString(c.GetName())
	case STAR_EMPIRE:
		return color.YellowString(c.GetName())
	}
	return c.GetName()
}

// Play plays a card
func Play(c Card, p internal.Player) {
	c.PlayEffect(p)
}

// ScrapableCard has an additional effect that applies on trashing
type ScrapableCard interface {
	Card
	ScrapEffect(internal.Player)
}

// AllyableCard has an additional effect that applies when another card
// of the same faction is played
type AllyableCard interface {
	Card
	AllyEffect(internal.Player)
}

// Base cards exist until they are destroyed
type Base interface {
	Card
	GetDefense() (strength int, isOutpost bool)
}
