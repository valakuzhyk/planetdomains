package card

func battleBot() Card {
	return &Ship{
		Name:        "Battle Bot",
		Cost:        1,
		Factions:    []Faction{MACHINE_CULT},
		PlayEffects: []Effect{AddCombat{2}, ScrapCards{1, ScrapLocation_HAND, false}},
		AllyEffects: []Effect{AddCombat{2}},
	}
}

func convoyBot() Card {
	return &Ship{
		Name:        "Convoy Bot",
		Cost:        3,
		Factions:    []Faction{MACHINE_CULT},
		PlayEffects: []Effect{AddCombat{4}, ScrapCards{1, ScrapLocation_HAND_OR_DISCARD, false}},
		AllyEffects: []Effect{AddCombat{2}},
	}
}

func mechCruiser() Card {
	return &Ship{
		Name:        "Mech Cruiser",
		Cost:        5,
		Factions:    []Faction{MACHINE_CULT},
		PlayEffects: []Effect{AddCombat{6}, ScrapCards{1, ScrapLocation_HAND_OR_DISCARD, false}},
		AllyEffects: []Effect{DestroyBase{}},
	}
}

func miningMech() Card {
	return &Ship{
		Name:        "Mining Mech",
		Cost:        4,
		Factions:    []Faction{MACHINE_CULT},
		PlayEffects: []Effect{AddTrade{3}, ScrapCards{1, ScrapLocation_HAND_OR_DISCARD, false}},
		AllyEffects: []Effect{AddCombat{3}},
	}
}

func repairBot() Card {
	return &Ship{
		Name:         "Repair Bot",
		Cost:         2,
		Factions:     []Faction{MACHINE_CULT},
		PlayEffects:  []Effect{AddTrade{2}, ScrapCards{1, ScrapLocation_DISCARD, false}},
		ScrapEffects: []Effect{AddCombat{2}},
	}
}

func theWrecker() Card {
	return &Ship{
		Name:        "The Wrecker",
		Cost:        7,
		Factions:    []Faction{MACHINE_CULT},
		PlayEffects: []Effect{AddCombat{6}, ScrapCards{2, ScrapLocation_HAND_OR_DISCARD, false}},
		AllyEffects: []Effect{DrawCards{1}},
	}
}

func frontierStation() Card {
	return &BaseImpl{
		Ship: Ship{
			Name:        "Frontier Station",
			Cost:        6,
			Factions:    []Faction{MACHINE_CULT},
			PlayEffects: []Effect{AddCombat{2} /* or 3 combat */},
		},
		Defense:       6,
		IsBaseOutpost: true,
	}
}

func stealthTower() Card {
	return &BaseImpl{
		Ship: Ship{
			Name:        "Stealth Tower",
			Cost:        5,
			Factions:    []Faction{MACHINE_CULT},
			PlayEffects: []Effect{ /* Copy Base */ },
		},
		Defense:       5,
		IsBaseOutpost: true,
	}
}

func theIncinerator() Card {
	return &BaseImpl{
		Ship: Ship{
			Name:        "The Incinerator",
			Cost:        8,
			Factions:    []Faction{MACHINE_CULT},
			PlayEffects: []Effect{ScrapCards{2, ScrapLocation_HAND_OR_DISCARD, false}},
			AllyEffects: []Effect{ /* effect for every card scrapped from hand or discard */ },
		},
		Defense:       6,
		IsBaseOutpost: true,
	}
}

func theOracle() Card {
	return &BaseImpl{
		Ship: Ship{
			Name:        "The Oracle",
			Cost:        4,
			Factions:    []Faction{MACHINE_CULT},
			PlayEffects: []Effect{ScrapCards{1, ScrapLocation_HAND, false}},
			AllyEffects: []Effect{AddCombat{3}},
		},
		Defense:       5,
		IsBaseOutpost: true,
	}
}

func warningBeacon() Card {
	return &BaseImpl{
		Ship: Ship{
			Name:        "Warning Beacon",
			Cost:        2,
			Factions:    []Faction{MACHINE_CULT},
			PlayEffects: []Effect{AddCombat{2} /* Purchase into hand */},
		},
		Defense:       2,
		IsBaseOutpost: true,
	}
}
