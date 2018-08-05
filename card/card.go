package card

import (
	"github.com/valakuzhyk/planetdomains/internal"
)

// Effect is the result of playing a card. These can be registered with a player,
// or, after playing a ship may be triggered immediately.
type Effect interface {
	Activate(c Card, p1, p2 internal.Player)
	String() string
}

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
	Reset()
	GetPlayEffects() []Effect
	GetAllyEffects() []Effect
	GetScrapEffects() []Effect
	String() string
}

// Base cards exist until they are destroyed
type Base interface {
	Card
	GetDefense() int
	IsOutpost() bool
}

// StringList returns a list of strings representing the cards given.
func StringList(cards ...Card) []string {
	cardlist := make([]string, len(cards))
	for i, c := range cards {
		cardlist[i] = c.String()
	}
	return cardlist
}
