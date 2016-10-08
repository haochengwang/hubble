package main

func InitBasicCardSchools() (result map[int]*CardSchool) {
	schools := make([]*CardSchool, 0)

	// Farms
	schools = append(schools, &CardSchool{
		schoolName:     "Agriculture",
		shortName:      "Agri.",
		age:            0,
		cardTypes:      []CardType{CARDTYPE_TECH, CARDTYPE_TECH_FARM},
		buildCost:      2,
		productionCrop: 1,
		cardCounts:     []int{2, 3, 4},
	})
	schools = append(schools, &CardSchool{
		schoolName:     "Irrigation",
		shortName:      "Irri.",
		age:            1,
		cardTypes:      []CardType{CARDTYPE_TECH, CARDTYPE_TECH_FARM},
		tech:           3,
		buildCost:      4,
		productionCrop: 2,
		cardCounts:     []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:     "Selective Breeding",
		shortName:      "S. Brd.",
		age:            2,
		cardTypes:      []CardType{CARDTYPE_TECH, CARDTYPE_TECH_FARM},
		tech:           5,
		buildCost:      6,
		productionCrop: 3,
		cardCounts:     []int{1, 2, 3},
	})
	schools = append(schools, &CardSchool{
		schoolName:     "Mech. Agriculture",
		shortName:      "M. Agr.",
		age:            3,
		cardTypes:      []CardType{CARDTYPE_TECH, CARDTYPE_TECH_FARM},
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
		cardTypes:          []CardType{CARDTYPE_TECH, CARDTYPE_TECH_MINE},
		buildCost:          2,
		productionResource: 1,
		cardCounts:         []int{2, 3, 4},
	})
	schools = append(schools, &CardSchool{
		schoolName:         "Iron",
		shortName:          "Iron",
		age:                1,
		cardTypes:          []CardType{CARDTYPE_TECH, CARDTYPE_TECH_MINE},
		tech:               5,
		buildCost:          5,
		productionResource: 2,
		cardCounts:         []int{2, 2, 3},
	})
	schools = append(schools, &CardSchool{
		schoolName:         "Coal",
		shortName:          "Coal",
		age:                2,
		cardTypes:          []CardType{CARDTYPE_TECH, CARDTYPE_TECH_MINE},
		tech:               7,
		buildCost:          8,
		productionResource: 3,
		cardCounts:         []int{1, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:         "Oil",
		shortName:          "Oil",
		age:                3,
		cardTypes:          []CardType{CARDTYPE_TECH, CARDTYPE_TECH_MINE},
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		shortName:  "Religion",
		age:        0,
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_MILI,
			CARDTYPE_TECH_MILI_INFANTRY},
		buildCost:       2,
		productionPower: 1,
		cardCounts:      []int{2, 3, 4},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Swordmen",
		shortName:  "Swordmen",
		age:        1,
		cardTypes: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_MILI,
			CARDTYPE_TECH_MILI_INFANTRY},
		tech:            4,
		buildCost:       3,
		productionPower: 2,
		cardCounts:      []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Riflemen",
		shortName:  "Riflemen",
		age:        2,
		cardTypes: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_MILI,
			CARDTYPE_TECH_MILI_INFANTRY},
		tech:            6,
		buildCost:       5,
		productionPower: 3,
		cardCounts:      []int{1, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Modern Infantry",
		shortName:  "M. Infa.",
		age:        3,
		cardTypes: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_MILI,
			CARDTYPE_TECH_MILI_INFANTRY},
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
		cardTypes: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_MILI,
			CARDTYPE_TECH_MILI_CAVALRY},
		tech:            5,
		buildCost:       3,
		productionPower: 2,
		cardCounts:      []int{2, 2, 3},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Cavalrymen",
		shortName:  "Caval.",
		age:        2,
		cardTypes: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_MILI,
			CARDTYPE_TECH_MILI_CAVALRY},
		tech:            6,
		buildCost:       5,
		productionPower: 3,
		cardCounts:      []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Tanks",
		shortName:  "Tanks",
		age:        3,
		cardTypes: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_MILI,
			CARDTYPE_TECH_MILI_CAVALRY},
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
		cardTypes: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_MILI,
			CARDTYPE_TECH_MILI_ARTILLERY},
		tech:            6,
		buildCost:       5,
		productionPower: 3,
		cardCounts:      []int{2, 2, 3},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Rockets",
		shortName:  "Rockets",
		age:        3,
		cardTypes: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_MILI,
			CARDTYPE_TECH_MILI_ARTILLERY},
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
		cardTypes: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_MILI,
			CARDTYPE_TECH_MILI_AIRFORCE},
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_SPECIAL,
			CARDTYPE_TECH_SPECIAL_CONSTRUCTION},
		tech:       3,
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Architecture",
		shortName:  "Archi.",
		age:        2,
		cardTypes: []CardType{CARDTYPE_TECH,
			CARDTYPE_TECH_SPECIAL,
			CARDTYPE_TECH_SPECIAL_CONSTRUCTION},
		tech:       6,
		cardCounts: []int{1, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Engineering",
		shortName:  "Engi.",
		age:        3,
		cardTypes: []CardType{CARDTYPE_TECH,
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
		cardTypes:           []CardType{CARDTYPE_WONDER},
		wonderBuildCosts:    []int{2, 2, 2},
		productionCulture:   1,
		productionHappiness: 2,
		cardCounts:          []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:        "Lib. of Alexandria",
		shortName:         "L. of A.",
		age:               0,
		cardTypes:         []CardType{CARDTYPE_WONDER},
		wonderBuildCosts:  []int{1, 4, 1},
		productionCulture: 1,
		productionTech:    1,
		cardCounts:        []int{1, 1, 1},
		traits:            []int{TRAIT_LIB_OF_ALEXANDRIA},
	})
	schools = append(schools, &CardSchool{ // TODO: other wonder effects
		schoolName:              "Colossus",
		shortName:               "Colossus",
		age:                     0,
		cardTypes:               []CardType{CARDTYPE_WONDER},
		wonderBuildCosts:        []int{3, 3},
		productionPower:         2,
		productionColonizePower: 1,
		cardCounts:              []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:           "Pyramid",
		shortName:            "Pyramid",
		age:                  0,
		cardTypes:            []CardType{CARDTYPE_WONDER},
		wonderBuildCosts:     []int{3, 2, 1},
		productionWhiteToken: 1,
		cardCounts:           []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{ // TODO: other wonder effects
		schoolName:          "Great Wall",
		shortName:           "G. Wall",
		age:                 1,
		cardTypes:           []CardType{CARDTYPE_WONDER},
		wonderBuildCosts:    []int{2, 2, 3, 2},
		productionCulture:   1,
		productionHappiness: 1,
		cardCounts:          []int{1, 1, 1},
		traits:              []int{TRAIT_GREAT_WALL},
	})
	schools = append(schools, &CardSchool{ // TODO: other wonder effects
		schoolName:          "St. Peter's Basilica",
		shortName:           "Basilica",
		age:                 1,
		cardTypes:           []CardType{CARDTYPE_WONDER},
		wonderBuildCosts:    []int{4, 4},
		productionCulture:   2,
		productionHappiness: 1,
		cardCounts:          []int{1, 1, 1},
		traits:              []int{TRAIT_ST_PETERS_BASILICA},
	})
	schools = append(schools, &CardSchool{
		schoolName:        "Univ. Carolina",
		shortName:         "U. Carol.",
		age:               1,
		cardTypes:         []CardType{CARDTYPE_WONDER},
		wonderBuildCosts:  []int{3, 3, 3},
		productionCulture: 1,
		productionTech:    2,
		cardCounts:        []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{ // TODO: other wonder effects
		schoolName:          "Taj Mahal",
		shortName:           "T. Mahal",
		age:                 1,
		cardTypes:           []CardType{CARDTYPE_WONDER},
		wonderBuildCosts:    []int{2, 4, 2},
		productionCulture:   3,
		productionBlueToken: 1,
		cardCounts:          []int{1, 1, 1},
		traits:              []int{TRAIT_TAJ_MAHAL},
	})
	schools = append(schools, &CardSchool{ // TODO: other wonder effects
		schoolName:       "Transcont. RR",
		shortName:        "T. RR",
		age:              2,
		cardTypes:        []CardType{CARDTYPE_WONDER},
		wonderBuildCosts: []int{3, 3, 3, 3},
		productionPower:  4,
		cardCounts:       []int{1, 1, 1},
		traits:           []int{TRAIT_TRANSCONT_RR},
	})
	schools = append(schools, &CardSchool{
		schoolName:          "Eiffel Tower",
		shortName:           "Eiffel T.",
		age:                 2,
		cardTypes:           []CardType{CARDTYPE_WONDER},
		wonderBuildCosts:    []int{3, 7, 3},
		productionCulture:   4,
		productionHappiness: 1,
		cardCounts:          []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:           "Kremlin",
		shortName:            "Kremlin",
		age:                  2,
		cardTypes:            []CardType{CARDTYPE_WONDER},
		wonderBuildCosts:     []int{4, 4, 4},
		productionCulture:    2,
		productionWhiteToken: 1,
		productionRedToken:   1,
		productionHappiness:  -1,
		cardCounts:           []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{ // TODO: other wonder effects
		schoolName:       "Ocean Liner Service",
		shortName:        "O. L. S.",
		age:              2,
		cardTypes:        []CardType{CARDTYPE_WONDER},
		wonderBuildCosts: []int{4, 2, 2, 4},
		cardCounts:       []int{1, 1, 1},
		traits:           []int{TRAIT_OCEAN_LINER_SERVICE},
	})
	schools = append(schools, &CardSchool{ // TODO: other wonder effects
		schoolName:       "Hollywood",
		shortName:        "Hollyw.",
		age:              3,
		cardTypes:        []CardType{CARDTYPE_WONDER},
		wonderBuildCosts: []int{5, 6, 5},
		cardCounts:       []int{1, 1, 1},
		traits:           []int{TRAIT_HOLLYWOOD},
	})
	schools = append(schools, &CardSchool{ // TODO: other wonder effects
		schoolName:       "Internet",
		shortName:        "Internet",
		age:              3,
		cardTypes:        []CardType{CARDTYPE_WONDER},
		wonderBuildCosts: []int{2, 3, 4, 3, 2},
		cardCounts:       []int{1, 1, 1},
		traits:           []int{TRAIT_INTERNET},
	})
	schools = append(schools, &CardSchool{ // TODO: other wonder effects
		schoolName:       "First Space Flight",
		shortName:        "Sp. Fl.",
		age:              3,
		cardTypes:        []CardType{CARDTYPE_WONDER},
		wonderBuildCosts: []int{1, 2, 4, 9},
		cardCounts:       []int{1, 1, 1},
		traits:           []int{TRAIT_FIRST_SPACE_FLIGHT},
	})
	schools = append(schools, &CardSchool{ // TODO: other wonder effects
		schoolName:       "Fast Food Chains",
		shortName:        "F. F. C.",
		age:              3,
		cardTypes:        []CardType{CARDTYPE_WONDER},
		wonderBuildCosts: []int{4, 4, 4, 4},
		cardCounts:       []int{1, 1, 1},
		traits:           []int{TRAIT_FAST_FOOD_CHAINS},
	})

	// Leaders all have special abilities
	schools = append(schools, &CardSchool{
		schoolName:         "Julius Caesar",
		shortName:          "Caesar",
		age:                0,
		cardTypes:          []CardType{CARDTYPE_LEADER},
		productionPower:    1,
		productionRedToken: 1,
		cardCounts:         []int{1, 1, 1},
		traits:             []int{TRAIT_JULIUS_CAESAR},
	})
	schools = append(schools, &CardSchool{
		schoolName:          "Homer",
		shortName:           "Homer",
		age:                 0,
		cardTypes:           []CardType{CARDTYPE_LEADER},
		productionHappiness: 1,
		cardCounts:          []int{1, 1, 1},
		traits:              []int{TRAIT_HOMER},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Moses",
		shortName:  "Moses",
		age:        0,
		cardTypes:  []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
		traits:     []int{TRAIT_MOSES},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Hammurabi",
		shortName:  "Hammu.",
		age:        0,
		cardTypes:  []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},

		traits: []int{TRAIT_HAMMURABI},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Aristotle",
		shortName:  "Arist.",
		age:        0,
		cardTypes:  []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
		traits:     []int{TRAIT_ARISTOTLE},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Alexander the Great",
		shortName:  "Alex.",
		age:        0,
		cardTypes:  []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
		traits:     []int{TRAIT_ALEXANDER_THE_GREAT},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Michelangelo",
		shortName:  "Michel.",
		age:        1,
		cardTypes:  []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
		traits:     []int{TRAIT_MICHELANGELO},
	})
	schools = append(schools, &CardSchool{
		schoolName:         "Joan of Arc",
		shortName:          "J. of A.",
		age:                1,
		cardTypes:          []CardType{CARDTYPE_LEADER},
		productionRedToken: 1,
		productionCulture:  1,
		cardCounts:         []int{1, 1, 1},
		traits:             []int{TRAIT_JOAN_OF_ARC},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Leonardo Da Vinci",
		shortName:  "Da Vinci",
		age:        1,
		cardTypes:  []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
		traits:     []int{TRAIT_LIB_LAB_AMPLIFY, TRAIT_LEONARDO_DA_VINCI},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Genghis Khan",
		shortName:  "Genghis.",
		age:        1,
		cardTypes:  []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
		traits:     []int{TRAIT_GENGHIS_KHAN},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Christopher Columbus",
		shortName:  "Columbus",
		age:        1,
		cardTypes:  []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
		traits:     []int{TRAIT_CHRISTOPHER_COLUMBUS},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Frederick Barbarosa",
		shortName:  "Barbaros.",
		age:        1,
		cardTypes:  []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
		traits:     []int{TRAIT_FREDERICK_BARBAROSSA},
	})
	schools = append(schools, &CardSchool{
		schoolName:          "William Shakespeare",
		shortName:           "Shakesp.",
		age:                 2,
		cardTypes:           []CardType{CARDTYPE_LEADER},
		productionHappiness: 1,
		cardCounts:          []int{1, 1, 1},
		traits:              []int{TRAIT_WILLIAM_SHAKESPEARE},
	})
	schools = append(schools, &CardSchool{
		schoolName: "James Cook",
		shortName:  "Cook",
		age:        2,
		cardTypes:  []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
		traits:     []int{TRAIT_JAMES_COOK},
	})
	schools = append(schools, &CardSchool{
		schoolName:         "Napoleon Bonaparte",
		shortName:          "Napoleon",
		age:                2,
		cardTypes:          []CardType{CARDTYPE_LEADER},
		productionRedToken: 2,
		cardCounts:         []int{1, 1, 1},
		traits:             []int{TRAIT_NAPOLEON_BONAPARTE},
	})
	schools = append(schools, &CardSchool{
		schoolName:         "Maximillien Robespierre",
		shortName:          "Robesp.",
		age:                2,
		cardTypes:          []CardType{CARDTYPE_LEADER},
		productionRedToken: 1,
		cardCounts:         []int{1, 1, 1},
		traits:             []int{TRAIT_MAXIMILLIEN_ROBESPIERRE},
	})
	schools = append(schools, &CardSchool{
		schoolName: "J.S. Bach",
		shortName:  "Bach",
		age:        2,
		cardTypes:  []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
		traits:     []int{TRAIT_J_S_BACH},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Isaac Newton",
		shortName:  "Newton",
		age:        2,
		cardTypes:  []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
		traits:     []int{TRAIT_LIB_LAB_AMPLIFY, TRAIT_ISAAC_NEWTON},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Albert Einstein",
		shortName:  "Einstein",
		age:        3,
		cardTypes:  []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
		traits:     []int{TRAIT_LIB_LAB_AMPLIFY, TRAIT_ALBERT_EINSTEIN},
	})
	schools = append(schools, &CardSchool{
		schoolName:        "Mahatma Gandhi",
		shortName:         "Gandhi",
		age:               3,
		cardTypes:         []CardType{CARDTYPE_LEADER},
		productionCulture: 2,
		cardCounts:        []int{1, 1, 1},
		traits:            []int{TRAIT_MAHATMA_GANDHI},
	})
	schools = append(schools, &CardSchool{
		schoolName:          "Charlie Chaplin",
		shortName:           "Chaplin",
		age:                 3,
		cardTypes:           []CardType{CARDTYPE_LEADER},
		productionHappiness: 2,
		cardCounts:          []int{1, 1, 1},
		traits:              []int{TRAIT_CHARLIE_CHAPLIN},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Bill Gates",
		shortName:  "B. Gates",
		age:        3,
		cardTypes:  []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
		traits:     []int{TRAIT_BILL_GATES},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Winston Churchill",
		shortName:  "Church.",
		age:        3,
		cardTypes:  []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
		traits:     []int{TRAIT_WINSTON_CHURCHILL},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Sid Meier",
		shortName:  "S. Meier",
		age:        3,
		cardTypes:  []CardType{CARDTYPE_LEADER},
		cardCounts: []int{1, 1, 1},
		traits:     []int{TRAIT_SID_MEIER},
	})

	// Actions
	schools = append(schools, &CardSchool{
		schoolName: "Breakthrough",
		age:        1,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_BREAKTHROUGH},
		actionBonus: 2,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Breakthrough",
		age:        2,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_BREAKTHROUGH},
		actionBonus: 3,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Cultural Heritage",
		age:        0,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_CULTURAL_HERITAGE},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Cultural Heritage",
		age:        1,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_CULTURAL_HERITAGE},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Efficient Upgrade",
		age:        2,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_EFFICIENT_UPGRADE},
		actionBonus: 3,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Efficient Upgrade",
		age:        3,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_EFFICIENT_UPGRADE},
		actionBonus: 4,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Endowment for Arts",
		age:        3,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_ENDOWMENT_FOR_ARTS},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Engineering Genius",
		age:        0,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_ENGINEERING_GENIUS},
		actionBonus: 2,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Engineering Genius",
		age:        1,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_ENGINEERING_GENIUS},
		actionBonus: 3,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Engineering Genius",
		age:        2,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_ENGINEERING_GENIUS},
		actionBonus: 4,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Engineering Genius",
		age:        3,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_ENGINEERING_GENIUS},
		actionBonus: 5,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Frugality",
		age:        0,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_FRUGALITY},
		actionBonus: 1,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Frugality",
		age:        1,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_FRUGALITY},
		actionBonus: 2,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Frugality",
		age:        1,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_FRUGALITY},
		actionBonus: 3,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Military Build-Up",
		age:        3,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_MILITARY_BUILD_UP},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Patriotism",
		age:        0,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_PATRIOTISM},
		actionBonus: 1,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Patriotism",
		age:        1,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_PATRIOTISM},
		actionBonus: 2,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Patriotism",
		age:        2,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_PATRIOTISM},
		actionBonus: 3,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Patriotism",
		age:        3,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_PATRIOTISM},
		actionBonus: 4,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Reserves",
		age:        1,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_RESERVES},
		actionBonus: 2,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Reserves",
		age:        2,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_RESERVES},
		actionBonus: 3,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Reserves",
		age:        3,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_RESERVES},
		actionBonus: 4,
		cardCounts:  []int{3, 3, 3},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Revolutionary Idea",
		age:        2,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_REVOLUTIONARY_IDEA},
		actionBonus: 4,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Revolutionary Idea",
		age:        3,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_REVOLUTIONARY_IDEA},
		actionBonus: 6,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Rich Land",
		age:        0,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_RICH_LAND},
		actionBonus: 1,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Rich Land",
		age:        1,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_RICH_LAND},
		actionBonus: 2,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Rich Land",
		age:        2,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_RICH_LAND},
		actionBonus: 3,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Stockpile",
		age:        0,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_STOCKPILE},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Urban Growth",
		age:        0,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_URBAN_GROWTH},
		actionBonus: 1,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Urban Growth",
		age:        1,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_URBAN_GROWTH},
		actionBonus: 2,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Urban Growth",
		age:        2,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_URBAN_GROWTH},
		actionBonus: 3,
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Urban Growth",
		age:        3,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_URBAN_GROWTH},
		actionBonus: 4,
		cardCounts:  []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Wave of Nationalism",
		age:        2,
		cardTypes: []CardType{CARDTYPE_ACTION,
			CARDTYPE_ACTION_WAVE_OF_NATIONALISM},
		cardCounts: []int{1, 1, 1},
	})

	// Defence / Colonise
	schools = append(schools, &CardSchool{
		schoolName: "Def+2 / Col+1",
		age:        1,
		cardTypes:  []CardType{CARDTYPE_DEFCOL},
		cardCounts: []int{6, 6, 6},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Def+4 / Col+2",
		age:        2,
		cardTypes:  []CardType{CARDTYPE_DEFCOL},
		cardCounts: []int{6, 6, 6},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Def+6 / Col+3",
		age:        3,
		cardTypes:  []CardType{CARDTYPE_DEFCOL},
		cardCounts: []int{6, 6, 6},
	})

	// Aggressions
	schools = append(schools, &CardSchool{
		schoolName:     "Enslave",
		age:            1,
		cardTypes:      []CardType{CARDTYPE_AGGRESSION},
		miliActionCost: 2,
		cardCounts:     []int{2, 2, 2},
		traits:         []int{TRAIT_ENSLAVE},
	})
	schools = append(schools, &CardSchool{
		schoolName:     "Plunder",
		age:            1,
		cardTypes:      []CardType{CARDTYPE_AGGRESSION},
		miliActionCost: 1,
		actionBonus:    3,
		cardCounts:     []int{2, 2, 2},
		traits:         []int{TRAIT_PLUNDER},
	})
	schools = append(schools, &CardSchool{
		schoolName:     "Raid",
		age:            1,
		cardTypes:      []CardType{CARDTYPE_AGGRESSION},
		miliActionCost: 1,
		cardCounts:     []int{2, 2, 2},
		traits:         []int{TRAIT_RAID},
	})
	schools = append(schools, &CardSchool{
		schoolName:     "Annex",
		age:            2,
		cardTypes:      []CardType{CARDTYPE_AGGRESSION},
		miliActionCost: 2,
		cardCounts:     []int{1, 1, 1},
		traits:         []int{TRAIT_ANNEX},
	})
	schools = append(schools, &CardSchool{
		schoolName:     "Infiltrate",
		age:            2,
		cardTypes:      []CardType{CARDTYPE_AGGRESSION},
		miliActionCost: 2,
		cardCounts:     []int{2, 2, 2},
		traits:         []int{TRAIT_INFILTRATE},
	})
	schools = append(schools, &CardSchool{
		schoolName:     "Plunder",
		age:            2,
		cardTypes:      []CardType{CARDTYPE_AGGRESSION},
		miliActionCost: 1,
		actionBonus:    5,
		cardCounts:     []int{2, 2, 2},
		traits:         []int{TRAIT_PLUNDER},
	})
	schools = append(schools, &CardSchool{
		schoolName:     "Raid",
		age:            2,
		cardTypes:      []CardType{CARDTYPE_AGGRESSION},
		miliActionCost: 2,
		cardCounts:     []int{2, 2, 2},
		traits:         []int{TRAIT_RAID},
	})
	schools = append(schools, &CardSchool{
		schoolName:     "Spy",
		age:            2,
		cardTypes:      []CardType{CARDTYPE_AGGRESSION},
		miliActionCost: 1,
		actionBonus:    5,
		cardCounts:     []int{2, 2, 2},
		traits:         []int{TRAIT_SPY},
	})
	schools = append(schools, &CardSchool{
		schoolName:     "Armed Intervention",
		age:            3,
		cardTypes:      []CardType{CARDTYPE_AGGRESSION},
		miliActionCost: 2,
		actionBonus:    7,
		cardCounts:     []int{4, 4, 4},
		traits:         []int{TRAIT_ARMED_INTERVENTION},
	})
	schools = append(schools, &CardSchool{
		schoolName:     "Plunder",
		age:            3,
		cardTypes:      []CardType{CARDTYPE_AGGRESSION},
		miliActionCost: 1,
		actionBonus:    7,
		cardCounts:     []int{2, 2, 2},
		traits:         []int{TRAIT_PLUNDER},
	})
	schools = append(schools, &CardSchool{
		schoolName:     "Raid",
		age:            3,
		cardTypes:      []CardType{CARDTYPE_AGGRESSION},
		miliActionCost: 3,
		cardCounts:     []int{2, 2, 2},
		traits:         []int{TRAIT_RAID},
	})

	// Wars
	schools = append(schools, &CardSchool{
		schoolName:     "War over Technology",
		age:            2,
		cardTypes:      []CardType{CARDTYPE_WAR},
		miliActionCost: 2,
		cardCounts:     []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:     "War over Territory",
		age:            2,
		cardTypes:      []CardType{CARDTYPE_WAR},
		miliActionCost: 2,
		cardCounts:     []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:     "War over Culture",
		age:            3,
		cardTypes:      []CardType{CARDTYPE_WAR},
		miliActionCost: 3,
		cardCounts:     []int{6, 6, 6},
	})

	// Pacts
	schools = append(schools, &CardSchool{
		schoolName:         "Open Borders Agreement",
		age:                1,
		cardTypes:          []CardType{CARDTYPE_PACT},
		symmetric:          true,
		canAttack:          true,
		endOnAttack:        false,
		productionRedToken: 1,
		traits:             []int{TRAIT_OPEN_BORDER_AGGREMENT},
		cardCounts:         []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Trade Routes Agreement",
		age:         1,
		cardTypes:   []CardType{CARDTYPE_PACT},
		symmetric:   false,
		canAttack:   true,
		endOnAttack: false,
		traits:      []int{TRAIT_TRADE_ROUTE_AGGREMENT_A},
		cardCounts:  []int{1, 1, 1},
		bSide: &CardSchool{
			schoolName:  "Trade Routes Agreement - Part B",
			age:         1,
			cardTypes:   []CardType{CARDTYPE_PACT},
			symmetric:   false,
			canAttack:   true,
			endOnAttack: false,
			traits:      []int{TRAIT_TRADE_ROUTE_AGGREMENT_B},
		},
	})
	schools = append(schools, &CardSchool{
		schoolName:         "Acceptance of Supremacy",
		age:                2,
		cardTypes:          []CardType{CARDTYPE_PACT},
		productionResource: 1,
		symmetric:          false,
		canAttack:          false,
		endOnAttack:        false,
		traits:             []int{TRAIT_ACCEPTANCE_OF_SUPREMACY_A},
		bSide: &CardSchool{
			schoolName:         "Acceptance of Supremacy - Part B",
			age:                2,
			cardTypes:          []CardType{CARDTYPE_PACT},
			productionResource: -1,
			symmetric:          false,
			canAttack:          false,
			endOnAttack:        false,
			traits:             []int{TRAIT_ACCEPTANCE_OF_SUPREMACY_B},
		},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:         "International Trade Agreement",
		age:                2,
		cardTypes:          []CardType{CARDTYPE_PACT},
		productionResource: 1,
		symmetric:          false,
		canAttack:          true,
		endOnAttack:        false,
		traits:             []int{TRAIT_INTERNATIONAL_TRADE_AGGREMENT_A},
		bSide: &CardSchool{
			schoolName:     "International Trade Agreement - Part B",
			age:            2,
			cardTypes:      []CardType{CARDTYPE_PACT},
			productionCrop: 1,
			symmetric:      false,
			canAttack:      true,
			endOnAttack:    false,
			traits:         []int{TRAIT_INTERNATIONAL_TRADE_AGGREMENT_B},
		},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:        "Promise of Military Protection",
		age:               2,
		cardTypes:         []CardType{CARDTYPE_PACT},
		productionCulture: 1,
		symmetric:         false,
		canAttack:         true,
		endOnAttack:       true,
		traits:            []int{TRAIT_PROMISE_OF_MILITARY_PROTECTION_A},
		bSide: &CardSchool{
			schoolName:        "Promise of Military Protection - Part B",
			age:               2,
			cardTypes:         []CardType{CARDTYPE_PACT},
			productionPower:   4,
			productionCulture: -1,
			symmetric:         false,
			canAttack:         true,
			endOnAttack:       true,
			traits:            []int{TRAIT_PROMISE_OF_MILITARY_PROTECTION_B},
		},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "Scientific Cooperation",
		age:         2,
		cardTypes:   []CardType{CARDTYPE_PACT},
		symmetric:   true,
		canAttack:   true,
		endOnAttack: false,
		traits:      []int{TRAIT_SCIENTIFIC_COOPERATION},
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:  "International Tourism",
		age:         3,
		cardTypes:   []CardType{CARDTYPE_PACT},
		symmetric:   true,
		canAttack:   true,
		endOnAttack: false,
		traits:      []int{TRAIT_INTERNATIONAL_TOURISM},
		cardCounts:  []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:        "Loss of Sovereignty",
		age:               3,
		cardTypes:         []CardType{CARDTYPE_PACT},
		productionCulture: 2,
		symmetric:         false,
		canAttack:         false,
		endOnAttack:       false,
		traits:            []int{TRAIT_LOSS_OF_SOVEREIGNTY_A},
		bSide: &CardSchool{
			schoolName:        "Loss of Sovereignty - Part B",
			age:               3,
			cardTypes:         []CardType{CARDTYPE_PACT},
			productionCulture: -2,
			symmetric:         false,
			canAttack:         false,
			endOnAttack:       false,
			traits:            []int{TRAIT_LOSS_OF_SOVEREIGNTY_B},
		},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:      "Military Alliance",
		age:             3,
		cardTypes:       []CardType{CARDTYPE_PACT},
		productionPower: 3,
		symmetric:       true,
		canAttack:       true,
		endOnAttack:     true,
		traits:          []int{TRAIT_MILITARY_ALLIANCE},
		cardCounts:      []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:        "Peace Treaty",
		age:               3,
		cardTypes:         []CardType{CARDTYPE_PACT},
		productionCulture: 1,
		symmetric:         true,
		canAttack:         false,
		endOnAttack:       false,
		traits:            []int{TRAIT_PEACE_TREATY},
		cardCounts:        []int{1, 1, 1},
	})

	// Tactics
	schools = append(schools, &CardSchool{
		schoolName:            "Fighting Band",
		age:                   1,
		cardTypes:             []CardType{CARDTYPE_TACTIC},
		productionPower:       1,
		productionPowerLesser: 1,
		formation:             []int{0, 0},
		cardCounts:            []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:            "Heavy Cavalry",
		age:                   1,
		cardTypes:             []CardType{CARDTYPE_TACTIC},
		productionPower:       4,
		productionPowerLesser: 4,
		formation:             []int{1, 1, 1},
		cardCounts:            []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:            "Legion",
		age:                   1,
		cardTypes:             []CardType{CARDTYPE_TACTIC},
		productionPower:       2,
		productionPowerLesser: 2,
		formation:             []int{0, 0, 0},
		cardCounts:            []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:            "Medieval Army",
		age:                   1,
		cardTypes:             []CardType{CARDTYPE_TACTIC},
		productionPower:       2,
		productionPowerLesser: 2,
		formation:             []int{0, 1},
		cardCounts:            []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:            "Phalanx",
		age:                   1,
		cardTypes:             []CardType{CARDTYPE_TACTIC},
		productionPower:       3,
		productionPowerLesser: 3,
		formation:             []int{0, 0, 1},
		cardCounts:            []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:            "Classic Army",
		age:                   2,
		cardTypes:             []CardType{CARDTYPE_TACTIC},
		productionPower:       8,
		productionPowerLesser: 4,
		formation:             []int{0, 0, 1, 1},
		cardCounts:            []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:            "Conquistadors",
		age:                   2,
		cardTypes:             []CardType{CARDTYPE_TACTIC},
		productionPower:       5,
		productionPowerLesser: 3,
		formation:             []int{0, 1, 1},
		cardCounts:            []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:            "Defensive Army",
		age:                   2,
		cardTypes:             []CardType{CARDTYPE_TACTIC},
		productionPower:       6,
		productionPowerLesser: 3,
		formation:             []int{0, 0, 2},
		cardCounts:            []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:            "Fortifications",
		age:                   2,
		cardTypes:             []CardType{CARDTYPE_TACTIC},
		productionPower:       5,
		productionPowerLesser: 3,
		formation:             []int{2, 2},
		cardCounts:            []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:            "Mobile Army",
		age:                   2,
		cardTypes:             []CardType{CARDTYPE_TACTIC},
		productionPower:       5,
		productionPowerLesser: 3,
		formation:             []int{1, 2},
		cardCounts:            []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:            "Napoleonic Army",
		age:                   2,
		cardTypes:             []CardType{CARDTYPE_TACTIC},
		productionPower:       7,
		productionPowerLesser: 4,
		formation:             []int{0, 1, 2},
		cardCounts:            []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:            "Entrenchments",
		age:                   3,
		cardTypes:             []CardType{CARDTYPE_TACTIC},
		productionPower:       9,
		productionPowerLesser: 5,
		formation:             []int{0, 2, 2},
		cardCounts:            []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:            "Mechanized Army",
		age:                   3,
		cardTypes:             []CardType{CARDTYPE_TACTIC},
		productionPower:       10,
		productionPowerLesser: 5,
		formation:             []int{1, 2, 2},
		cardCounts:            []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:            "Modern Army",
		age:                   3,
		cardTypes:             []CardType{CARDTYPE_TACTIC},
		productionPower:       13,
		productionPowerLesser: 7,
		formation:             []int{0, 0, 1, 2},
		cardCounts:            []int{2, 2, 2},
	})
	schools = append(schools, &CardSchool{
		schoolName:            "Shock Troops",
		age:                   3,
		cardTypes:             []CardType{CARDTYPE_TACTIC},
		productionPower:       11,
		productionPowerLesser: 6,
		formation:             []int{0, 1, 1, 1},
		cardCounts:            []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Development of Agriculture",
		age:        0,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Development of Civilization",
		age:        0,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Development of Crafts",
		age:        0,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Development of Markets",
		age:        0,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Development of Politics",
		age:        0,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Development of Religion",
		age:        0,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Development of Science",
		age:        0,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Development of Settlement",
		age:        0,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Development of Trade Route",
		age:        0,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Development of Warfare",
		age:        0,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Barbarians",
		age:        1,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Border Conflict",
		age:        1,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Crusades",
		age:        1,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Cultural Influence",
		age:        1,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Foray",
		age:        1,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Good Harvest",
		age:        1,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Immigration",
		age:        1,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "New Deposits",
		age:        1,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Pestilence",
		age:        1,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Raiders",
		age:        1,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Rats",
		age:        1,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Rebellion",
		age:        1,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Reign of Terror",
		age:        1,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Scientific Breakthrough",
		age:        1,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Uncertain Borders",
		age:        1,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Civil Unrest",
		age:        2,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Cold War",
		age:        2,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Crime Wave",
		age:        2,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Economic Progress",
		age:        2,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Emigration",
		age:        2,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Iconoclasm",
		age:        2,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Independence Declaration",
		age:        2,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "International Agreement",
		age:        2,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "National Pride",
		age:        2,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Politics of Strength",
		age:        2,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Popularization of Science",
		age:        2,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Prosperity",
		age:        2,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Ravages of Time",
		age:        2,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Refugees",
		age:        2,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Terrorism",
		age:        2,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Impact of Agriculture",
		age:        3,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Impact of Architecture",
		age:        3,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Impact of Balance",
		age:        3,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Impact of Colonies",
		age:        3,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Impact of Competition",
		age:        3,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Impact of Government",
		age:        3,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Impact of Happiness",
		age:        3,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Impact of Industry",
		age:        3,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Impact of Population",
		age:        3,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Impact of Progress",
		age:        3,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Impact of Science",
		age:        3,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Impact of Strength",
		age:        3,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Impact of Technology",
		age:        3,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Impact of Variety",
		age:        3,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName: "Impact of Wonders",
		age:        3,
		cardTypes:  []CardType{CARDTYPE_EVENT},
		cardCounts: []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:            "Developed Territory I",
		age:                   1,
		cardTypes:             []CardType{CARDTYPE_TERRITORY},
		productionYellowToken: 1,
		productionBlueToken:   1,
		actionBonus:           3,
		traits:                []int{TRAIT_DEVELOPED_TERRITORY},
		cardCounts:            []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:          "Historic Territory I",
		age:                 1,
		cardTypes:           []CardType{CARDTYPE_TERRITORY},
		productionHappiness: 1,
		actionBonus:         6,
		traits:              []int{TRAIT_HISTORIC_TERRITORY},
		cardCounts:          []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:            "Inhabited Territory I",
		age:                   1,
		cardTypes:             []CardType{CARDTYPE_TERRITORY},
		productionYellowToken: 2,
		actionBonus:           1,
		traits:                []int{TRAIT_INHABITED_TERRITORY},
		cardCounts:            []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:      "Strategic Territory I",
		age:             1,
		cardTypes:       []CardType{CARDTYPE_TERRITORY},
		productionPower: 2,
		actionBonus:     3,
		traits:          []int{TRAIT_STRATEGIC_TERRITORY},
		cardCounts:      []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:            "Vast Territory I",
		age:                   1,
		cardTypes:             []CardType{CARDTYPE_TERRITORY},
		productionYellowToken: 3,
		productionBlueToken:   -1,
		actionBonus:           3,
		traits:                []int{TRAIT_VAST_TERRITORY},
		cardCounts:            []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:          "Wealthly Territory I",
		age:                 1,
		cardTypes:           []CardType{CARDTYPE_TERRITORY},
		productionBlueToken: 3,
		actionBonus:         5,
		traits:              []int{TRAIT_WEALTHLY_TERRITORY},
		cardCounts:          []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:            "Developed Territory II",
		age:                   2,
		cardTypes:             []CardType{CARDTYPE_TERRITORY},
		productionYellowToken: 2,
		productionBlueToken:   2,
		actionBonus:           5,
		traits:                []int{TRAIT_DEVELOPED_TERRITORY},
		cardCounts:            []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:          "Historic Territory II",
		age:                 2,
		cardTypes:           []CardType{CARDTYPE_TERRITORY},
		productionHappiness: 2,
		actionBonus:         11,
		traits:              []int{TRAIT_HISTORIC_TERRITORY},
		cardCounts:          []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:            "Inhabited Territory II",
		age:                   2,
		cardTypes:             []CardType{CARDTYPE_TERRITORY},
		productionYellowToken: 3,
		actionBonus:           2,
		traits:                []int{TRAIT_INHABITED_TERRITORY},
		cardCounts:            []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:      "Strategic Territory II",
		age:             2,
		cardTypes:       []CardType{CARDTYPE_TERRITORY},
		productionPower: 4,
		actionBonus:     5,
		traits:          []int{TRAIT_STRATEGIC_TERRITORY},
		cardCounts:      []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:            "Vast Territory II",
		age:                   2,
		cardTypes:             []CardType{CARDTYPE_TERRITORY},
		productionYellowToken: 4,
		productionBlueToken:   -1,
		actionBonus:           4,
		traits:                []int{TRAIT_VAST_TERRITORY},
		cardCounts:            []int{1, 1, 1},
	})
	schools = append(schools, &CardSchool{
		schoolName:          "Wealthly Territory II",
		age:                 2,
		cardTypes:           []CardType{CARDTYPE_TERRITORY},
		productionBlueToken: 4,
		actionBonus:         9,
		traits:              []int{TRAIT_WEALTHLY_TERRITORY},
		cardCounts:          []int{1, 1, 1},
	})

	result = make(map[int]*CardSchool)
	for i, s := range schools {
		s.schoolId = i + 1
		result[s.schoolId] = s
	}
	return
}
