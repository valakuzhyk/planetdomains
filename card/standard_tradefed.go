package card

func colonySeedShip() Card {
	return &Ship{
		Name:        "Colony Seed Ship",
		Cost:        5,
		Factions:    []Faction{TRADE_FED},
		PlayEffects: []Effect{AddTrade{3}, AddCombat{3}, AddAuthority{3}},
		/* if you aquire, place in hand */
	}
}

func frontierFerry() Card {
	return &Ship{
		Name:         "Frontier Ferry",
		Cost:         4,
		Factions:     []Faction{TRADE_FED},
		PlayEffects:  []Effect{AddTrade{3}, AddAuthority{4}},
		ScrapEffects: []Effect{DestroyBase{}},
	}
}

func patrolCutter() Card {
	return &Ship{
		Name:        "Patrol Cutter",
		Cost:        3,
		Factions:    []Faction{TRADE_FED},
		PlayEffects: []Effect{AddTrade{2}, AddCombat{3}},
		AllyEffects: []Effect{AddAuthority{4}},
	}
}

func peacekeeper() Card {
	return &Ship{
		Name:        "Peacekeeper",
		Cost:        6,
		Factions:    []Faction{TRADE_FED},
		PlayEffects: []Effect{AddCombat{6}, AddAuthority{6}},
		AllyEffects: []Effect{DrawCards{1}},
	}
}

func solarSkiff() Card {
	return &Ship{
		Name:        "Solar Skiff",
		Cost:        1,
		Factions:    []Faction{TRADE_FED},
		PlayEffects: []Effect{AddTrade{2}},
		AllyEffects: []Effect{DrawCards{1}},
	}
}

func tradeHauler() Card {
	return &Ship{
		Name:        "Trade Hauler",
		Cost:        2,
		Factions:    []Faction{TRADE_FED},
		PlayEffects: []Effect{AddTrade{3}},
		AllyEffects: []Effect{AddAuthority{3}},
	}
}

func centralStation() Card {
	return &BaseImpl{
		Ship: Ship{
			Name:        "Central Station",
			Cost:        4,
			Factions:    []Faction{TRADE_FED},
			PlayEffects: []Effect{AddTrade{2} /* conditional gain 4 auth and draw */},
		},
		Defense:       5,
		IsBaseOutpost: false,
	}
}

func factoryWorld() Card {
	return &BaseImpl{
		Ship: Ship{
			Name:        "Factory World",
			Cost:        8,
			Factions:    []Faction{TRADE_FED},
			PlayEffects: []Effect{AddTrade{3} /* put next ship purchased in hand */},
		},
		Defense:       6,
		IsBaseOutpost: true,
	}
}

func federationShipyard() Card {
	return &BaseImpl{
		Ship: Ship{
			Name:        "Federation Shipyeard",
			Cost:        6,
			Factions:    []Faction{TRADE_FED},
			PlayEffects: []Effect{AddTrade{2}},
			AllyEffects: []Effect{ /* next card on top of deck */ },
		},
		Defense:       6,
		IsBaseOutpost: true,
	}
}

func loyalColony() Card {
	return &BaseImpl{
		Ship: Ship{
			Name:        "Loyal Colony",
			Cost:        7,
			Factions:    []Faction{TRADE_FED},
			PlayEffects: []Effect{AddTrade{3}, AddCombat{3}, AddAuthority{3}},
		},
		Defense:       6,
		IsBaseOutpost: false,
	}
}

func storageSilo() Card {
	return &BaseImpl{
		Ship: Ship{
			Name:        "Storage Silo",
			Cost:        2,
			Factions:    []Faction{TRADE_FED},
			PlayEffects: []Effect{AddAuthority{2}},
			AllyEffects: []Effect{AddTrade{2}},
		},
		Defense:       3,
		IsBaseOutpost: false,
	}
}
