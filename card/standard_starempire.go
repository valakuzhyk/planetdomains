package card

func agingBattleship() Card {
	return &Ship{
		Name:         "Aging Battleship",
		Cost:         5,
		Factions:     []Faction{STAR_EMPIRE},
		PlayEffects:  []Effect{AddCombat{5}},
		AllyEffects:  []Effect{DrawCards{1}},
		ScrapEffects: []Effect{AddCombat{2}, DrawCards{2}},
	}
}

func emperorsDreadnaught() Card {
	return &Ship{
		Name:        "Emperor's Dreadnaught",
		Cost:        8,
		Factions:    []Faction{STAR_EMPIRE},
		PlayEffects: []Effect{AddCombat{8}, DrawCards{1}, OpponentDiscards{1} /* buy in hand */},
	}
}

func falcon() Card {
	return &Ship{
		Name:         "Falcon",
		Cost:         3,
		Factions:     []Faction{STAR_EMPIRE},
		PlayEffects:  []Effect{AddCombat{2}, DrawCards{1}},
		ScrapEffects: []Effect{OpponentDiscards{1}},
	}
}

func gunship() Card {
	return &Ship{
		Name:         "Gunship",
		Cost:         4,
		Factions:     []Faction{STAR_EMPIRE},
		PlayEffects:  []Effect{AddCombat{5}, OpponentDiscards{1}},
		ScrapEffects: []Effect{AddTrade{4}},
	}
}

func heavyCruiser() Card {
	return &Ship{
		Name:        "Heavy Cruiser",
		Cost:        5,
		Factions:    []Faction{STAR_EMPIRE},
		PlayEffects: []Effect{AddCombat{4}, DrawCards{1}},
		AllyEffects: []Effect{DrawCards{1}},
	}
}

func lancer() Card {
	return &Ship{
		Name:        "Lancer",
		Cost:        2,
		Factions:    []Faction{STAR_EMPIRE},
		PlayEffects: []Effect{AddCombat{4} /* conditional add 2 combat*/},
		AllyEffects: []Effect{OpponentDiscards{1}},
	}
}

func starBarge() Card {
	return &Ship{
		Name:        "Star Barge",
		Cost:        1,
		Factions:    []Faction{STAR_EMPIRE},
		PlayEffects: []Effect{AddTrade{2}},
		AllyEffects: []Effect{AddCombat{2}, OpponentDiscards{1}},
	}
}

func commandCenter() Card {
	return &BaseImpl{
		Ship: Ship{
			Name:        "Command Center",
			Cost:        4,
			Factions:    []Faction{STAR_EMPIRE},
			PlayEffects: []Effect{AddTrade{2} /* whenever star is played, add combat */},
		},
		Defense:       4,
		IsBaseOutpost: true,
	}
}

func imperialPalace() Card {
	return &BaseImpl{
		Ship: Ship{
			Name:        "Imperial Palace",
			Cost:        7,
			Factions:    []Faction{STAR_EMPIRE},
			PlayEffects: []Effect{DrawCards{1}, OpponentDiscards{1}},
			AllyEffects: []Effect{AddCombat{4}},
		},
		Defense:       6,
		IsBaseOutpost: true,
	}
}

func orbitalPlatform() Card {
	return &BaseImpl{
		Ship: Ship{
			Name:        "Orbital Platform",
			Cost:        3,
			Factions:    []Faction{STAR_EMPIRE},
			PlayEffects: []Effect{ /* If you discard, draw */ },
			AllyEffects: []Effect{AddCombat{3}},
		},
		Defense:       4,
		IsBaseOutpost: false,
	}
}

func supplyDepot() Card {
	return &BaseImpl{
		Ship: Ship{
			Name:        "Supply Depot",
			Cost:        6,
			Factions:    []Faction{STAR_EMPIRE},
			PlayEffects: []Effect{ /* Discard up to two, gain either 2 gold or 2 combat for each card */ },
			AllyEffects: []Effect{DrawCards{1}},
		},
		Defense:       5,
		IsBaseOutpost: true,
	}
}
