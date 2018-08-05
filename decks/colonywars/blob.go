package colonywars

import (
	"github.com/valakuzhyk/planetdomains/card"
	"github.com/valakuzhyk/planetdomains/cardimpl"
	"github.com/valakuzhyk/planetdomains/effect"
)

func bioformer() card.Card {
	return &cardimpl.BaseImpl{
		Ship: cardimpl.Ship{
			Name:         "Bioformer",
			Cost:         4,
			Factions:     []card.Faction{card.BLOB},
			PlayEffects:  []card.Effect{effect.AddCombat{3}},
			ScrapEffects: []card.Effect{effect.AddTrade{3}},
		},
		Defense:       4,
		IsBaseOutpost: false,
	}
}

func plasmaVent() card.Card {
	return &cardimpl.BaseImpl{
		Ship: cardimpl.Ship{
			Name:         "Plasma Vent",
			Cost:         6,
			Factions:     []card.Faction{card.BLOB},
			PlayEffects:  []card.Effect{effect.AddCombat{4} /*  Put on top of deck when buying */},
			ScrapEffects: []card.Effect{effect.DestroyBase{}},
		},
		Defense:       5,
		IsBaseOutpost: false,
	}
}

func stellarReef() card.Card {
	return &cardimpl.BaseImpl{
		Ship: cardimpl.Ship{
			Name:         "Stellar Reef",
			Cost:         2,
			Factions:     []card.Faction{card.BLOB},
			PlayEffects:  []card.Effect{effect.AddTrade{1}},
			ScrapEffects: []card.Effect{effect.AddCombat{3}},
		},
		Defense:       3,
		IsBaseOutpost: false,
	}
}

func cargoPod() card.Card {
	return &cardimpl.Ship{
		Name:         "Cargo Pod",
		Cost:         3,
		Factions:     []card.Faction{card.BLOB},
		PlayEffects:  []card.Effect{effect.AddTrade{3}},
		AllyEffects:  []card.Effect{effect.AddCombat{3}},
		ScrapEffects: []card.Effect{effect.AddCombat{1}},
	}
}

func leviathan() card.Card {
	return &cardimpl.Ship{
		Name:     "Leviathan",
		Cost:     8,
		Factions: []card.Faction{card.BLOB},
		PlayEffects: []card.Effect{
			effect.AddCombat{9},
			effect.DrawCards{1},
			effect.DestroyBase{}},
		AllyEffects: []card.Effect{effect.AcquireCardLessThan{3}},
	}
}

func moonwurm() card.Card {
	return &cardimpl.Ship{
		Name:        "Moonwurm",
		Cost:        7,
		Factions:    []card.Faction{card.BLOB},
		PlayEffects: []card.Effect{effect.AddCombat{8}, effect.DrawCards{1}},
		AllyEffects: []card.Effect{effect.AcquireCardLessThan{2}},
	}
}

func parasite() card.Card {
	return &cardimpl.Ship{
		Name:     "Parasite",
		Cost:     5,
		Factions: []card.Faction{card.BLOB},
		PlayEffects: []card.Effect{
			effect.ChooseBetween{[]card.Effect{
				effect.AddCombat{6},
				effect.AcquireCardLessThan{6}},
			}},
		AllyEffects: []card.Effect{effect.DrawCards{1}},
	}
}

func ravager() card.Card {
	return &cardimpl.Ship{
		Name:     "Ravager",
		Cost:     3,
		Factions: []card.Faction{card.BLOB},
		PlayEffects: []card.Effect{
			effect.AddCombat{6},
			effect.ScrapCards{2, effect.ScrapLocation_TRADE_ROW, false},
		},
	}
}

func predator() card.Card {
	return &cardimpl.Ship{
		Name:        "Predator",
		Cost:        2,
		Factions:    []card.Faction{card.BLOB},
		PlayEffects: []card.Effect{effect.AddCombat{4}},
		AllyEffects: []card.Effect{effect.DrawCards{1}},
	}
}

func swarmer() card.Card {
	return &cardimpl.Ship{
		Name:     "Swarmer",
		Cost:     1,
		Factions: []card.Faction{card.BLOB},
		PlayEffects: []card.Effect{
			effect.AddCombat{3},
			effect.ScrapCards{1, effect.ScrapLocation_TRADE_ROW, false},
		},
		AllyEffects: []card.Effect{effect.AddCombat{2}},
	}
}
