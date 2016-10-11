package sw_duel

const (
	RESOURCE_GOLD int = iota
	RESOURCE_WOOD
	RESOURCE_CLAY
	RESOURCE_STONE
	RESOURCE_CLOTH
	RESOURCE_GLASS
	RESOURCE_LIMIT
)

const (
	// Basic resource
	PRODUCE_GOLD int = iota
	PRODUCE_WOOD
	PRODUCE_CLAY
	PRODUCE_STONE
	PRODUCE_CLOTH
	PRODUCE_GLASS
	PRODUCE_SHIELD
	PRODUCE_SCORE

	// Alternative goods
	PRODUCE_RAW_GOODS
	PRODUCE_MANUFACTURED_GOODS

	// Science tokens
	PRODUCE_SCIENCE_0 // Pendulum
	PRODUCE_SCIENCE_1 // Wheels
	PRODUCE_SCIENCE_2
	PRODUCE_SCIENCE_3
	PRODUCE_SCIENCE_4
	PRODUCE_SCIENCE_5
	PRODUCE_SCIENCE_6

	// Cheap resources
	PRODUCE_CHEAP_WOOD
	PRODUCE_CHEAP_CLAY
	PRODUCE_CHEAP_STONE
	PRODUCE_CHEAP_CLOTH
	PRODUCE_CHEAP_GLASS

	// Yellow card effects
	PRODUCE_GRAY_INCOME
	PRODUCE_BROWN_INCOME
	PRODUCE_RED_INCOME
	PRODUCE_YELLOW_INCOME
	PRODUCE_WONDER_INCOME

	// Guide effects
	PRODUCE_TRAIT_MERCHANTS_GUILD
	PRODUCE_TRAIT_SHIPOWNERS_GUILD
	PRODUCE_TRAIT_BUILDERS_GUILD
	PRODUCE_TRAIT_MAGISTRATES_GUILD
	PRODUCE_TRAIT_SCIENTIST_GUILD
	PRODUCE_TRAIT_MONEYLENDERS_GUILD
	PRODUCE_TRAIT_TACTICIANS_GUILD

	// Special wonder effects and progress tokens
	PRODUCE_TRAIT_ADDITIONAL_ACTION
	PRODUCE_TRAIT_BURN_BROWN
	PRODUCE_TRAIT_BURN_GRAY
	PRODUCE_TRAIT_BURN_GOLD
	PRODUCE_TRAIT_MAUSOLEUM
	PRODUCE_TRAIT_AGRICULTURE
	PRODUCE_TRAIT_ARCHITECTURE
	PRODUCE_TRAIT_ECONOMY
	PRODUCE_TRAIT_LAW
	PRODUCE_TRAIT_MASONRY
	PRODUCE_TRAIT_MATHEMATICS
	PRODUCE_TRAIT_PHILOSOPHY
	PRODUCE_TRAIT_STRATEGY
	PRODUCE_TRAIT_THEOLOGY
	PRODUCE_TRAIT_URBANISM

	PRODUCE_LIMIT
)

const (
	COLOR_BROWN = iota
	COLOR_GRAY
	COLOR_YELLOW
	COLOR_RED
	COLOR_GREEN
	COLOR_BLUE
	COLOR_PURPLE
)

type BuildPrerequisite struct {
	Resources []int
	LinkId    int
}

type Product struct {
	Product []int
}

type Buildable struct {
	Id   int
	Name string

	Prereq   BuildPrerequisite
	Produces []int
}

type Card struct {
	Buildable
	Color int
}

func generateProductMetrics(gold, score int, additional []int) (result []int) {
	result = make([]int, PRODUCE_LIMIT)
	result[PRODUCE_GOLD] = gold
	result[PRODUCE_SCORE] = score
	for _, product := range additional {
		result[product]++
	}
	return
}

func GetAllCards() map[int]*Card {
	return map[int]*Card{
		// Age 1
		1: &Card{Buildable: Buildable{
			Id:   1,
			Name: "Lumber Yard",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_WOOD}),
		}},

		2: &Card{Buildable: Buildable{
			Id:   2,
			Name: "Logging Camp",
			Prereq: BuildPrerequisite{
				Resources: []int{1, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_WOOD}),
		}},

		3: &Card{Buildable: Buildable{
			Id:   3,
			Name: "Clay Pool",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_CLAY}),
		}},

		4: &Card{Buildable: Buildable{
			Id:   4,
			Name: "Clay Pit",
			Prereq: BuildPrerequisite{
				Resources: []int{1, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_CLAY}),
		}},

		5: &Card{Buildable: Buildable{
			Id:   5,
			Name: "Quarry",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_STONE}),
		}},

		6: &Card{Buildable: Buildable{
			Id:   6,
			Name: "Stone Pit",
			Prereq: BuildPrerequisite{
				Resources: []int{1, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_STONE}),
		}},

		7: &Card{Buildable: Buildable{
			Id:   7,
			Name: "Glassworks",
			Prereq: BuildPrerequisite{
				Resources: []int{1, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_GLASS}),
		}},

		8: &Card{Buildable: Buildable{
			Id:   8,
			Name: "Press",
			Prereq: BuildPrerequisite{
				Resources: []int{1, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_CLOTH}),
		}},

		9: &Card{Buildable: Buildable{
			Id:   9,
			Name: "Guard Tower",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_SHIELD}),
		}},

		10: &Card{Buildable: Buildable{
			Id:   10,
			Name: "Workshop",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 0, 0, 1, 0},
			},
			Produces: generateProductMetrics(0, 1, []int{PRODUCE_SCIENCE_0}),
		}},

		11: &Card{Buildable: Buildable{
			Id:   11,
			Name: "Apothecary",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 0, 0, 0, 1},
			},
			Produces: generateProductMetrics(0, 1, []int{PRODUCE_SCIENCE_1}),
		}},

		12: &Card{Buildable: Buildable{
			Id:   12,
			Name: "Stone Reserve",
			Prereq: BuildPrerequisite{
				Resources: []int{3, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_CHEAP_STONE}),
		}},

		13: &Card{Buildable: Buildable{
			Id:   13,
			Name: "Clay Reserve",
			Prereq: BuildPrerequisite{
				Resources: []int{3, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_CHEAP_CLAY}),
		}},

		14: &Card{Buildable: Buildable{
			Id:   14,
			Name: "Wood Reserve",
			Prereq: BuildPrerequisite{
				Resources: []int{3, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_CHEAP_WOOD}),
		}},

		15: &Card{Buildable: Buildable{
			Id:   15,
			Name: "Stable",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 1, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_SHIELD}),
		}},

		16: &Card{Buildable: Buildable{
			Id:   16,
			Name: "Garrison",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 1, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_SHIELD}),
		}},

		17: &Card{Buildable: Buildable{
			Id:   17,
			Name: "Palisade",
			Prereq: BuildPrerequisite{
				Resources: []int{2, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_SHIELD}),
		}},

		18: &Card{Buildable: Buildable{
			Id:   18,
			Name: "Scriptorium",
			Prereq: BuildPrerequisite{
				Resources: []int{2, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_SCIENCE_2}),
		}},

		19: &Card{Buildable: Buildable{
			Id:   19,
			Name: "Pharmacist",
			Prereq: BuildPrerequisite{
				Resources: []int{2, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_SCIENCE_3}),
		}},

		20: &Card{Buildable: Buildable{
			Id:   20,
			Name: "Theater",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 3, []int{}),
		}},

		21: &Card{Buildable: Buildable{
			Id:   21,
			Name: "Altar",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 3, []int{}),
		}},

		22: &Card{Buildable: Buildable{
			Id:   22,
			Name: "Baths",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 0, 1, 0, 0},
			},
			Produces: generateProductMetrics(0, 3, []int{}),
		}},

		23: &Card{Buildable: Buildable{
			Id:   23,
			Name: "Tavern",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(4, 0, []int{}),
		}},

		// Age 2
		24: &Card{Buildable: Buildable{
			Id:   24,
			Name: "Sawmill",
			Prereq: BuildPrerequisite{
				Resources: []int{2, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_WOOD, PRODUCE_WOOD}),
		}},

		25: &Card{Buildable: Buildable{
			Id:   25,
			Name: "Brickyard",
			Prereq: BuildPrerequisite{
				Resources: []int{2, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_CLAY, PRODUCE_CLAY}),
		}},

		26: &Card{Buildable: Buildable{
			Id:   26,
			Name: "Shelf Quarry",
			Prereq: BuildPrerequisite{
				Resources: []int{2, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_STONE, PRODUCE_STONE}),
		}},

		27: &Card{Buildable: Buildable{
			Id:   27,
			Name: "Glass Blower",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_GLASS}),
		}},

		28: &Card{Buildable: Buildable{
			Id:   28,
			Name: "Drying Room",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_CLOTH}),
		}},

		29: &Card{Buildable: Buildable{
			Id:   29,
			Name: "Walls",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 0, 2, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_SHIELD, PRODUCE_SHIELD}),
		}},

		30: &Card{Buildable: Buildable{
			Id:   30,
			Name: "Forum",
			Prereq: BuildPrerequisite{
				Resources: []int{3, 0, 1, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_MANUFACTURED_GOODS}),
		}},

		31: &Card{Buildable: Buildable{
			Id:   31,
			Name: "Caravansery",
			Prereq: BuildPrerequisite{
				Resources: []int{2, 0, 0, 0, 1, 1},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_RAW_GOODS}),
		}},

		32: &Card{Buildable: Buildable{
			Id:   32,
			Name: "Customs House",
			Prereq: BuildPrerequisite{
				Resources: []int{4, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_CHEAP_GLASS, PRODUCE_CHEAP_CLOTH}),
		}},

		33: &Card{Buildable: Buildable{
			Id:   33,
			Name: "Tribunal",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 2, 0, 0, 0, 1},
			},
			Produces: generateProductMetrics(0, 5, []int{}),
		}},

		34: &Card{Buildable: Buildable{
			Id:   34,
			Name: "House Breeders",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 1, 1, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_SHIELD}),
		}},

		35: &Card{Buildable: Buildable{
			Id:   35,
			Name: "Barracks",
			Prereq: BuildPrerequisite{
				Resources: []int{4, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_SHIELD}),
		}},

		36: &Card{Buildable: Buildable{
			Id:   36,
			Name: "Archery Range",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 1, 0, 1, 1, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_SHIELD, PRODUCE_SHIELD}),
		}},

		37: &Card{Buildable: Buildable{
			Id:   37,
			Name: "Parade Ground",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 2, 0, 0, 1},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_SHIELD, PRODUCE_SHIELD}),
		}},

		38: &Card{Buildable: Buildable{
			Id:   38,
			Name: "Library",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 1, 0, 1, 0, 1},
			},
			Produces: generateProductMetrics(0, 2, []int{PRODUCE_SCIENCE_2}),
		}},

		39: &Card{Buildable: Buildable{
			Id:   39,
			Name: "Dispensary",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 2, 1, 0, 0},
			},
			Produces: generateProductMetrics(0, 2, []int{PRODUCE_SCIENCE_3}),
		}},

		40: &Card{Buildable: Buildable{
			Id:   40,
			Name: "School",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 1, 0, 0, 2, 0},
			},
			Produces: generateProductMetrics(0, 1, []int{PRODUCE_SCIENCE_1}),
		}},

		41: &Card{Buildable: Buildable{
			Id:   41,
			Name: "Laboratory",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 1, 0, 0, 0, 2},
			},
			Produces: generateProductMetrics(0, 1, []int{PRODUCE_SCIENCE_0}),
		}},

		42: &Card{Buildable: Buildable{
			Id:   42,
			Name: "Statue",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 2, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 4, []int{}),
		}},

		43: &Card{Buildable: Buildable{
			Id:   43,
			Name: "Statue",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 1, 0, 0, 1, 0},
			},
			Produces: generateProductMetrics(0, 4, []int{}),
		}},

		44: &Card{Buildable: Buildable{
			Id:   44,
			Name: "Aqueduct",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 0, 3, 0, 0},
			},
			Produces: generateProductMetrics(0, 5, []int{}),
		}},

		45: &Card{Buildable: Buildable{
			Id:   45,
			Name: "Rostrum",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 1, 0, 1, 0, 0},
			},
			Produces: generateProductMetrics(0, 4, []int{}),
		}},

		46: &Card{Buildable: Buildable{
			Id:   46,
			Name: "Brewery",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(6, 0, []int{}),
		}},

		// Age 3
		47: &Card{Buildable: Buildable{
			Id:   47,
			Name: "Arsenal",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 2, 3, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_SHIELD, PRODUCE_SHIELD, PRODUCE_SHIELD}),
		}},

		48: &Card{Buildable: Buildable{
			Id:   48,
			Name: "Courthouse",
			Prereq: BuildPrerequisite{
				Resources: []int{8, 0, 0, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_SHIELD, PRODUCE_SHIELD, PRODUCE_SHIELD}),
		}},

		49: &Card{Buildable: Buildable{
			Id:   49,
			Name: "Academy",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 1, 0, 1, 0, 2},
			},
			Produces: generateProductMetrics(0, 3, []int{PRODUCE_SCIENCE_4}),
		}},

		50: &Card{Buildable: Buildable{
			Id:   50,
			Name: "Study",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 2, 0, 0, 1, 1},
			},
			Produces: generateProductMetrics(0, 3, []int{PRODUCE_SCIENCE_4}),
		}},

		51: &Card{Buildable: Buildable{
			Id:   51,
			Name: "Chamber of Commerce",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 0, 0, 2, 0},
			},
			Produces: generateProductMetrics(0, 3, []int{PRODUCE_GRAY_INCOME}),
		}},

		52: &Card{Buildable: Buildable{
			Id:   52,
			Name: "Port",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 1, 0, 0, 1, 1},
			},
			Produces: generateProductMetrics(0, 3, []int{PRODUCE_BROWN_INCOME}),
		}},

		53: &Card{Buildable: Buildable{
			Id:   53,
			Name: "Armory",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 0, 2, 0, 1},
			},
			Produces: generateProductMetrics(0, 3, []int{PRODUCE_RED_INCOME}),
		}},

		54: &Card{Buildable: Buildable{
			Id:   54,
			Name: "Palace",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 1, 1, 1, 0, 2},
			},
			Produces: generateProductMetrics(0, 7, []int{}),
		}},

		55: &Card{Buildable: Buildable{
			Id:   55,
			Name: "Town Hall",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 2, 0, 3, 0, 0},
			},
			Produces: generateProductMetrics(0, 7, []int{}),
		}},

		56: &Card{Buildable: Buildable{
			Id:   56,
			Name: "Obelisk",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 0, 2, 0, 1},
			},
			Produces: generateProductMetrics(0, 5, []int{}),
		}},

		57: &Card{Buildable: Buildable{
			Id:   57,
			Name: "Fortifications",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 1, 2, 1, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_SHIELD, PRODUCE_SHIELD}),
		}},

		58: &Card{Buildable: Buildable{
			Id:   58,
			Name: "Siege Workshop",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 3, 0, 0, 0, 1},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_SHIELD, PRODUCE_SHIELD}),
		}},

		59: &Card{Buildable: Buildable{
			Id:   59,
			Name: "Circus",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 2, 2, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_SHIELD, PRODUCE_SHIELD}),
		}},

		60: &Card{Buildable: Buildable{
			Id:   60,
			Name: "University",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 1, 0, 1, 1},
			},
			Produces: generateProductMetrics(0, 2, []int{PRODUCE_SCIENCE_5}),
		}},

		61: &Card{Buildable: Buildable{
			Id:   61,
			Name: "Observatory",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 0, 1, 2, 0},
			},
			Produces: generateProductMetrics(0, 2, []int{PRODUCE_SCIENCE_5}),
		}},

		62: &Card{Buildable: Buildable{
			Id:   62,
			Name: "Gardens",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 2, 2, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 6, []int{}),
		}},

		63: &Card{Buildable: Buildable{
			Id:   63,
			Name: "Pantheon",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 1, 1, 0, 2, 0},
			},
			Produces: generateProductMetrics(0, 6, []int{}),
		}},

		64: &Card{Buildable: Buildable{
			Id:   64,
			Name: "Lighthouse",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 2, 0, 0, 0, 1},
			},
			Produces: generateProductMetrics(0, 3, []int{PRODUCE_YELLOW_INCOME}),
		}},

		65: &Card{Buildable: Buildable{
			Id:   65,
			Name: "Arena",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 1, 1, 1, 0, 0},
			},
			Produces: generateProductMetrics(0, 3, []int{PRODUCE_WONDER_INCOME}),
		}},

		// Guilds
		100: &Card{Buildable: Buildable{
			Id:   100,
			Name: "Merchants Guild",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 1, 1, 0, 1, 1},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_TRAIT_MERCHANTS_GUILD}),
		}},

		101: &Card{Buildable: Buildable{
			Id:   101,
			Name: "Shipowners Guild",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 1, 1, 1, 1},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_TRAIT_SHIPOWNERS_GUILD}),
		}},

		102: &Card{Buildable: Buildable{
			Id:   102,
			Name: "Builders Guild",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 1, 1, 2, 0, 1},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_TRAIT_BUILDERS_GUILD}),
		}},

		103: &Card{Buildable: Buildable{
			Id:   103,
			Name: "Magistrates Guild",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 2, 1, 0, 1, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_TRAIT_MAGISTRATES_GUILD}),
		}},

		104: &Card{Buildable: Buildable{
			Id:   104,
			Name: "Scientists Guild",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 2, 2, 0, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_TRAIT_SCIENTIST_GUILD}),
		}},

		105: &Card{Buildable: Buildable{
			Id:   105,
			Name: "Moneylenders Guild",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 2, 0, 2, 0, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_TRAIT_MONEYLENDERS_GUILD}),
		}},

		106: &Card{Buildable: Buildable{
			Id:   106,
			Name: "Tacticians Guild",
			Prereq: BuildPrerequisite{
				Resources: []int{0, 0, 1, 2, 1, 0},
			},
			Produces: generateProductMetrics(0, 0, []int{PRODUCE_TRAIT_TACTICIANS_GUILD}),
		}},
	}
}
