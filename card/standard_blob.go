package card

import "github.com/valakuzhyk/planetdomains/internal"

type bioformer struct{}

func (e bioformer) GetName() string               { return "Bioformer" }
func (e bioformer) GetCost() int                  { return 4 }
func (e bioformer) GetFaction() Faction           { return BLOB }
func (e bioformer) PlayEffect(p internal.Player)  { p.AddCombat(3) }
func (e bioformer) ScrapEffect(p internal.Player) { p.AddTrade(3) }
func (e bioformer) GetDefense() (int, bool)       { return 4, false }

type plasmaVent struct{}

func (e plasmaVent) GetName() string     { return "Plasma Vent" }
func (e plasmaVent) GetCost() int        { return 6 }
func (e plasmaVent) GetFaction() Faction { return BLOB }
func (e plasmaVent) PlayEffect(p internal.Player) {
	p.AddCombat(4) /* if you aquired this after playing a blob, direct to hand */
}
func (e plasmaVent) ScrapEffect(p internal.Player) { /* Destroy base */ }
func (e plasmaVent) GetDefense() (int, bool)       { return 5, false }

type stellarReef struct{}

func (e stellarReef) GetName() string               { return "Stellar Reef" }
func (e stellarReef) GetCost() int                  { return 2 }
func (e stellarReef) GetFaction() Faction           { return BLOB }
func (e stellarReef) PlayEffect(p internal.Player)  { p.AddTrade(1) }
func (e stellarReef) ScrapEffect(p internal.Player) { p.AddCombat(3) }
func (e stellarReef) GetDefense() (int, bool)       { return 3, false }

type cargoPod struct{}

func (e cargoPod) GetName() string               { return "Cargo Pod" }
func (e cargoPod) GetCost() int                  { return 3 }
func (e cargoPod) GetFaction() Faction           { return BLOB }
func (e cargoPod) PlayEffect(p internal.Player)  { p.AddTrade(3) }
func (e cargoPod) AllyEffect(p internal.Player)  { p.AddCombat(3) }
func (e cargoPod) ScrapEffect(p internal.Player) { p.AddCombat(3) }

type leviathan struct{}

func (e leviathan) GetName() string              { return "Leviathan" }
func (e leviathan) GetCost() int                 { return 8 }
func (e leviathan) GetFaction() Faction          { return BLOB }
func (e leviathan) PlayEffect(p internal.Player) { p.AddCombat(9); p.DrawCards(1) /* destroy base */ }
func (e leviathan) AllyEffect(p internal.Player) { /* acquire a cost of three or less, put into your hand*/
}

type moonwurm struct{}

func (e moonwurm) GetName() string              { return "Moonwurm" }
func (e moonwurm) GetCost() int                 { return 7 }
func (e moonwurm) GetFaction() Faction          { return BLOB }
func (e moonwurm) PlayEffect(p internal.Player) { p.AddCombat(8); p.DrawCards(1) }
func (e moonwurm) AllyEffect(p internal.Player) { /* acquire a cost of two or less, put into your hand*/
}

type parasite struct{}

func (e parasite) GetName() string              { return "Parasite" }
func (e parasite) GetCost() int                 { return 5 }
func (e parasite) GetFaction() Faction          { return BLOB }
func (e parasite) PlayEffect(p internal.Player) { p.AddCombat(6) /* or card of 6 or less for free */ }
func (e parasite) AllyEffect(p internal.Player) { p.DrawCards(1) }

type ravager struct{}

func (e ravager) GetName() string     { return "Ravager" }
func (e ravager) GetCost() int        { return 3 }
func (e ravager) GetFaction() Faction { return BLOB }
func (e ravager) PlayEffect(p internal.Player) {
	p.AddCombat(6) /* scrap up to two cards in the trade row */
}

type predator struct{}

func (e predator) GetName() string              { return "Predator" }
func (e predator) GetCost() int                 { return 2 }
func (e predator) GetFaction() Faction          { return BLOB }
func (e predator) PlayEffect(p internal.Player) { p.AddCombat(4) }
func (e predator) AllyEffect(p internal.Player) { p.DrawCards(1) }

type swarmer struct{}

func (e swarmer) GetName() string     { return "Swarmer" }
func (e swarmer) GetCost() int        { return 1 }
func (e swarmer) GetFaction() Faction { return BLOB }
func (e swarmer) PlayEffect(p internal.Player) {
	p.AddCombat(3) /* you can scrap a trade in the trade row */
}
func (e swarmer) AllyEffect(p internal.Player) { p.AddCombat(2) }
