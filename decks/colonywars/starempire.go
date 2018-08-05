package colonywars

import (
	"github.com/valakuzhyk/planetdomains/card"
	"github.com/valakuzhyk/planetdomains/cardimpl"
	"github.com/valakuzhyk/planetdomains/effect"
)

func agingBattleship() card.Card {
	return &cardimpl.Ship{
		Name:         "Aging Battleship",
		Cost:         5,
		Factions:     []card.Faction{card.STAR_EMPIRE},
		PlayEffects:  []card.Effect{effect.AddCombat{5}},
		AllyEffects:  []card.Effect{effect.DrawCards{1}},
		ScrapEffects: []card.Effect{effect.AddCombat{2}, effect.DrawCards{2}},
	}
}

func emperorsDreadnaught() card.Card {
	return &cardimpl.Ship{
		Name:        "Emperor's Dreadnaught",
		Cost:        8,
		Factions:    []card.Faction{card.STAR_EMPIRE},
		PlayEffects: []card.Effect{effect.AddCombat{8}, effect.DrawCards{1}, effect.OpponentDiscards{1} /* buy in hand */},
	}
}

func falcon() card.Card {
	return &cardimpl.Ship{
		Name:         "Falcon",
		Cost:         3,
		Factions:     []card.Faction{card.STAR_EMPIRE},
		PlayEffects:  []card.Effect{effect.AddCombat{2}, effect.DrawCards{1}},
		ScrapEffects: []card.Effect{effect.OpponentDiscards{1}},
	}
}

func gunship() card.Card {
	return &cardimpl.Ship{
		Name:         "Gunship",
		Cost:         4,
		Factions:     []card.Faction{card.STAR_EMPIRE},
		PlayEffects:  []card.Effect{effect.AddCombat{5}, effect.OpponentDiscards{1}},
		ScrapEffects: []card.Effect{effect.AddTrade{4}},
	}
}

func heavyCruiser() card.Card {
	return &cardimpl.Ship{
		Name:        "Heavy Cruiser",
		Cost:        5,
		Factions:    []card.Faction{card.STAR_EMPIRE},
		PlayEffects: []card.Effect{effect.AddCombat{4}, effect.DrawCards{1}},
		AllyEffects: []card.Effect{effect.DrawCards{1}},
	}
}

func lancer() card.Card {
	return &cardimpl.Ship{
		Name:        "Lancer",
		Cost:        2,
		Factions:    []card.Faction{card.STAR_EMPIRE},
		PlayEffects: []card.Effect{effect.AddCombat{4} /* conditional add 2 combat*/},
		AllyEffects: []card.Effect{effect.OpponentDiscards{1}},
	}
}

func starBarge() card.Card {
	return &cardimpl.Ship{
		Name:        "Star Barge",
		Cost:        1,
		Factions:    []card.Faction{card.STAR_EMPIRE},
		PlayEffects: []card.Effect{effect.AddTrade{2}},
		AllyEffects: []card.Effect{effect.AddCombat{2}, effect.OpponentDiscards{1}},
	}
}

func commandCenter() card.Card {
	return &cardimpl.BaseImpl{
		Ship: cardimpl.Ship{
			Name:        "Command Center",
			Cost:        4,
			Factions:    []card.Faction{card.STAR_EMPIRE},
			PlayEffects: []card.Effect{effect.AddTrade{2} /* whenever star is played, add combat */},
		},
		Defense:       4,
		IsBaseOutpost: true,
	}
}

func imperialPalace() card.Card {
	return &cardimpl.BaseImpl{
		Ship: cardimpl.Ship{
			Name:        "Imperial Palace",
			Cost:        7,
			Factions:    []card.Faction{card.STAR_EMPIRE},
			PlayEffects: []card.Effect{effect.DrawCards{1}, effect.OpponentDiscards{1}},
			AllyEffects: []card.Effect{effect.AddCombat{4}},
		},
		Defense:       6,
		IsBaseOutpost: true,
	}
}

func orbitalPlatform() card.Card {
	return &cardimpl.BaseImpl{
		Ship: cardimpl.Ship{
			Name:        "Orbital Platform",
			Cost:        3,
			Factions:    []card.Faction{card.STAR_EMPIRE},
			PlayEffects: []card.Effect{ /* If you discard, draw */ },
			AllyEffects: []card.Effect{effect.AddCombat{3}},
		},
		Defense:       4,
		IsBaseOutpost: false,
	}
}

func supplyDepot() card.Card {
	return &cardimpl.BaseImpl{
		Ship: cardimpl.Ship{
			Name:        "Supply Depot",
			Cost:        6,
			Factions:    []card.Faction{card.STAR_EMPIRE},
			PlayEffects: []card.Effect{ /* Discard.Card up to two, gain either 2 gold or 2 combat for each card.Card */ },
			AllyEffects: []card.Effect{effect.DrawCards{1}},
		},
		Defense:       5,
		IsBaseOutpost: true,
	}
}
