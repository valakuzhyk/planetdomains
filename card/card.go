package card

import (
	"fmt"

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

// Print returns a string representing the card.
func Print(c Card) string {
	return fmt.Sprintf("%s", c.GetName())
}

// Play plays a card
func Play(c Card, p internal.Player) {
	c.PlayEffect(p)
}

// TrashableCard has an additional affect that applies on trashing
type TrashableCard interface {
	Card
	TrashEffect(internal.Player)
}

// Trash trashes a card
func Trash(c TrashableCard, p internal.Player) {
	c.PlayEffect(p)
	c.TrashEffect(p)
}
