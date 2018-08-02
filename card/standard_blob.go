package card

import "github.com/valakuzhyk/planetdomains/internal"

type bioformer struct{}

func (e bioformer) GetName() string                    { return "Bioformer" }
func (e bioformer) GetCost() int                       { return 4 }
func (e bioformer) GetFaction() Faction                { return BLOB }
func (e bioformer) PlayEffect(p1, p2 internal.Player)  { p1.AddCombat(3) }
func (e bioformer) ScrapEffect(p1, p2 internal.Player) { p1.AddTrade(3) }
func (e bioformer) GetDefense() (int, bool)            { return 4, false }

type plasmaVent struct{}

func (e plasmaVent) GetName() string     { return "Plasma Vent" }
func (e plasmaVent) GetCost() int        { return 6 }
func (e plasmaVent) GetFaction() Faction { return BLOB }
func (e plasmaVent) PlayEffect(p1, p2 internal.Player) {
	p1.AddCombat(4) /* if you aquired this after playing a blob, direct to hand */
}
func (e plasmaVent) ScrapEffect(p1, p2 internal.Player) { p2.DestroyBase() }
func (e plasmaVent) GetDefense() (int, bool)            { return 5, false }

type stellarReef struct{}

func (e stellarReef) GetName() string                    { return "Stellar Reef" }
func (e stellarReef) GetCost() int                       { return 2 }
func (e stellarReef) GetFaction() Faction                { return BLOB }
func (e stellarReef) PlayEffect(p1, p2 internal.Player)  { p1.AddTrade(1) }
func (e stellarReef) ScrapEffect(p1, p2 internal.Player) { p1.AddCombat(3) }
func (e stellarReef) GetDefense() (int, bool)            { return 3, false }

type cargoPod struct{}

func (e cargoPod) GetName() string                    { return "Cargo Pod" }
func (e cargoPod) GetCost() int                       { return 3 }
func (e cargoPod) GetFaction() Faction                { return BLOB }
func (e cargoPod) PlayEffect(p1, p2 internal.Player)  { p1.AddTrade(3) }
func (e cargoPod) AllyEffect(p1, p2 internal.Player)  { p1.AddCombat(3) }
func (e cargoPod) ScrapEffect(p1, p2 internal.Player) { p1.AddCombat(3) }

type leviathan struct{}

func (e leviathan) GetName() string     { return "Leviathan" }
func (e leviathan) GetCost() int        { return 8 }
func (e leviathan) GetFaction() Faction { return BLOB }
func (e leviathan) PlayEffect(p1, p2 internal.Player) {
	p1.AddCombat(9)
	p1.DrawCards(1)
	p2.DestroyBase()
}
func (e leviathan) AllyEffect(p1, p2 internal.Player) { p1.AcquireCardLessThan(3) }

type moonwurm struct{}

func (e moonwurm) GetName() string                   { return "Moonwurm" }
func (e moonwurm) GetCost() int                      { return 7 }
func (e moonwurm) GetFaction() Faction               { return BLOB }
func (e moonwurm) PlayEffect(p1, p2 internal.Player) { p1.AddCombat(8); p1.DrawCards(1) }
func (e moonwurm) AllyEffect(p1, p2 internal.Player) { p1.AcquireCardLessThan(2) }

type parasite struct{}

func (e parasite) GetName() string     { return "Parasite" }
func (e parasite) GetCost() int        { return 5 }
func (e parasite) GetFaction() Faction { return BLOB }
func (e parasite) PlayEffect(p1, p2 internal.Player) {
	p1.AddCombat(6) /* or card of 6 or less for free */
}
func (e parasite) AllyEffect(p1, p2 internal.Player) { p1.DrawCards(1) }

type ravager struct{}

func (e ravager) GetName() string     { return "Ravager" }
func (e ravager) GetCost() int        { return 3 }
func (e ravager) GetFaction() Faction { return BLOB }
func (e ravager) PlayEffect(p1, p2 internal.Player) {
	p1.AddCombat(6)
	p1.ScrapFromTradeRow(2)
}

type predator struct{}

func (e predator) GetName() string                   { return "Predator" }
func (e predator) GetCost() int                      { return 2 }
func (e predator) GetFaction() Faction               { return BLOB }
func (e predator) PlayEffect(p1, p2 internal.Player) { p1.AddCombat(4) }
func (e predator) AllyEffect(p1, p2 internal.Player) { p1.DrawCards(1) }

type swarmer struct{}

func (e swarmer) GetName() string     { return "Swarmer" }
func (e swarmer) GetCost() int        { return 1 }
func (e swarmer) GetFaction() Faction { return BLOB }
func (e swarmer) PlayEffect(p1, p2 internal.Player) {
	p1.AddCombat(3)
	p1.ScrapFromTradeRow(1)
}
func (e swarmer) AllyEffect(p1, p2 internal.Player) { p1.AddCombat(2) }
