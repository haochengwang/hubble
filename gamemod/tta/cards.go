package main

type CardType int

const (
	CARDTYPE_LEADER                    = 1
	CARDTYPE_WONDER                    = 2
	CARDTYPE_ACTION                    = 3
	CARDTYPE_TECH                      = 4
	CARDTYPE_TECH_SPECIAL              = 5
	CARDTYPE_TECH_SPECIAL_CIVIL        = 6
	CARDTYPE_TECH_SPECIAL_MILITARY     = 7
	CARDTYPE_TECH_SPECIAL_COLONIZE     = 8
	CARDTYPE_TECH_SPECIAL_CONSTRUCTION = 9
	CARDTYPE_TECH_FARM                 = 10
	CARDTYPE_TECH_MINE                 = 11
	CARDTYPE_TECH_URBAN                = 12
	CARDTYPE_TECH_URBAN_TEMPLE         = 13
	CARDTYPE_TECH_URBAN_LAB            = 14
	CARDTYPE_TECH_URBAN_ARENA          = 15
	CARDTYPE_TECH_URBAN_THEATER        = 16
	CARDTYPE_TECH_URBAN_LIBRARY        = 17
	CARDTYPE_TECH_MILLI                = 18
	CARDTYPE_TECH_MILLI_INFANTRY       = 19
	CARDTYPE_TECH_MILLI_CAVALRY        = 20
	CARDTYPE_TECH_MILLI_ARTILLERY      = 21
	CARDTYPE_TECH_MILLI_AIRFORCE       = 22
	CARDTYPE_TECH_GOVERNMENT           = 23

	CARDTYPE_DEFCOL     = 101
	CARDTYPE_TACTIC     = 102
	CARDTYPE_COLONY     = 103
	CARDTYPE_EVENT      = 104
	CARDTYPE_PACT       = 105
	CARDTYPE_AGGRESSION = 106
	CARDTYPE_WAR        = 107
)

type CardSchool struct {
	schoolId    int
	schoolName  string
	shortName   string
	description string

	age      int
	cardType []CardType

	tech             int
	techRevolution   int // Only for governments
	buildCost        int
	wonderBuildCosts []int // Only for wonders
	actionBonus      int   // Only for actions

	productionCrop          int
	productionResource      int
	productionCulture       int
	productionTech          int
	productionHappiness     int
	productionPower         int
	productionColonizePower int
	productionUrbanLimit    int
	productionWhiteToken    int // White tokens
	productionRedToken      int // Red tokens
	productionYellowToken   int // Yellow tokens
	productionBlueToken     int // Blue tokens

	cardCounts []int
}

type Card struct {
	id       int
	schoolId int
}

type CardStack []Card

type CardPosition struct {
	stackId  int
	position int
}

type AddCardRequest struct {
	schoolId int
	stackId  int
}

type BanishCardRequest struct {
	position CardPosition
}

type MoveCardRequest struct {
	sourcePosition CardPosition
	targetPosition CardPosition
}

type ShuffleStackRequest struct {
	stackId int
}

type CardStackUniversalManager struct {
	nextStackId   int
	nextCardId    int
	cardStacks    map[int]CardStack
	cardPositions map[int]CardPosition
}

func NewCardStackUniversalManager() *CardStackUniversalManager {
	return &CardStackUniversalManager{
		nextStackId:   1,
		nextCardId:    1,
		cardStacks:    make(map[int]CardStack),
		cardPositions: make(map[int]CardPosition),
	}
}

func (m *CardStackUniversalManager) newStack() (result int) {
	result = m.nextStackId
	stack := CardStack(make([]Card, 0))
	m.cardStacks[m.nextStackId] = stack
	m.nextStackId++

	return result
}

func (m *CardStackUniversalManager) newCard(schoolId int) (result Card) {
	id := m.nextCardId
	result = Card{
		id:       id,
		schoolId: schoolId,
	}
	m.nextStackId++

	return
}

func (m *CardStackUniversalManager) processRequest(request interface{}) {
	switch request := request.(type) {
	case *AddCardRequest:
		stack := m.cardStacks[request.stackId]
		newStack := append([]Card(stack), m.newCard(request.schoolId))
		m.cardStacks[request.stackId] = newStack
	case *BanishCardRequest:
		//stack := m.cardStacks[request.position.stackId]
	case *MoveCardRequest:
	case *ShuffleStackRequest:
	}
}
