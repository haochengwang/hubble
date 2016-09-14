package main

func InitBasicCardSchools() (result map[int]*CardSchool) {
	schools := make([]*CardSchool, 0)

	// Farms
	schools = append(schools, &CardSchool{
		schoolName:     "Agriculture",
		shortName:      "Agri.",
		age:            0,
		cardType:       []CardType{CARDTYPE_TECH, CARDTYPE_TECH_FARM},
		buildCost:      2,
		productionCrop: 1,
		cardCounts:     []int{2, 3, 4},
	})
	schools = append(schools, &CardSchool{
		schoolName:     "Irrigation",
		shortName:      "Irri.",
		age:            1,
		cardType:       []CardType{CARDTYPE_TECH, CARDTYPE_TECH_FARM},
		tech:           3,
		buildCost:      4,
		productionCrop: 2,
		cardCounts:     []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:     "Selective Breeding",
		shortName:      "S. Brd.",
		age:            2,
		cardType:       []CardType{CARDTYPE_TECH, CARDTYPE_TECH_FARM},
		tech:           5,
		buildCost:      6,
		productionCrop: 3,
		cardCounts:     []int{1, 2, 3},
	})
	schools = append(schools, &CardSchool{
		schoolName:     "Mech. Agriculture",
		shortName:      "M. Agr.",
		age:            3,
		cardType:       []CardType{CARDTYPE_TECH, CARDTYPE_TECH_FARM},
		tech:           7,
		buildCost:      8,
		productionCrop: 5,
		cardCounts:     []int{1, 2, 2},
	})

	// Mines
	schools = append(schools, &CardSchool{
		schoolName:         "Bronze",
		shortName:          "Bronze",
		age:                0,
		cardType:           []CardType{CARDTYPE_TECH, CARDTYPE_TECH_MINE},
		buildCost:          2,
		productionResource: 1,
		cardCounts:         []int{2, 3, 4},
	})
	schools = append(schools, &CardSchool{
		schoolName:         "Iron",
		shortName:          "Iron",
		age:                1,
		cardType:           []CardType{CARDTYPE_TECH, CARDTYPE_TECH_MINE},
		tech:               5,
		buildCost:          5,
		productionResource: 2,
		cardCounts:         []int{2, 2, 3},
	})
	schools = append(schools, &CardSchool{
		schoolName:         "Coal",
		shortName:          "Coal",
		age:                2,
		cardType:           []CardType{CARDTYPE_TECH, CARDTYPE_TECH_MINE},
		tech:               7,
		buildCost:          8,
		productionResource: 3,
		cardCounts:         []int{1, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:         "Oil",
		shortName:          "Oil",
		age:                3,
		cardType:           []CardType{CARDTYPE_TECH, CARDTYPE_TECH_MINE},
		tech:               9,
		buildCost:          11,
		productionResource: 5,
		cardCounts:         []int{1, 2, 2},
	})

	// Urban - Labs
	schools = append(schools, &CardSchool{
		schoolName: "Philosophy",
		shortName:  "Phil.",
		age:        0,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_URBAN,
			CARDTYPE_TECH_URBAN_LAB},
		buildCost:      3,
		productionTech: 1,
		cardCounts:     []int{2, 3, 4},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Alchemy",
		shortName:  "Alchemy",
		age:        1,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_URBAN,
			CARDTYPE_TECH_URBAN_LAB},
		tech:           4,
		buildCost:      6,
		productionTech: 2,
		cardCounts:     []int{2, 2, 3},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Scientific Method",
		shortName:  "S. Mth.",
		age:        2,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_URBAN,
			CARDTYPE_TECH_URBAN_LAB},
		tech:           6,
		buildCost:      8,
		productionTech: 3,
		cardCounts:     []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Computers",
		shortName:  "Comp.",
		age:        3,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_URBAN,
			CARDTYPE_TECH_URBAN_LAB},
		tech:           8,
		buildCost:      11,
		productionTech: 5,
		cardCounts:     []int{2, 2, 2},
	})

	// Urban - temples
	schools = append(schools, &CardSchool{
		schoolName: "Religion",
		shortName:  "Rel.",
		age:        0,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_URBAN,
			CARDTYPE_TECH_URBAN_TEMPLE},
		buildCost:           3,
		productionHappiness: 1,
		productionCulture:   1,
		cardCounts:          []int{2, 3, 4},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Theology",
		shortName:  "Theo.",
		age:        1,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_URBAN,
			CARDTYPE_TECH_URBAN_TEMPLE},
		tech:                2,
		buildCost:           5,
		productionHappiness: 2,
		productionCulture:   1,
		cardCounts:          []int{1, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Org. Religion",
		shortName:  "O. Rel.",
		age:        2,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_URBAN,
			CARDTYPE_TECH_URBAN_TEMPLE},
		tech:                4,
		buildCost:           7,
		productionHappiness: 3,
		productionCulture:   1,
		cardCounts:          []int{2, 2, 2},
	})

	// Urban - Arena
	schools = append(schools, &CardSchool{
		schoolName: "Bread & Circuses",
		shortName:  "B. & C.",
		age:        1,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_URBAN,
			CARDTYPE_TECH_URBAN_ARENA},
		tech:                3,
		buildCost:           3,
		productionPower:     1,
		productionHappiness: 2,
		cardCounts:          []int{1, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Team Sports",
		shortName:  "T. Sp.",
		age:        2,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_URBAN,
			CARDTYPE_TECH_URBAN_ARENA},
		tech:                5,
		buildCost:           5,
		productionPower:     2,
		productionHappiness: 3,
		cardCounts:          []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Pro. Sports",
		shortName:  "P. Sp.",
		age:        3,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_URBAN,
			CARDTYPE_TECH_URBAN_ARENA},
		tech:                7,
		buildCost:           8,
		productionPower:     3,
		productionHappiness: 4,
		cardCounts:          []int{1, 1, 2},
	})

	// Urban - Library
	schools = append(schools, &CardSchool{
		schoolName: "Printing Press",
		shortName:  "P. Pr.",
		age:        1,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_URBAN,
			CARDTYPE_TECH_URBAN_LIBRARY},
		tech:              3,
		buildCost:         3,
		productionCulture: 1,
		productionTech:    1,
		cardCounts:        []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Journalism",
		shortName:  "Journ.",
		age:        2,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_URBAN,
			CARDTYPE_TECH_URBAN_LIBRARY},
		tech:              6,
		buildCost:         8,
		productionCulture: 2,
		productionTech:    2,
		cardCounts:        []int{1, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Multimedia",
		shortName:  "Multim.",
		age:        3,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_URBAN,
			CARDTYPE_TECH_URBAN_LIBRARY},
		tech:              9,
		buildCost:         11,
		productionCulture: 3,
		productionTech:    3,
		cardCounts:        []int{2, 2, 2},
	})

	// Urban - Theater
	schools = append(schools, &CardSchool{
		schoolName: "Drama",
		shortName:  "Drama",
		age:        1,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_URBAN,
			CARDTYPE_TECH_URBAN_THEATER},
		tech:                3,
		buildCost:           4,
		productionCulture:   2,
		productionHappiness: 1,
		cardCounts:          []int{1, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Opera",
		shortName:  "Opera",
		age:        2,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_URBAN,
			CARDTYPE_TECH_URBAN_THEATER},
		tech:                7,
		buildCost:           8,
		productionCulture:   3,
		productionHappiness: 1,
		cardCounts:          []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Movies",
		shortName:  "Movies",
		age:        3,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_URBAN,
			CARDTYPE_TECH_URBAN_THEATER},
		tech:                10,
		buildCost:           11,
		productionCulture:   4,
		productionHappiness: 1,
		cardCounts:          []int{2, 2, 2},
	})

	// Infantry
	schools = append(schools, &CardSchool{
		schoolName: "Warriors",
		shortName:  "Warriors",
		age:        0,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_MILLI,
			CARDTYPE_TECH_MILLI_INFANTRY},
		buildCost:       2,
		productionPower: 1,
		cardCounts:      []int{2, 3, 4},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Swordmen",
		shortName:  "Swordmen",
		age:        1,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_MILLI,
			CARDTYPE_TECH_MILLI_INFANTRY},
		tech:            4,
		buildCost:       3,
		productionPower: 2,
		cardCounts:      []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Riflemen",
		shortName:  "Riflemen",
		age:        2,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_MILLI,
			CARDTYPE_TECH_MILLI_INFANTRY},
		tech:            6,
		buildCost:       5,
		productionPower: 3,
		cardCounts:      []int{1, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Modern Infantry",
		shortName:  "M. Infa.",
		age:        3,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_MILLI,
			CARDTYPE_TECH_MILLI_INFANTRY},
		tech:            10,
		buildCost:       7,
		productionPower: 5,
		cardCounts:      []int{1, 2, 2},
	})

	// Cavalry
	schools = append(schools, &CardSchool{
		schoolName: "Knights",
		shortName:  "Knights",
		age:        1,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_MILLI,
			CARDTYPE_TECH_MILLI_CAVALRY},
		tech:            5,
		buildCost:       3,
		productionPower: 2,
		cardCounts:      []int{2, 2, 3},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Cavalrymen",
		shortName:  "Caval.",
		age:        2,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_MILLI,
			CARDTYPE_TECH_MILLI_CAVALRY},
		tech:            6,
		buildCost:       5,
		productionPower: 3,
		cardCounts:      []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Tanks",
		shortName:  "Tanks",
		age:        3,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_MILLI,
			CARDTYPE_TECH_MILLI_CAVALRY},
		tech:            9,
		buildCost:       7,
		productionPower: 5,
		cardCounts:      []int{1, 2, 2},
	})

	// Artilery
	schools = append(schools, &CardSchool{
		schoolName: "Cannon",
		shortName:  "Cannon",
		age:        2,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_MILLI,
			CARDTYPE_TECH_MILLI_ARTILLERY},
		tech:            6,
		buildCost:       5,
		productionPower: 3,
		cardCounts:      []int{2, 2, 3},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Rockets",
		shortName:  "Rockets",
		age:        3,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_MILLI,
			CARDTYPE_TECH_MILLI_ARTILLERY},
		tech:            8,
		buildCost:       7,
		productionPower: 5,
		cardCounts:      []int{1, 2, 2},
	})

	// Ari Force
	schools = append(schools, &CardSchool{
		schoolName: "Air Forces",
		shortName:  "Air F.",
		age:        3,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_MILLI,
			CARDTYPE_TECH_MILLI_AIRFORCE},
		tech:            12,
		buildCost:       7,
		productionPower: 5,
		cardCounts:      []int{2, 2, 3},
	})

	// Governments
	schools = append(schools, &CardSchool{
		schoolName: "Despotism",
		shortName:  "Despot.",
		age:        0,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_GOVERNMENT},
		productionWhiteToken: 4,
		productionRedToken:   2,
		productionUrbanLimit: 2,
		cardCounts:           []int{2, 3, 4},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Theocracy",
		shortName:  "Theoc.",
		age:        1,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_GOVERNMENT},
		tech:                 6,
		techRevolution:       1,
		productionWhiteToken: 4,
		productionRedToken:   3,
		productionUrbanLimit: 3,
		productionCulture:    1,
		productionPower:      1,
		productionHappiness:  1,
		cardCounts:           []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Monarchy",
		shortName:  "Monarchy",
		age:        1,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_GOVERNMENT},
		tech:                 8,
		techRevolution:       2,
		productionWhiteToken: 5,
		productionRedToken:   3,
		productionUrbanLimit: 3,
		cardCounts:           []int{1, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Republic",
		shortName:  "Republic",
		age:        2,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_GOVERNMENT},
		tech:                 13,
		techRevolution:       3,
		productionWhiteToken: 7,
		productionRedToken:   2,
		productionUrbanLimit: 3,
		cardCounts:           []int{1, 1, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Const. Monarchy",
		shortName:  "C. Mon.",
		age:        2,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_GOVERNMENT},
		tech:                 12,
		techRevolution:       6,
		productionWhiteToken: 6,
		productionRedToken:   4,
		productionUrbanLimit: 3,
		cardCounts:           []int{1, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Communism",
		shortName:  "Comm.",
		age:        3,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_GOVERNMENT},
		tech:                 19,
		techRevolution:       5,
		productionWhiteToken: 7,
		productionRedToken:   5,
		productionUrbanLimit: 4,
		productionHappiness:  -1,
		cardCounts:           []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Fundamentalism",
		shortName:  "Fundam.",
		age:        3,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_GOVERNMENT},
		tech:                 18,
		techRevolution:       7,
		productionWhiteToken: 6,
		productionRedToken:   5,
		productionUrbanLimit: 4,
		productionTech:       -2,
		productionPower:      5,
		cardCounts:           []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Democracy",
		shortName:  "Democ.",
		age:        3,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_GOVERNMENT},
		tech:                 17,
		techRevolution:       9,
		productionWhiteToken: 7,
		productionRedToken:   3,
		productionUrbanLimit: 4,
		productionCulture:    3,
		cardCounts:           []int{1, 2, 2},
	})

	// Special tech - military
	schools = append(schools, &CardSchool{
		schoolName: "Warfare",
		shortName:  "Warfare",
		age:        1,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_SPECIAL,
			CARDTYPE_TECH_SPECIAL_MILITARY},
		tech:               5,
		productionPower:    1,
		productionRedToken: 1,
		cardCounts:         []int{1, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Strategy",
		shortName:  "Strategy",
		age:        2,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_SPECIAL,
			CARDTYPE_TECH_SPECIAL_MILITARY},
		tech:               8,
		productionPower:    3,
		productionRedToken: 2,
		cardCounts:         []int{1, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Military Theory",
		shortName:  "M. Th.",
		age:        3,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_SPECIAL,
			CARDTYPE_TECH_SPECIAL_MILITARY},
		tech:               11,
		productionPower:    5,
		productionRedToken: 3,
		cardCounts:         []int{1, 2, 2},
	})

	// Special tech - civil
	schools = append(schools, &CardSchool{
		schoolName: "Code of Laws",
		shortName:  "C. of L.",
		age:        1,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_SPECIAL,
			CARDTYPE_TECH_SPECIAL_CIVIL},
		tech:                 6,
		productionWhiteToken: 1,
		cardCounts:           []int{1, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Justice System",
		shortName:  "J. Sys.",
		age:        2,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_SPECIAL,
			CARDTYPE_TECH_SPECIAL_CIVIL},
		tech:                 7,
		productionWhiteToken: 1,
		productionBlueToken:  3,
		cardCounts:           []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Civil Service",
		shortName:  "C. Serv.",
		age:        3,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_SPECIAL,
			CARDTYPE_TECH_SPECIAL_CIVIL},
		tech:                 10,
		productionWhiteToken: 2,
		productionBlueToken:  3,
		cardCounts:           []int{1, 1, 1},
	})

	// Special tech - colonize
	schools = append(schools, &CardSchool{
		schoolName: "Cartography",
		shortName:  "Carto.",
		age:        1,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_SPECIAL,
			CARDTYPE_TECH_SPECIAL_COLONIZE},
		tech:                    4,
		productionPower:         1,
		productionColonizePower: 2,
		cardCounts:              []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Navigation",
		shortName:  "Navi.",
		age:        2,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_SPECIAL,
			CARDTYPE_TECH_SPECIAL_COLONIZE},
		tech:                    6,
		productionPower:         2,
		productionColonizePower: 3,
		cardCounts:              []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Satellites",
		shortName:  "Sate.s",
		age:        3,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_SPECIAL,
			CARDTYPE_TECH_SPECIAL_COLONIZE},
		tech:                    8,
		productionPower:         3,
		productionColonizePower: 4,
		cardCounts:              []int{1, 1, 1},
	})

	// Special tech - construction
	schools = append(schools, &CardSchool{
		schoolName: "Masonry",
		shortName:  "Masonry",
		age:        1,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_SPECIAL,
			CARDTYPE_TECH_SPECIAL_CONSTRUCTION},
		tech:       3,
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Architecture",
		shortName:  "Archi.",
		age:        2,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_SPECIAL,
			CARDTYPE_TECH_SPECIAL_CONSTRUCTION},
		tech:       6,
		cardCounts: []int{1, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Engineering",
		shortName:  "Engi.",
		age:        3,
		cardType: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_SPECIAL,
			CARDTYPE_TECH_SPECIAL_CONSTRUCTION},
		tech:       9,
		cardCounts: []int{1, 1, 1},
	})

	// Wonders
	schools = append(schools, &CardSchool{
		schoolName:          "Hanging Garden",
		shortName:           "H. Gard.",
		age:                 0,
		cardType:            []CardType{CARDTYPE_WONDER},
		wonderBuildCosts:    []int{2, 2, 2},
		productionCulture:   1,
		productionHappiness: 2,
		cardCounts:          []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:        "Lib. of Alexandria",
		shortName:         "L. of A.",
		age:               0,
		cardType:          []CardType{CARDTYPE_WONDER},
		wonderBuildCosts:  []int{1, 4, 1},
		productionCulture: 1,
		productionTech:    1,
		cardCounts:        []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{ // TODO: other wonder effects
		schoolName:              "Colossus",
		shortName:               "Colossus",
		age:                     0,
		cardType:                []CardType{CARDTYPE_WONDER},
		wonderBuildCosts:        []int{3, 3},
		productionPower:         2,
		productionColonizePower: 1,
		cardCounts:              []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:           "Pyramid",
		shortName:            "Pyramid",
		age:                  0,
		cardType:             []CardType{CARDTYPE_WONDER},
		wonderBuildCosts:     []int{3, 2, 1},
		productionWhiteToken: 2,
		cardCounts:           []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{ // TODO: other wonder effects
		schoolName:          "Great Wall",
		shortName:           "G. Wall",
		age:                 1,
		cardType:            []CardType{CARDTYPE_WONDER},
		wonderBuildCosts:    []int{2, 2, 3, 2},
		productionCulture:   1,
		productionHappiness: 1,
		cardCounts:          []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{ // TODO: other wonder effects
		schoolName:          "St. Peter's Basilica",
		shortName:           "Basilica",
		age:                 1,
		cardType:            []CardType{CARDTYPE_WONDER},
		wonderBuildCosts:    []int{4, 4},
		productionCulture:   2,
		productionHappiness: 1,
		cardCounts:          []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:        "Univ. Carolina",
		shortName:         "U. Carol.",
		age:               1,
		cardType:          []CardType{CARDTYPE_WONDER},
		wonderBuildCosts:  []int{3, 3, 3},
		productionCulture: 1,
		productionTech:    2,
		cardCounts:        []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{ // TODO: other wonder effects
		schoolName:          "Taj Mahal",
		shortName:           "T. Mahal",
		age:                 1,
		cardType:            []CardType{CARDTYPE_WONDER},
		wonderBuildCosts:    []int{2, 4, 2},
		productionCulture:   3,
		productionBlueToken: 1,
		cardCounts:          []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{ // TODO: other wonder effects
		schoolName:       "Transcont. RR",
		shortName:        "T. RR",
		age:              2,
		cardType:         []CardType{CARDTYPE_WONDER},
		wonderBuildCosts: []int{3, 3, 3, 3},
		productionPower:  4,
		cardCounts:       []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:          "Eiffel Tower",
		shortName:           "Eiffel T.",
		age:                 2,
		cardType:            []CardType{CARDTYPE_WONDER},
		wonderBuildCosts:    []int{3, 7, 3},
		productionCulture:   4,
		productionHappiness: 1,
		cardCounts:          []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:           "Kremlin",
		shortName:            "Kremlin",
		age:                  2,
		cardType:             []CardType{CARDTYPE_WONDER},
		wonderBuildCosts:     []int{4, 4, 4},
		productionCulture:    3,
		productionWhiteToken: 1,
		productionRedToken:   1,
		productionHappiness:  -1,
		cardCounts:           []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{ // TODO: other wonder effects
		schoolName:       "Ocean Liner Service",
		shortName:        "O. L. S.",
		age:              2,
		cardType:         []CardType{CARDTYPE_WONDER},
		wonderBuildCosts: []int{4, 2, 2, 4},
		cardCounts:       []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{ // TODO: other wonder effects
		schoolName:       "Hollywood",
		shortName:        "Hollyw.",
		age:              3,
		cardType:         []CardType{CARDTYPE_WONDER},
		wonderBuildCosts: []int{5, 6, 5},
		cardCounts:       []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{ // TODO: other wonder effects
		schoolName:       "Internet",
		shortName:        "Internet",
		age:              3,
		cardType:         []CardType{CARDTYPE_WONDER},
		wonderBuildCosts: []int{2, 3, 4, 3, 2},
		cardCounts:       []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{ // TODO: other wonder effects
		schoolName:       "First Space Flight",
		shortName:        "Sp. Fl.",
		age:              3,
		cardType:         []CardType{CARDTYPE_WONDER},
		wonderBuildCosts: []int{1, 2, 4, 9},
		cardCounts:       []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{ // TODO: other wonder effects
		schoolName:       "Fast Food Chains",
		shortName:        "F. F. C.",
		age:              3,
		cardType:         []CardType{CARDTYPE_WONDER},
		wonderBuildCosts: []int{4, 4, 4, 4},
		cardCounts:       []int{1, 1, 1},
	})

	// Leaders all have special abilities
	schools = append(schools, &CardSchool{
		schoolName:         "Julius Caesar",
		shortName:          "Caesar",
		age:                0,
		cardType:           []CardType{CARDTYPE_LEADER},
		productionPower:    1,
		productionRedToken: 1,
		cardCounts:         []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:          "Homer",
		shortName:           "Homer",
		age:                 0,
		cardType:            []CardType{CARDTYPE_LEADER},
		productionHappiness: 1,
		cardCounts:          []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Moses",
		shortName:  "Moses",
		age:        0,
		cardType:   []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Hammurabi",
		shortName:  "Hammu.",
		age:        0,
		cardType:   []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Aristotle",
		shortName:  "Arist.",
		age:        0,
		cardType:   []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Alexander the Great",
		shortName:  "Alex.",
		age:        0,
		cardType:   []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Michelangelo",
		shortName:  "Michel.",
		age:        1,
		cardType:   []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:         "Joan of Arc",
		shortName:          "J. of A.",
		age:                1,
		cardType:           []CardType{CARDTYPE_LEADER},
		productionRedToken: 1,
		productionCulture:  1,
		cardCounts:         []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Leonardo Da Vinci",
		shortName:  "Da Vinci",
		age:        1,
		cardType:   []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Genghis Khan",
		shortName:  "Genghis.",
		age:        1,
		cardType:   []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Christopher Columbus",
		shortName:  "Columbus",
		age:        1,
		cardType:   []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Frederick Barbarosa",
		shortName:  "Barbaros.",
		age:        1,
		cardType:   []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:          "William Shakespeare",
		shortName:           "Shakesp.",
		age:                 2,
		cardType:            []CardType{CARDTYPE_LEADER},
		productionHappiness: 1,
		cardCounts:          []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "James Cook",
		shortName:  "Cook",
		age:        2,
		cardType:   []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:         "Napoleon Bonaparte",
		shortName:          "Napoleon",
		age:                2,
		cardType:           []CardType{CARDTYPE_LEADER},
		productionRedToken: 2,
		cardCounts:         []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:         "Maximillien Robespierre",
		shortName:          "Robesp.",
		age:                2,
		cardType:           []CardType{CARDTYPE_LEADER},
		productionRedToken: 1,
		cardCounts:         []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "J.S. Bach",
		shortName:  "Bach",
		age:        2,
		cardType:   []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Isaac Newton",
		shortName:  "Newton",
		age:        2,
		cardType:   []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Albert Einstein",
		shortName:  "Einstein",
		age:        3,
		cardType:   []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:        "Mahatma Gandhi",
		shortName:         "Gandhi",
		age:               3,
		cardType:          []CardType{CARDTYPE_LEADER},
		productionCulture: 2,
		cardCounts:        []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:          "Chalie Chaplin",
		shortName:           "Chaplin",
		age:                 3,
		cardType:            []CardType{CARDTYPE_LEADER},
		productionHappiness: 2,
		cardCounts:          []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Bill Gates",
		shortName:  "B. Gates",
		age:        3,
		cardType:   []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Winston Churchill",
		shortName:  "Church.",
		age:        3,
		cardType:   []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Sid Meier",
		shortName:  "S. Meier",
		age:        3,
		cardType:   []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
	})

	// Actions
	schools = append(schools, &CardSchool{
		schoolName:  "Breakthrough - Age I",
		age:         1,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 2,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Breakthrough - Age II",
		age:         2,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 3,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Cultural Heritage - Age A",
		age:        0,
		cardType:   []CardType{CARDTYPE_ACTION},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Cultural Heritage - Age I",
		age:        1,
		cardType:   []CardType{CARDTYPE_ACTION},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Efficient Upgrade - Age II",
		age:         2,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 3,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Efficient Upgrade - Age III",
		age:         3,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 4,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Endowment for Arts",
		age:        3,
		cardType:   []CardType{CARDTYPE_ACTION},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Engineering Genius - Age A",
		age:         0,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 2,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Engineering Genius - Age I",
		age:         1,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 3,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Engineering Genius - Age II",
		age:         2,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 4,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Engineering Genius - Age III",
		age:         3,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 5,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Frugality - Age A",
		age:         0,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 1,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Frugality - Age I",
		age:         1,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 2,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Frugality - Age II",
		age:         1,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 3,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Military Build-Up",
		age:        3,
		cardType:   []CardType{CARDTYPE_ACTION},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Patriotism - Age A",
		age:         0,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 1,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Patriotism - Age I",
		age:         1,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 2,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Patriotism - Age II",
		age:         2,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 3,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Patriotism - Age III",
		age:         3,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 4,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Reserves - Age I",
		age:         1,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 2,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Reserves - Age III",
		age:         2,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 3,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Reserves - Age III",
		age:         3,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 4,
		cardCounts:  []int{3, 3, 3},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Revolutionary Idea - Age II",
		age:         2,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 4,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Revolutionary Idea - Age III",
		age:         3,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 6,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Rich Land - Age A",
		age:         0,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 1,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Rich Land - Age I",
		age:         1,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 2,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Rich Land - Age II",
		age:         2,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 3,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Stockpile",
		age:        0,
		cardType:   []CardType{CARDTYPE_ACTION},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Urban Growth - Age A",
		age:         0,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 1,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Urban Growth - Age I",
		age:         1,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 2,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Urban Growth - Age II",
		age:         2,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 3,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Urban Growth - Age III",
		age:         3,
		cardType:    []CardType{CARDTYPE_ACTION},
		actionBonus: 4,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Wave of Nationalism",
		age:        2,
		cardType:   []CardType{CARDTYPE_ACTION},
		cardCounts: []int{1, 1, 1},
	})

	result = make(map[int]*CardSchool)
	for i, s := range schools {
		s.schoolId = i + 1
		result[s.schoolId] = s
	}
	return
}
