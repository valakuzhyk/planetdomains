package effect

import (
	"fmt"
	"strings"

	"github.com/valakuzhyk/planetdomains/card"
	"github.com/valakuzhyk/planetdomains/internal"
)

// AddCombat adds the amount of damage a player does to an opponent.
type AddCombat struct{ Amount int }

func (e AddCombat) Activate(c card.Card, p1, p2 internal.Player) { p1.AddCombat(e.Amount) }
func (e AddCombat) String() string                               { return fmt.Sprintf("Add %d combat", e.Amount) }

type AddAuthority struct{ Amount int }

func (e AddAuthority) Activate(c card.Card, p1, p2 internal.Player) { p1.AddAuthority(e.Amount) }
func (e AddAuthority) String() string                               { return fmt.Sprintf("Add %d authority", e.Amount) }

type AddTrade struct{ Amount int }

func (e AddTrade) Activate(c card.Card, p1, p2 internal.Player) { p1.AddTrade(e.Amount) }
func (e AddTrade) String() string                               { return fmt.Sprintf("Add %d trade", e.Amount) }

type DestroyBase struct{}

func (e DestroyBase) Activate(c card.Card, p1, p2 internal.Player) { p1.DestroyBase(p2) }
func (e DestroyBase) String() string                               { return "Destroy opponent base" }

type AcquireCardLessThan struct{ Cost int }

func (e AcquireCardLessThan) Activate(c card.Card, p1, p2 internal.Player) {
	// TODO
}
func (e AcquireCardLessThan) String() string {
	return fmt.Sprintf("Acquire card from the trade row less than %d", e.Cost)
}

type DrawCards struct{ Amount uint }

func (e DrawCards) Activate(c card.Card, p1, p2 internal.Player) {
	p1.DrawCards(e.Amount)
}
func (e DrawCards) String() string {
	return fmt.Sprintf("Draw %d cards", e.Amount)
}

type OpponentDiscards struct{ Amount uint }

func (e OpponentDiscards) Activate(c card.Card, p1, p2 internal.Player) {
	for i := uint(0); i < e.Amount; i++ {
		p2.MustDiscard()
	}
}
func (e OpponentDiscards) String() string {
	return fmt.Sprintf("Opponent discards %d cards", e.Amount)
}

type ScrapLocation int

const (
	ScrapLocation_UNKNOWN ScrapLocation = iota
	ScrapLocation_HAND
	ScrapLocation_DISCARD
	ScrapLocation_HAND_OR_DISCARD
	ScrapLocation_TRADE_ROW
)

func (s ScrapLocation) String() string {
	switch s {
	case ScrapLocation_HAND:
		return "hand"
	case ScrapLocation_DISCARD:
		return "discard"
	case ScrapLocation_HAND_OR_DISCARD:
		return "hand or discard"
	case ScrapLocation_TRADE_ROW:
		return "trade row"
	}
	return "Unknown??"
}

type ScrapCards struct {
	Amount   uint
	Location ScrapLocation
	Required bool
}

func (e ScrapCards) Activate(c card.Card, p1, p2 internal.Player) {
	switch e.Location {
	case ScrapLocation_DISCARD:
		p1.ScrapFromDiscard()
	case ScrapLocation_HAND:
		p1.ScrapFromHand()
	case ScrapLocation_HAND_OR_DISCARD:
	case ScrapLocation_TRADE_ROW:
	}
}
func (e ScrapCards) String() string {
	return fmt.Sprintf("Scrap %d cards from %s", e.Amount, e.Location.String())
}

type ChooseBetween struct {
	Options []card.Effect
}

func (e ChooseBetween) Activate(c card.Card, p1, p2 internal.Player) {
	// TODO
}
func (e ChooseBetween) String() string {
	strs := []string{}
	for _, e := range e.Options {
		strs = append(strs, e.String())
	}
	return strings.Join(strs, " or ")
}
