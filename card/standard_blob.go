package card

func bioformer() Card {
	return &BaseImpl{
		Ship: Ship{
			Name:         "Bioformer",
			Cost:         4,
			Factions:     []Faction{BLOB},
			PlayEffects:  []Effect{AddCombat{3}},
			ScrapEffects: []Effect{AddTrade{3}},
		},
		Defense:       4,
		IsBaseOutpost: false,
	}
}

func plasmaVent() Card {
	return &BaseImpl{
		Ship: Ship{
			Name:         "Plasma Vent",
			Cost:         6,
			Factions:     []Faction{BLOB},
			PlayEffects:  []Effect{AddCombat{4} /*  Put on top of deck when buying */},
			ScrapEffects: []Effect{DestroyBase{}},
		},
		Defense:       5,
		IsBaseOutpost: false,
	}
}

func stellarReef() Card {
	return &BaseImpl{
		Ship: Ship{
			Name:         "Stellar Reef",
			Cost:         2,
			Factions:     []Faction{BLOB},
			PlayEffects:  []Effect{AddTrade{1}},
			ScrapEffects: []Effect{AddCombat{3}},
		},
		Defense:       3,
		IsBaseOutpost: false,
	}
}

func cargoPod() Card {
	return &Ship{
		Name:         "Cargo Pod",
		Cost:         3,
		Factions:     []Faction{BLOB},
		PlayEffects:  []Effect{AddTrade{3}},
		AllyEffects:  []Effect{AddCombat{3}},
		ScrapEffects: []Effect{AddCombat{1}},
	}
}

func leviathan() Card {
	return &Ship{
		Name:        "Leviathan",
		Cost:        8,
		Factions:    []Faction{BLOB},
		PlayEffects: []Effect{AddCombat{9}, DrawCards{1}, DestroyBase{}},
		AllyEffects: []Effect{AcquireCardLessThan{3}},
	}
}

func moonwurm() Card {
	return &Ship{
		Name:        "Moonwurm",
		Cost:        7,
		Factions:    []Faction{BLOB},
		PlayEffects: []Effect{AddCombat{8}, DrawCards{1}},
		AllyEffects: []Effect{AcquireCardLessThan{2}},
	}
}

func parasite() Card {
	return &Ship{
		Name:     "Parasite",
		Cost:     5,
		Factions: []Faction{BLOB},
		PlayEffects: []Effect{
			ChooseBetween{[]Effect{
				AddCombat{6},
				AcquireCardLessThan{6}},
			}},
		AllyEffects: []Effect{DrawCards{1}},
	}
}

func ravager() Card {
	return &Ship{
		Name:     "Ravager",
		Cost:     3,
		Factions: []Faction{BLOB},
		PlayEffects: []Effect{
			AddCombat{6},
			ScrapCards{2, ScrapLocation_TRADE_ROW, false},
		},
	}
}

func predator() Card {
	return &Ship{
		Name:        "Predator",
		Cost:        2,
		Factions:    []Faction{BLOB},
		PlayEffects: []Effect{AddCombat{4}},
		AllyEffects: []Effect{DrawCards{1}},
	}
}

func swarmer() Card {
	return &Ship{
		Name:     "Swarmer",
		Cost:     1,
		Factions: []Faction{BLOB},
		PlayEffects: []Effect{
			AddCombat{3},
			ScrapCards{1, ScrapLocation_TRADE_ROW, false},
		},
		AllyEffects: []Effect{AddCombat{2}},
	}
}
