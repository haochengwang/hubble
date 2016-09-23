package main

type CardType int

const (
	CARDTYPE_LEADER                    = 1
	CARDTYPE_WONDER                    = 2
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

	CARDTYPE_ACTION                     = 24
	CARDTYPE_ACTION_BREAKTHROUGH        = 25
	CARDTYPE_ACTION_CULTURAL_HERITAGE   = 26
	CARDTYPE_ACTION_EFFICIENT_UPGRADE   = 27
	CARDTYPE_ACTION_ENDOWMENT_FOR_ARTS  = 28
	CARDTYPE_ACTION_ENGINEERING_GENIUS  = 29
	CARDTYPE_ACTION_FRUGALITY           = 30
	CARDTYPE_ACTION_MILITARY_BUILD_UP   = 31
	CARDTYPE_ACTION_PATRIOTISM          = 32
	CARDTYPE_ACTION_RESERVES            = 33
	CARDTYPE_ACTION_REVOLUTIONARY_IDEA  = 34
	CARDTYPE_ACTION_RICH_LAND           = 35
	CARDTYPE_ACTION_STOCKPILE           = 36
	CARDTYPE_ACTION_URBAN_GROWTH        = 37
	CARDTYPE_ACTION_WAVE_OF_NATIONALISM = 38

	CARDTYPE_DEFCOL     = 101
	CARDTYPE_TACTIC     = 102
	CARDTYPE_COLONY     = 103
	CARDTYPE_EVENT      = 104
	CARDTYPE_PACT       = 105
	CARDTYPE_AGGRESSION = 106
	CARDTYPE_WAR        = 107
)

const (
	SA_LIB_OF_ALEXANDRIA int = iota
	SA_GREAT_WALL
	SA_ST_PETERS_BASILICA
	SA_TAJ_MAHAL
	SA_TRANSCONT_RR
	SA_OCEAN_LINER_SERVICE
	SA_HOLLYWOOD
	SA_INTERNET
	SA_FIRST_SPACE_FLIGHT
	SA_FAST_FOOD_CHAINS
	SA_LIB_LAB_AMPLIFY // Da Vinci, Newton and Einstein
	SA_JULIUS_CAESAR
	SA_HOMER
	SA_MOSES
	SA_HAMMURABI
	SA_ARISTOTLE
	SA_ALEXANDER_THE_GREAT
	SA_MICHELANGELO
	SA_JOAN_OF_ARC
	SA_LEONARDO_DA_VINCI
	SA_GENGHIS_KHAN
	SA_CHRISTOPHER_COLUMBUS
	SA_FREDERICK_BARBAROSSA
	SA_WILLIAM_SHAKESPEARE
	SA_JAMES_COOK
	SA_NAPOLEON_BONAPARTE
	SA_MAXIMILLIEN_ROBESPIERRE
	SA_J_S_BACH
	SA_ISAAC_NEWTON
	SA_ALBERT_EINSTEIN
	SA_MAHATMA_GANDHI
	SA_CHARLIE_CHAPLIN
	SA_BILL_GATES
	SA_WINSTON_CHURCHILL
	SA_SID_MEIER
)

type CardSchool struct {
	schoolId    int
	schoolName  string
	shortName   string
	description string

	age       int
	cardTypes []CardType

	tech             int
	techRevolution   int // Only for governments
	buildCost        int
	wonderBuildCosts []int // Only for wonders
	actionBonus      int   // Only for actions
	specialAbilities []int

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

func (s *CardSchool) hasType(cardType CardType) bool {
	for _, t := range s.cardTypes {
		if t == cardType {
			return true
		}
	}
	return false
}

func (s *CardSchool) hasSpecialAbility(specialAbility int) bool {
	for _, sa := range s.specialAbilities {
		if sa == specialAbility {
			return true
		}
	}
	return false
}

/*type Card struct {
	id       int
	schoolId int
}

type CardStack []Card

func pushCard(stack CardStack, card Card, position int) CardStack {
	if position == len(stack) {
		return append(stack, card)
	}
	rear := append([]Card{}, stack[position:]...)
	result := append(stack[:position], card)
	result = append(result, rear...)
	return result
}

func popCard(stack CardStack, position int) (CardStack, Card) {
	card := stack[position]
	if position == len(stack) {
		return stack[:position], card
	}
	return append(stack[:position], stack[position+1:]...), card
}

type CardPosition struct {
	stackId  int
	position int
}

type AddCardRequest struct {
	schoolId int
	position CardPosition
}

type AddCardToTopRequest struct {
	schoolId int
	stackId  int
}

type BanishCardRequest struct {
	position CardPosition
}

type BanishAllCardsInStackRequest struct {
	stackId int
}

type MoveCardRequest struct {
	sourcePosition CardPosition
	targetPosition CardPosition
}

type SwapCardRequest struct {
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
	m.nextCardId++

	return
}

func (m *CardStackUniversalManager) processRequest(request interface{}) {
	switch request := request.(type) {
	case *AddCardRequest:
		pos := request.position
		m.cardStacks[pos.stackId] = pushCard(
			m.cardStacks[pos.stackId], m.newCard(request.schoolId), pos.position)
	case *AddCardToTopRequest:
		m.processRequest(&AddCardRequest{
			schoolId: request.schoolId,
			position: CardPosition{
				stackId:  request.stackId,
				position: m.getStackSize(request.stackId),
			},
		})
	case *BanishCardRequest:
		pos := request.position
		m.cardStacks[pos.stackId], _ = popCard(
			m.cardStacks[pos.stackId], pos.position)
	case *BanishAllCardsInStackRequest:
		stackId := request.stackId
		m.cardStacks[stackId] = []Card{}
	case *MoveCardRequest:
		pos := request.sourcePosition
		var card Card
		m.cardStacks[pos.stackId], card = popCard(
			m.cardStacks[pos.stackId], pos.position)

		pos = request.targetPosition
		m.cardStacks[pos.stackId] = pushCard(
			m.cardStacks[pos.stackId], card, pos.position)
	case *SwapCardRequest:
		pos1 := request.sourcePosition
		pos2 := request.targetPosition
		card1 := m.cardStacks[pos1.stackId][pos1.position]
		card2 := m.cardStacks[pos2.stackId][pos2.position]
		m.cardStacks[pos2.stackId][pos2.position] = card1
		m.cardStacks[pos1.stackId][pos1.position] = card2
	}
}

func (m *CardStackUniversalManager) getFirstCard(stackId int) *Card {
	if stack, ok := m.cardStacks[stackId]; ok {
		if len(stack) <= 0 {
			return nil
		}
		return &stack[0]
	}
	return nil
}

func (m *CardStackUniversalManager) getCardAt(stackId int, position int) *Card {
	if stack, ok := m.cardStacks[stackId]; ok {
		return &stack[position]
	}
	return nil
}

func (m *CardStackUniversalManager) getStackSize(stackId int) int {
	if stack, ok := m.cardStacks[stackId]; ok {
		return len(stack)
	}
	return -1
}*/
