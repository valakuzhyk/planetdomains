package colonywars

import (
	"github.com/valakuzhyk/planetdomains/card"
	"github.com/valakuzhyk/planetdomains/cardimpl"
	"github.com/valakuzhyk/planetdomains/effect"
)

func colonySeedShip() card.Card {
	return &cardimpl.Ship{
		Name:        "Colony Seed Ship",
		Cost:        5,
		Factions:    []card.Faction{card.TRADE_FED},
		PlayEffects: []card.Effect{effect.AddTrade{3}, effect.AddCombat{3}, effect.AddAuthority{3}},
		/* if you aquire, place in hand */
	}
}

func frontierFerry() card.Card {
	return &cardimpl.Ship{
		Name:         "Frontier Ferry",
		Cost:         4,
		Factions:     []card.Faction{card.TRADE_FED},
		PlayEffects:  []card.Effect{effect.AddTrade{3}, effect.AddAuthority{4}},
		ScrapEffects: []card.Effect{effect.DestroyBase{}},
	}
}

func patrolCutter() card.Card {
	return &cardimpl.Ship{
		Name:        "Patrol Cutter",
		Cost:        3,
		Factions:    []card.Faction{card.TRADE_FED},
		PlayEffects: []card.Effect{effect.AddTrade{2}, effect.AddCombat{3}},
		AllyEffects: []card.Effect{effect.AddAuthority{4}},
	}
}

func peacekeeper() card.Card {
	return &cardimpl.Ship{
		Name:        "Peacekeeper",
		Cost:        6,
		Factions:    []card.Faction{card.TRADE_FED},
		PlayEffects: []card.Effect{effect.AddCombat{6}, effect.AddAuthority{6}},
		AllyEffects: []card.Effect{effect.DrawCards{1}},
	}
}

func solarSkiff() card.Card {
	return &cardimpl.Ship{
		Name:        "Solar Skiff",
		Cost:        1,
		Factions:    []card.Faction{card.TRADE_FED},
		PlayEffects: []card.Effect{effect.AddTrade{2}},
		AllyEffects: []card.Effect{effect.DrawCards{1}},
	}
}

func tradeHauler() card.Card {
	return &cardimpl.Ship{
		Name:        "Trade Hauler",
		Cost:        2,
		Factions:    []card.Faction{card.TRADE_FED},
		PlayEffects: []card.Effect{effect.AddTrade{3}},
		AllyEffects: []card.Effect{effect.AddAuthority{3}},
	}
}

func centralStation() card.Card {
	return &cardimpl.BaseImpl{
		Ship: cardimpl.Ship{
			Name:        "Central Station",
			Cost:        4,
			Factions:    []card.Faction{card.TRADE_FED},
			PlayEffects: []card.Effect{effect.AddTrade{2} /* conditional gain 4 auth and draw */},
		},
		Defense:       5,
		IsBaseOutpost: false,
	}
}

func factoryWorld() card.Card {
	return &cardimpl.BaseImpl{
		Ship: cardimpl.Ship{
			Name:        "Factory World",
			Cost:        8,
			Factions:    []card.Faction{card.TRADE_FED},
			PlayEffects: []card.Effect{effect.AddTrade{3} /* put next ship purchased in hand */},
		},
		Defense:       6,
		IsBaseOutpost: true,
	}
}

func federationShipyard() card.Card {
	return &cardimpl.BaseImpl{
		Ship: cardimpl.Ship{
			Name:        "Federation Shipyeard",
			Cost:        6,
			Factions:    []card.Faction{card.TRADE_FED},
			PlayEffects: []card.Effect{effect.AddTrade{2}},
			AllyEffects: []card.Effect{ /* next card.Card on top of deck */ },
		},
		Defense:       6,
		IsBaseOutpost: true,
	}
}

func loyalColony() card.Card {
	return &cardimpl.BaseImpl{
		Ship: cardimpl.Ship{
			Name:        "Loyal Colony",
			Cost:        7,
			Factions:    []card.Faction{card.TRADE_FED},
			PlayEffects: []card.Effect{effect.AddTrade{3}, effect.AddCombat{3}, effect.AddAuthority{3}},
		},
		Defense:       6,
		IsBaseOutpost: false,
	}
}

func storageSilo() card.Card {
	return &cardimpl.BaseImpl{
		Ship: cardimpl.Ship{
			Name:        "Storage Silo",
			Cost:        2,
			Factions:    []card.Faction{card.TRADE_FED},
			PlayEffects: []card.Effect{effect.AddAuthority{2}},
			AllyEffects: []card.Effect{effect.AddTrade{2}},
		},
		Defense:       3,
		IsBaseOutpost: false,
	}
}
