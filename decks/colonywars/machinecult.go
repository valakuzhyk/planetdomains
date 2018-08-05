package colonywars

import (
	"github.com/valakuzhyk/planetdomains/card"
	"github.com/valakuzhyk/planetdomains/cardimpl"
	"github.com/valakuzhyk/planetdomains/effect"
)

func battleBot() card.Card {
	return &cardimpl.Ship{
		Name:        "Battle Bot",
		Cost:        1,
		Factions:    []card.Faction{card.MACHINE_CULT},
		PlayEffects: []card.Effect{effect.AddCombat{2}, effect.ScrapCards{1, effect.ScrapLocation_HAND, false}},
		AllyEffects: []card.Effect{effect.AddCombat{2}},
	}
}

func convoyBot() card.Card {
	return &cardimpl.Ship{
		Name:        "Convoy Bot",
		Cost:        3,
		Factions:    []card.Faction{card.MACHINE_CULT},
		PlayEffects: []card.Effect{effect.AddCombat{4}, effect.ScrapCards{1, effect.ScrapLocation_HAND_OR_DISCARD, false}},
		AllyEffects: []card.Effect{effect.AddCombat{2}},
	}
}

func mechCruiser() card.Card {
	return &cardimpl.Ship{
		Name:        "Mech Cruiser",
		Cost:        5,
		Factions:    []card.Faction{card.MACHINE_CULT},
		PlayEffects: []card.Effect{effect.AddCombat{6}, effect.ScrapCards{1, effect.ScrapLocation_HAND_OR_DISCARD, false}},
		AllyEffects: []card.Effect{effect.DestroyBase{}},
	}
}

func miningMech() card.Card {
	return &cardimpl.Ship{
		Name:        "Mining Mech",
		Cost:        4,
		Factions:    []card.Faction{card.MACHINE_CULT},
		PlayEffects: []card.Effect{effect.AddTrade{3}, effect.ScrapCards{1, effect.ScrapLocation_HAND_OR_DISCARD, false}},
		AllyEffects: []card.Effect{effect.AddCombat{3}},
	}
}

func repairBot() card.Card {
	return &cardimpl.Ship{
		Name:         "Repair Bot",
		Cost:         2,
		Factions:     []card.Faction{card.MACHINE_CULT},
		PlayEffects:  []card.Effect{effect.AddTrade{2}, effect.ScrapCards{1, effect.ScrapLocation_DISCARD, false}},
		ScrapEffects: []card.Effect{effect.AddCombat{2}},
	}
}

func theWrecker() card.Card {
	return &cardimpl.Ship{
		Name:        "The Wrecker",
		Cost:        7,
		Factions:    []card.Faction{card.MACHINE_CULT},
		PlayEffects: []card.Effect{effect.AddCombat{6}, effect.ScrapCards{2, effect.ScrapLocation_HAND_OR_DISCARD, false}},
		AllyEffects: []card.Effect{effect.DrawCards{1}},
	}
}

func frontierStation() card.Card {
	return &cardimpl.BaseImpl{
		Ship: cardimpl.Ship{
			Name:     "Frontier Station",
			Cost:     6,
			Factions: []card.Faction{card.MACHINE_CULT},
			PlayEffects: []card.Effect{effect.ChooseBetween{[]card.Effect{
				effect.AddTrade{2},
				effect.AddCombat{3},
			}}},
		},
		Defense:       6,
		IsBaseOutpost: true,
	}
}

func stealthTower() card.Card {
	return &cardimpl.BaseImpl{
		Ship: cardimpl.Ship{
			Name:        "Stealth Tower",
			Cost:        5,
			Factions:    []card.Faction{card.MACHINE_CULT},
			PlayEffects: []card.Effect{ /* Copy Base */ },
		},
		Defense:       5,
		IsBaseOutpost: true,
	}
}

func theIncinerator() card.Card {
	return &cardimpl.BaseImpl{
		Ship: cardimpl.Ship{
			Name:        "The Incinerator",
			Cost:        8,
			Factions:    []card.Faction{card.MACHINE_CULT},
			PlayEffects: []card.Effect{effect.ScrapCards{2, effect.ScrapLocation_HAND_OR_DISCARD, false}},
			AllyEffects: []card.Effect{ /* effect for every card.Card scrapped from hand or discard.Card */ },
		},
		Defense:       6,
		IsBaseOutpost: true,
	}
}

func theOracle() card.Card {
	return &cardimpl.BaseImpl{
		Ship: cardimpl.Ship{
			Name:        "The Oracle",
			Cost:        4,
			Factions:    []card.Faction{card.MACHINE_CULT},
			PlayEffects: []card.Effect{effect.ScrapCards{1, effect.ScrapLocation_HAND, false}},
			AllyEffects: []card.Effect{effect.AddCombat{3}},
		},
		Defense:       5,
		IsBaseOutpost: true,
	}
}

func warningBeacon() card.Card {
	return &cardimpl.BaseImpl{
		Ship: cardimpl.Ship{
			Name:        "Warning Beacon",
			Cost:        2,
			Factions:    []card.Faction{card.MACHINE_CULT},
			PlayEffects: []card.Effect{effect.AddCombat{2} /* Purchase into hand */},
		},
		Defense:       2,
		IsBaseOutpost: true,
	}
}
