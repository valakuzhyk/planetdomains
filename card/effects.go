package card

import (
	"fmt"

	"github.com/valakuzhyk/planetdomains/internal"
)

// Effect is the result of playing a card. These can be registered with a player,
// or, after playing a ship may be triggered immediately.
type Effect interface {
	Activate(c Card, p internal.Player)
	String() string
}

// AddCombat adds the amount of damage a player does to an opponent.
type AddCombat struct{ Amount int }

func (e AddCombat) Activate(c Card, p internal.Player) { p.AddCombat(e.Amount) }
func (e AddCombat) String() string                     { return fmt.Sprintf("Add %d combat", e.Amount) }

type AddAuthority struct{ Amount int }

func (e AddAuthority) Activate(c Card, p internal.Player) { p.AddAuthority(e.Amount) }
func (e AddAuthority) String() string                     { return fmt.Sprintf("Add %d authority", e.Amount) }

type AddTrade struct{ Amount int }

func (e AddTrade) Activate(c Card, p internal.Player) { p.AddTrade(e.Amount) }
func (e AddTrade) String() string                     { return fmt.Sprintf("Add %d trade", e.Amount) }

type DestroyBase struct{ Amount int }

func (e DestroyBase) Activate(c Card, p internal.Player) {
	for i := 0; i < e.Amount; i++ {
		p.DestroyBase()
	}
}
func (e DestroyBase) String() string { return fmt.Sprintf("Destroy %d base(s)", e.Amount) }

type AcquireCardLessThan struct{ Cost int }

func (e AcquireCardLessThan) Activate(c Card, p internal.Player) {
	// TODO
}
func (e AcquireCardLessThan) String() string {
	return fmt.Sprintf("Acquire card from the trade row less than %d", e.Cost)
}

type DrawCards struct{ Amount uint }

func (e DrawCards) Activate(c Card, p internal.Player) {
	p.DrawCards(e.Amount)
}
func (e DrawCards) String() string {
	return fmt.Sprintf("Draw %d cards", e.Amount)
}

type OpponentDiscards struct{ Amount uint }

func (e OpponentDiscards) Activate(c Card, p internal.Player) {
	// TODO
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

type ScrapCards struct {
	Amount   uint
	Location ScrapLocation
	Required bool
}

func (e ScrapCards) Activate(c Card, p internal.Player) {
	// TODO
}
func (e ScrapCards) String() string {
	// TODO location
	return fmt.Sprintf("Scrap %d cards", e.Amount)
}
