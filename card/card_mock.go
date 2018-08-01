package card

import "github.com/valakuzhyk/planetdomains/internal"

// Mock is used to mock a card
type Mock struct {
	Name           string
	GetCostImpl    func() int
	GetFactionImpl func() Faction
	PlayEffectImpl func(internal.Player)
}

func (m Mock) GetName() string {
	return m.Name
}
func (m Mock) GetCost() int {
	return m.GetCostImpl()
}
func (m Mock) GetFaction() Faction {
	return m.GetFactionImpl()
}
func (m Mock) PlayEffect(p internal.Player) {
	m.PlayEffectImpl(p)
}
