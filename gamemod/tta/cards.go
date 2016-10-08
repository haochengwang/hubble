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
	CARDTYPE_TECH_MILI                 = 18
	CARDTYPE_TECH_MILI_INFANTRY        = 19
	CARDTYPE_TECH_MILI_CAVALRY         = 20
	CARDTYPE_TECH_MILI_ARTILLERY       = 21
	CARDTYPE_TECH_MILI_AIRFORCE        = 22
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
	CARDTYPE_TERRITORY  = 103
	CARDTYPE_EVENT      = 104
	CARDTYPE_PACT       = 105
	CARDTYPE_AGGRESSION = 106
	CARDTYPE_WAR        = 107
)

const (
	TRAIT_LIB_OF_ALEXANDRIA int = iota
	TRAIT_GREAT_WALL
	TRAIT_ST_PETERS_BASILICA
	TRAIT_TAJ_MAHAL
	TRAIT_TRANSCONT_RR
	TRAIT_OCEAN_LINER_SERVICE
	TRAIT_HOLLYWOOD
	TRAIT_INTERNET
	TRAIT_FIRST_SPACE_FLIGHT
	TRAIT_FAST_FOOD_CHAINS
	TRAIT_LIB_LAB_AMPLIFY // Da Vinci, Newton and Einstein
	TRAIT_JULIUS_CAESAR
	TRAIT_HOMER
	TRAIT_MOSES
	TRAIT_HAMMURABI
	TRAIT_ARISTOTLE
	TRAIT_ALEXANDER_THE_GREAT
	TRAIT_MICHELANGELO
	TRAIT_JOAN_OF_ARC
	TRAIT_LEONARDO_DA_VINCI
	TRAIT_GENGHIS_KHAN
	TRAIT_CHRISTOPHER_COLUMBUS
	TRAIT_FREDERICK_BARBAROSSA
	TRAIT_WILLIAM_SHAKESPEARE
	TRAIT_JAMES_COOK
	TRAIT_NAPOLEON_BONAPARTE
	TRAIT_MAXIMILLIEN_ROBESPIERRE
	TRAIT_J_S_BACH
	TRAIT_ISAAC_NEWTON
	TRAIT_ALBERT_EINSTEIN
	TRAIT_MAHATMA_GANDHI
	TRAIT_CHARLIE_CHAPLIN
	TRAIT_BILL_GATES
	TRAIT_WINSTON_CHURCHILL
	TRAIT_SID_MEIER

	// Aggressions
	TRAIT_ENSLAVE
	TRAIT_PLUNDER
	TRAIT_RAID
	TRAIT_ANNEX
	TRAIT_INFILTRATE
	TRAIT_SPY
	TRAIT_ARMED_INTERVENTION

	// Pacts
	TRAIT_OPEN_BORDER_AGGREMENT
	TRAIT_TRADE_ROUTE_AGGREMENT_A
	TRAIT_TRADE_ROUTE_AGGREMENT_B
	TRAIT_ACCEPTANCE_OF_SUPREMACY_A
	TRAIT_ACCEPTANCE_OF_SUPREMACY_B
	TRAIT_INTERNATIONAL_TRADE_AGGREMENT_A
	TRAIT_INTERNATIONAL_TRADE_AGGREMENT_B
	TRAIT_PROMISE_OF_MILITARY_PROTECTION_A
	TRAIT_PROMISE_OF_MILITARY_PROTECTION_B
	TRAIT_SCIENTIFIC_COOPERATION
	TRAIT_INTERNATIONAL_TOURISM
	TRAIT_LOSS_OF_SOVEREIGNTY_A
	TRAIT_LOSS_OF_SOVEREIGNTY_B
	TRAIT_MILITARY_ALLIANCE
	TRAIT_PEACE_TREATY

	// Territories
	TRAIT_DEVELOPED_TERRITORY
	TRAIT_HISTORIC_TERRITORY
	TRAIT_INHABITED_TERRITORY
	TRAIT_STRATEGIC_TERRITORY
	TRAIT_VAST_TERRITORY
	TRAIT_WEALTHLY_TERRITORY
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
	traits           []int

	productionCrop          int
	productionResource      int
	productionCulture       int
	productionTech          int
	productionHappiness     int
	productionPower         int
	productionPowerLesser   int // For tactics only
	productionColonizePower int
	productionUrbanLimit    int
	productionWhiteToken    int // White tokens
	productionRedToken      int // Red tokens
	productionYellowToken   int // Yellow tokens
	productionBlueToken     int // Blue tokens

	miliActionCost int   // For aggressions and wars
	formation      []int // For tactics only
	symmetric      bool  // For pacts
	canAttack      bool  // For pacts
	endOnAttack    bool  // For pacts

	bSide *CardSchool // For pacts, the virtual card of b side pact

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

func (s *CardSchool) hasTrait(trait int) bool {
	for _, t := range s.traits {
		if t == trait {
			return true
		}
	}
	return false
}

func (s *CardSchool) isCivilCard() bool {
	return s.hasType(CARDTYPE_LEADER) ||
		s.hasType(CARDTYPE_WONDER) ||
		s.hasType(CARDTYPE_TECH) ||
		s.hasType(CARDTYPE_ACTION)
}

type Card struct {
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
}
