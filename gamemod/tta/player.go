package main

import (
	"fmt"
)

const (
	MILI_INFANTRY int = iota
	MILI_CAVALRY
	MILI_ARTILERY
	MILI_AIRFORCE
	FARM
	MINE
	URBAN_TEMPLE
	URBAN_LAB
	URBAN_ARENA
	URBAN_LIBRARY
	URBAN_THEATER
	GOVERNMENT
	LEADER
	WONDER_NOT_COMPLETED
	WONDER_COMPLETED
	TECH_SPECIAL_CIVIL
	TECH_SPECIAL_WARFARE
	TECH_SPECIAL_COLONIZE
	TECH_SPECIAL_CONSTRUCTION
	HAND
	USER_STACK_SIZE
)

const (
	FREE_YELLOW = iota
	FREE_BLUE
	FREE_WORKER
	WONDER
	POLITICAL_UNUSED
	POLITICAL_USED

	CULTURE_COUNTER
	TECH_COUNTER
	LEADER_A_TAKEN
	LEADER_I_TAKEN
	LEADER_II_TAKEN
	LEADER_III_TAKEN
)

const (
	TOKEN_DEFAULT = iota
	TOKEN_YELLOW
	TOKEN_BLUE
	TOKEN_WHITE
	TOKEN_RED
)

type PlayerBoard struct {
	game *TtaGame

	stacks              []int
	techTokenManager    *TokenBankUniversalManager
	specialTokenManager *TokenBankUniversalManager
}

func initPlayerBoard(game *TtaGame) (result *PlayerBoard) {
	csm := game.cardStackManager
	stacks := make([]int, USER_STACK_SIZE)

	// Prepare stacks
	for i := 0; i < USER_STACK_SIZE; i++ {
		stacks[i] = csm.newStack()
	}

	// Prepare initial cards
	csm.processRequest(&AddCardRequest{
		position: CardPosition{
			stackId:  stacks[MILI_INFANTRY],
			position: 0,
		},
		schoolId: 25,
	})
	csm.processRequest(&AddCardRequest{
		position: CardPosition{
			stackId:  stacks[FARM],
			position: 0,
		},
		schoolId: 1,
	})
	csm.processRequest(&AddCardRequest{
		position: CardPosition{
			stackId:  stacks[MINE],
			position: 0,
		},
		schoolId: 5,
	})
	csm.processRequest(&AddCardRequest{
		position: CardPosition{
			stackId:  stacks[MINE],
			position: 1,
		},
		schoolId: 8,
	})
	csm.processRequest(&AddCardRequest{
		position: CardPosition{
			stackId:  stacks[URBAN_TEMPLE],
			position: 0,
		},
		schoolId: 13,
	})
	csm.processRequest(&AddCardRequest{
		position: CardPosition{
			stackId:  stacks[URBAN_LAB],
			position: 0,
		},
		schoolId: 9,
	})
	csm.processRequest(&AddCardRequest{
		position: CardPosition{
			stackId:  stacks[GOVERNMENT],
			position: 0,
		},
		schoolId: 35,
	})
	// Prepare initial tokens for agriculture and bronze
	agricultureCard := csm.getFirstCard(stacks[FARM])
	bronzeCard := csm.getFirstCard(stacks[MINE])

	oilCard := csm.cardStacks[stacks[MINE]][1]
	warriorCard := csm.getFirstCard(stacks[MILI_INFANTRY])
	governmentCard := csm.getFirstCard(stacks[GOVERNMENT])

	labCard := csm.getFirstCard(stacks[URBAN_LAB])

	game.cardTokenManager.setTokenCount(agricultureCard.id, TOKEN_YELLOW, 2)
	game.cardTokenManager.setTokenCount(bronzeCard.id, TOKEN_YELLOW, 2)
	game.cardTokenManager.setTokenCount(oilCard.id, TOKEN_YELLOW, 1)
	game.cardTokenManager.setTokenCount(warriorCard.id, TOKEN_YELLOW, 1)
	game.cardTokenManager.setTokenCount(governmentCard.id, TOKEN_WHITE, 4)
	game.cardTokenManager.setTokenCount(governmentCard.id, TOKEN_RED, 2)

	game.cardTokenManager.setTokenCount(labCard.id, TOKEN_YELLOW, 1)

	// Prepare special bank manager
	specialTokenManager := NewTokenBankUniversalManager()
	specialTokenManager.setTokenCount(FREE_WORKER, TOKEN_YELLOW, 1)
	specialTokenManager.setTokenCount(FREE_YELLOW, TOKEN_YELLOW, 15)
	specialTokenManager.setTokenCount(FREE_BLUE, TOKEN_BLUE, 16)

	// Prepare tech bank Manager
	techTokenManager := NewTokenBankUniversalManager()

	// Finish
	return &PlayerBoard{
		game:                game,
		stacks:              stacks,
		specialTokenManager: specialTokenManager,
		techTokenManager:    techTokenManager,
	}
}

func (p *PlayerBoard) getCultureTotal() int {
	return p.specialTokenManager.getTokenCount(CULTURE_COUNTER, TOKEN_DEFAULT)
}

func (p *PlayerBoard) getTechTotal() int {
	return p.specialTokenManager.getTokenCount(TECH_COUNTER, TOKEN_DEFAULT)
}

func (p *PlayerBoard) getUsableWhiteTokens() int {
	csm := p.game.cardStackManager
	governmentCard := csm.getFirstCard(p.stacks[GOVERNMENT])
	return p.game.cardTokenManager.getTokenCount(governmentCard.id, TOKEN_WHITE)
}

func (p *PlayerBoard) getUsableRedTokens() int {
	csm := p.game.cardStackManager
	governmentCard := csm.getFirstCard(p.stacks[GOVERNMENT])
	return p.game.cardTokenManager.getTokenCount(governmentCard.id, TOKEN_RED)
}

func (p *PlayerBoard) getFreeWorkers() int {
	return p.specialTokenManager.getTokenCount(FREE_WORKER, TOKEN_YELLOW)
}

func (p *PlayerBoard) getFreeYellowTokens() int {
	return p.specialTokenManager.getTokenCount(FREE_YELLOW, TOKEN_YELLOW)
}

func (p *PlayerBoard) getFreeBlueTokens() int {
	return p.specialTokenManager.getTokenCount(FREE_BLUE, TOKEN_BLUE)
}

func (p *PlayerBoard) getHandSize() int {
	csm := p.game.cardStackManager
	return csm.getStackSize(p.stacks[HAND])
}

func (p *PlayerBoard) getResourceCorruption() int {
	freeBlue := p.getFreeBlueTokens()
	switch freeBlue {
	case 0:
		return 6
	case 1, 2, 3, 4, 5:
		return 4
	case 6, 7, 8, 9, 10:
		return 2
	default:
		return 0
	}
}
func (p *PlayerBoard) getCropConsume() int {
	freeYellow := p.getFreeYellowTokens()
	switch freeYellow {
	case 0:
		return 6
	case 1, 2, 3, 4:
		return 4
	case 5, 6, 7, 8:
		return 3
	case 9, 10, 11, 12:
		return 2
	case 13, 14, 15, 16:
		return 1
	default:
		return 0
	}
}

func (p *PlayerBoard) removeUsableWhiteTokens(count int) {
	csm := p.game.cardStackManager
	governmentCard := csm.getFirstCard(p.stacks[GOVERNMENT])
	p.game.cardTokenManager.modifyToken(governmentCard.id, TOKEN_WHITE, -count)
}

func (p *PlayerBoard) refillWhiteRedTokens() {
	csm := p.game.cardStackManager
	governmentCard := csm.getFirstCard(p.stacks[GOVERNMENT])
	p.game.cardTokenManager.processRequest(&SetTokenRequest{
		bankId:     governmentCard.id,
		tokenType:  TOKEN_WHITE,
		tokenCount: p.calcWhiteTokenLimit(),
	})
	p.game.cardTokenManager.processRequest(&SetTokenRequest{
		bankId:     governmentCard.id,
		tokenType:  TOKEN_RED,
		tokenCount: p.calcRedTokenLimit(),
	})
}

func (p *PlayerBoard) isBuildingWonder() bool {
	csm := p.game.cardStackManager
	return csm.getStackSize(p.stacks[WONDER_NOT_COMPLETED]) != 0
}

func (p *PlayerBoard) getBlueTokensOnCurrentWonder() int {
	csm := p.game.cardStackManager
	wonderCard := csm.getFirstCard(p.stacks[WONDER_NOT_COMPLETED])
	return p.game.cardTokenManager.getTokenCount(wonderCard.id, TOKEN_BLUE)
}

func (p *PlayerBoard) getCurrentHandSize() int {
	csm := p.game.cardStackManager
	return csm.getStackSize(p.stacks[HAND])
}

func (p *PlayerBoard) hasLeader() bool {
	csm := p.game.cardStackManager
	return csm.getStackSize(p.stacks[LEADER]) > 0
}
func (p *PlayerBoard) calcPlayerFeatureSum(f func(*CardSchool) int, canBeNegative bool) int {
	csm := p.game.cardStackManager
	allSchools := p.game.cardSchools
	result := 0
	// Government
	governmentCard := csm.getFirstCard(p.stacks[GOVERNMENT])
	result += f(allSchools[governmentCard.schoolId])

	// Army, Farms, mines and urban buildings counts yellow tokens
	for _, t := range []int{MILI_INFANTRY,
		MILI_CAVALRY,
		MILI_ARTILERY,
		MILI_AIRFORCE,
		FARM,
		MINE,
		URBAN_TEMPLE,
		URBAN_LAB,
		URBAN_ARENA,
		URBAN_LIBRARY,
		URBAN_THEATER,
	} {
		for _, card := range csm.cardStacks[t] {
			result += f(allSchools[card.schoolId]) *
				p.game.cardTokenManager.getTokenCount(card.id, TOKEN_YELLOW)
		}
	}

	// Special technology
	for _, t := range []int{
		TECH_SPECIAL_CIVIL,
		TECH_SPECIAL_WARFARE,
		TECH_SPECIAL_COLONIZE,
		TECH_SPECIAL_CONSTRUCTION,
	} {
		specialTechCard := csm.getFirstCard(p.stacks[t])
		if specialTechCard != nil {
			result += f(allSchools[specialTechCard.schoolId])
		}
	}

	// Leader
	for _, card := range csm.cardStacks[p.stacks[LEADER]] {
		result += f(allSchools[card.schoolId])
	}
	// Wonders
	for _, card := range csm.cardStacks[p.stacks[WONDER_COMPLETED]] {
		result += f(allSchools[card.schoolId])
	}

	if result < 0 && !canBeNegative {
		result = 0
	}
	return result
}

func (p *PlayerBoard) calcWhiteTokenLimit() int {
	return p.calcPlayerFeatureSum(func(school *CardSchool) int {
		return school.productionWhiteToken
	}, false)
}

func (p *PlayerBoard) calcRedTokenLimit() int {
	return p.calcPlayerFeatureSum(func(school *CardSchool) int {
		return school.productionRedToken
	}, false)
}

func (p *PlayerBoard) calcCultureInc() int {
	return p.calcPlayerFeatureSum(func(school *CardSchool) int {
		return school.productionCulture
	}, false)
}

func (p *PlayerBoard) calcTechInc() int {
	return p.calcPlayerFeatureSum(func(school *CardSchool) int {
		return school.productionTech
	}, false)
}

func (p *PlayerBoard) calcPower() int {
	return p.calcPlayerFeatureSum(func(school *CardSchool) int {
		return school.productionPower
	}, false)
}

func (p *PlayerBoard) canTakeCardFromWheel(index int) bool {
	card := p.game.getCardOnGreatWheel(index)
	// Assure card exists
	if card == nil {
		return false
	}

	// White tokens enough
	if p.getUsableWhiteTokens() < p.takeCardFromWheelCost(index) {
		return false
	}

	// Cannot take card if hand is full
	if p.getCurrentHandSize() >= p.calcWhiteTokenLimit() {
		return false
	}

	school := p.game.cardSchools[card.schoolId]
	// Cannot take wonder if another is under construction
	if school.hasType(CARDTYPE_WONDER) {
		if p.isBuildingWonder() {
			return false
		}
	}

	// Cannot take duplicate tech cards
	if school.hasType(CARDTYPE_TECH) {
		if p.techTokenManager.getTokenCount(card.schoolId, TOKEN_DEFAULT) > 0 {
			return false
		}
	}

	// Cannot take leader with duplicate ages
	if school.hasType(CARDTYPE_LEADER) {
		if p.specialTokenManager.getTokenCount(LEADER_A_TAKEN+school.age,
			TOKEN_DEFAULT) > 0 {
			return false
		}
	}
	return true
}

func (p *PlayerBoard) takeCardFromWheelCost(index int) (cost int) {
	csm := p.game.cardStackManager
	switch index {
	case 0, 1, 2, 3, 4:
		cost = 1
	case 5, 6, 7, 8:
		cost = 2
	case 9, 10, 11, 12:
		cost = 3
	default:
		return -1
	}

	card := csm.getFirstCard(p.game.greatWheel[index])
	school := p.game.cardSchools[card.schoolId]
	if school.hasType(CARDTYPE_WONDER) {
		cost += csm.getStackSize(p.stacks[WONDER_COMPLETED])
		// TODO: Taj Mahal, Hammurabi, Mich. here
	}
	return
}

func (p *PlayerBoard) takeCardFromWheel(index int) {
	csm := p.game.cardStackManager
	card := p.game.getCardOnGreatWheel(index)

	if card == nil {
		return
	}
	cost := p.takeCardFromWheelCost(index)
	if p.canTakeCardFromWheel(index) {
		// Move the card
		school := p.game.cardSchools[card.schoolId]
		if school.hasType(CARDTYPE_WONDER) {
			csm.processRequest(&MoveCardRequest{
				sourcePosition: CardPosition{
					stackId:  p.game.greatWheel[index],
					position: 0,
				},
				targetPosition: CardPosition{
					stackId:  p.stacks[WONDER_NOT_COMPLETED],
					position: 0,
				},
			})
		} else {
			csm.processRequest(&MoveCardRequest{
				sourcePosition: CardPosition{
					stackId:  p.game.greatWheel[index],
					position: 0,
				},
				targetPosition: CardPosition{
					stackId:  p.stacks[HAND],
					position: 0,
				},
			})
		}

		// Remove white tokens used
		p.removeUsableWhiteTokens(cost)

		// Mark tech card, no duplicate
		if school.hasType(CARDTYPE_TECH) {
			p.techTokenManager.processRequest(&AddTokenRequest{
				bankId:     school.schoolId,
				tokenType:  TOKEN_DEFAULT,
				tokenCount: 1,
			})
		}

		// Mark leader
		if school.hasType(CARDTYPE_LEADER) {
			p.specialTokenManager.processRequest(&AddTokenRequest{
				bankId:     LEADER_A_TAKEN + school.age,
				tokenType:  TOKEN_DEFAULT,
				tokenCount: 1,
			})
		}
	}
}

func (p *PlayerBoard) getResourceTotal() (result int) {
	csm := p.game.cardStackManager
	result = 0
	for _, card := range csm.cardStacks[p.stacks[MINE]] {
		school := p.game.cardSchools[card.schoolId]
		amount := p.game.cardTokenManager.getTokenCount(card.id, TOKEN_BLUE)
		result += school.productionResource * amount
	}
	return
}

func (p *PlayerBoard) getCropTotal() (result int) {
	csm := p.game.cardStackManager
	result = 0
	for _, card := range csm.cardStacks[p.stacks[FARM]] {
		school := p.game.cardSchools[card.schoolId]
		amount := p.game.cardTokenManager.getTokenCount(card.id, TOKEN_BLUE)
		result += school.productionCrop * amount
	}
	return
}

// Calculate the "best" solution to spent an amount of resource/crop
func (p *PlayerBoard) tryArrangeSpend(q, unit []int, free, spent int) (possible bool, arranged []int) {
	for i, _ := range q {
		for {
			if q[i] > 0 && unit[i] <= spent {
				q[i]--
				free++
				spent -= unit[i]
			} else {
				break
			}
		}

		if spent == 0 {
			return true, q
		}

		if q[i] > 0 && unit[i] > spent {
			q[i]--
			free++
			spent -= unit[i] // Spent will be < 0
			for {
				i--
				for {
					if unit[i]+spent <= 0 {
						if free > 0 {
							free--
						} else {
							return true, q
						}
						q[i] += 1
						spent += unit[i]
					} else {
						break
					}
				}

				if spent == 0 {
					return true, q
				}
			}
		}
	}

	return false, nil
}

func (p *PlayerBoard) spendResource(amount int) {
	csm := p.game.cardStackManager
	quantity := make([]int, csm.getStackSize(p.stacks[MINE]))
	unit := make([]int, csm.getStackSize(p.stacks[MINE]))

	for i, card := range csm.cardStacks[p.stacks[MINE]] {
		school := p.game.cardSchools[card.schoolId]
		q := p.game.cardTokenManager.getTokenCount(card.id, TOKEN_BLUE)
		quantity[i] = q
		unit[i] = school.productionResource
	}
	sum := 0
	for _, q := range quantity {
		sum += q
	}

	possible, newQuantity := p.tryArrangeSpend(quantity, unit, p.getFreeBlueTokens(), amount)
	if !possible {
		return
	}
	if amount == 0 {
		fmt.Println(p.getResourceTotal())
		fmt.Println(quantity)
		fmt.Println(newQuantity)
	}
	newSum := 0
	for _, q := range newQuantity {
		newSum += q
	}

	for i, card := range csm.cardStacks[p.stacks[MINE]] {
		p.game.cardTokenManager.processRequest(&SetTokenRequest{
			bankId:     card.id,
			tokenType:  TOKEN_BLUE,
			tokenCount: newQuantity[i],
		})
	}
	p.specialTokenManager.processRequest(&AddTokenRequest{
		bankId:     FREE_BLUE,
		tokenType:  TOKEN_BLUE,
		tokenCount: sum - newSum,
	})
	return
}

func (p *PlayerBoard) spendAllResource() {
	p.spendResource(p.getResourceTotal())
}

func (p *PlayerBoard) spendCrop(amount int) {
	csm := p.game.cardStackManager
	quantity := make([]int, csm.getStackSize(p.stacks[FARM]))
	unit := make([]int, csm.getStackSize(p.stacks[FARM]))

	for i, card := range csm.cardStacks[p.stacks[FARM]] {
		school := p.game.cardSchools[card.schoolId]
		q := p.game.cardTokenManager.getTokenCount(card.id, TOKEN_BLUE)
		quantity[i] = q
		unit[i] = school.productionCrop
	}
	sum := 0
	for _, q := range quantity {
		sum += q
	}

	possible, newQuantity := p.tryArrangeSpend(quantity, unit, p.getFreeBlueTokens(), amount)
	if !possible {
		return
	}
	newSum := 0
	for _, q := range newQuantity {
		newSum += q
	}

	for i, card := range csm.cardStacks[p.stacks[FARM]] {
		p.game.cardTokenManager.processRequest(&SetTokenRequest{
			bankId:     card.id,
			tokenType:  TOKEN_BLUE,
			tokenCount: newQuantity[i],
		})
	}
	p.specialTokenManager.processRequest(&AddTokenRequest{
		bankId:     FREE_BLUE,
		tokenType:  TOKEN_BLUE,
		tokenCount: sum - newSum,
	})
	return
}

func (p *PlayerBoard) spendAllCrop() {
	p.spendCrop(p.getCropTotal())
}

func (p *PlayerBoard) consumeCrop() {
	c := p.getCropConsume()
	crop := p.getCropTotal()
	if c < crop {
		p.spendCrop(c)
	} else {
		p.spendAllCrop()
		p.specialTokenManager.processRequest(&RemoveTokenRequest{
			bankId:     CULTURE_COUNTER,
			tokenType:  TOKEN_DEFAULT,
			tokenCount: 4 * (c - crop),
		})
	}
}

func (p *PlayerBoard) corrupt() {
	c := p.getResourceCorruption()
	//fmt.Println(p.getFreeBlueTokens(), c)

	res := p.getResourceTotal()
	if c <= res {
		p.spendResource(c)
	} else {
		c -= res
		p.spendAllResource()
		crop := p.getCropTotal()
		if c < crop {
			p.spendCrop(c)
		} else {
			p.spendAllCrop()
			// TODO: reduce culture?
		}
	}
}

func (p *PlayerBoard) productCrop() {
	csm := p.game.cardStackManager
	// Need reverse iterate
	for i := csm.getStackSize(p.stacks[FARM]) - 1; i >= 0; i-- {
		card := csm.cardStacks[p.stacks[FARM]][i]
		amount := p.game.cardTokenManager.getTokenCount(card.id, TOKEN_YELLOW)
		if amount > p.getFreeBlueTokens() {
			amount = p.getFreeBlueTokens()
		}
		fmt.Println(p.getFreeYellowTokens())
		if amount == 0 {
			return
		}
		p.specialTokenManager.processRequest(&RemoveTokenRequest{
			bankId:     FREE_BLUE,
			tokenType:  TOKEN_BLUE,
			tokenCount: amount,
		})
		fmt.Println(p.getFreeYellowTokens())
		p.game.cardTokenManager.processRequest(&AddTokenRequest{
			bankId:     card.id,
			tokenType:  TOKEN_BLUE,
			tokenCount: amount,
		})
		fmt.Println(p.getFreeYellowTokens())
	}
}

func (p *PlayerBoard) productResource() {
	csm := p.game.cardStackManager
	// Need reverse iterate
	for i := csm.getStackSize(p.stacks[MINE]) - 1; i >= 0; i-- {
		card := csm.cardStacks[p.stacks[MINE]][i]
		amount := p.game.cardTokenManager.getTokenCount(card.id, TOKEN_YELLOW)
		if amount > p.getFreeBlueTokens() {
			amount = p.getFreeBlueTokens()
		}

		if amount == 0 {
			return
		}
		p.specialTokenManager.processRequest(&RemoveTokenRequest{
			bankId:     FREE_BLUE,
			tokenType:  TOKEN_BLUE,
			tokenCount: amount,
		})
		p.game.cardTokenManager.processRequest(&AddTokenRequest{
			bankId:     card.id,
			tokenType:  TOKEN_BLUE,
			tokenCount: amount,
		})
	}
}

func (p *PlayerBoard) productCultureAndTech() {
	p.specialTokenManager.processRequest(&AddTokenRequest{
		bankId:     CULTURE_COUNTER,
		tokenType:  TOKEN_DEFAULT,
		tokenCount: p.calcCultureInc(),
	})
	p.specialTokenManager.processRequest(&AddTokenRequest{
		bankId:     TECH_COUNTER,
		tokenType:  TOKEN_DEFAULT,
		tokenCount: p.calcTechInc(),
	})
}

func (p *PlayerBoard) doProductionPhase() {
	p.productCultureAndTech()
	p.corrupt()
	p.productCrop()
	p.consumeCrop()
	p.productResource()
	p.refillWhiteRedTokens()
}

func (p *PlayerBoard) canPlayCard(card Card, attachment interface{}) bool {
	school := p.game.cardSchools[card.schoolId]
	if school.hasType(CARDTYPE_TECH) {
		// TODO tech cost reduction effects here
		return p.getTechTotal() >= school.tech
	} else if school.hasType(CARDTYPE_ACTION) {
		// TODO
		return false
	} else if school.hasType(CARDTYPE_LEADER) {
		return true
	}
	return false
}

func (p *PlayerBoard) canPlayHand(index int, attachment interface{}) bool {
	csm := p.game.cardStackManager
	if index < 0 || index >= p.getHandSize() {
		return false
	}

	if p.getUsableWhiteTokens() < p.takeCardFromWheelCost(index) {
		return false
	}

	card := csm.cardStacks[p.stacks[HAND]][index]
	return p.canPlayCard(card, attachment)
}

func (p *PlayerBoard) playStructureTechCard(card Card, index int, stackId int) {
	csm := p.game.cardStackManager
	school := p.game.cardSchools[card.schoolId]
	stack := csm.cardStacks[p.stacks[stackId]]
	i := 0
	for {
		if i >= len(stack) {
			break
		}

		cardAti := stack[i]
		schoolAti := p.game.cardSchools[cardAti.schoolId]
		if schoolAti.age >= school.age {
			break
		}
		i += 1
	}
	csm.processRequest(&MoveCardRequest{
		sourcePosition: CardPosition{
			stackId:  p.stacks[HAND],
			position: index,
		},
		targetPosition: CardPosition{
			stackId:  p.stacks[stackId],
			position: i,
		},
	})
}

func (p *PlayerBoard) playCard(card Card, index int, attachment interface{}) {
	csm := p.game.cardStackManager
	school := p.game.cardSchools[card.schoolId]
	if school.hasType(CARDTYPE_TECH) {
		// TODO: all tech cost modifiers
		p.specialTokenManager.processRequest(&RemoveTokenRequest{
			bankId:     TECH_COUNTER,
			tokenType:  TOKEN_DEFAULT,
			tokenCount: school.tech,
		})
		if school.hasType(CARDTYPE_TECH_FARM) {
			p.playStructureTechCard(card, index, FARM)
		} else if school.hasType(CARDTYPE_TECH_MINE) {
			p.playStructureTechCard(card, index, MINE)
		} else if school.hasType(CARDTYPE_TECH_MILLI_INFANTRY) {
			p.playStructureTechCard(card, index, MILI_INFANTRY)
		} else if school.hasType(CARDTYPE_TECH_MILLI_CAVALRY) {
			p.playStructureTechCard(card, index, MILI_CAVALRY)
		} else if school.hasType(CARDTYPE_TECH_MILLI_ARTILLERY) {
			p.playStructureTechCard(card, index, MILI_ARTILERY)
		} else if school.hasType(CARDTYPE_TECH_MILLI_AIRFORCE) {
			p.playStructureTechCard(card, index, MILI_AIRFORCE)
		} else if school.hasType(CARDTYPE_TECH_URBAN_TEMPLE) {
			p.playStructureTechCard(card, index, URBAN_TEMPLE)
		} else if school.hasType(CARDTYPE_TECH_URBAN_LAB) {
			p.playStructureTechCard(card, index, URBAN_LAB)
		} else if school.hasType(CARDTYPE_TECH_URBAN_ARENA) {
			p.playStructureTechCard(card, index, URBAN_ARENA)
		} else if school.hasType(CARDTYPE_TECH_URBAN_LIBRARY) {
			p.playStructureTechCard(card, index, URBAN_LIBRARY)
		} else if school.hasType(CARDTYPE_TECH_URBAN_THEATER) {
			p.playStructureTechCard(card, index, URBAN_THEATER)
		} else if school.hasType(CARDTYPE_TECH_GOVERNMENT) {

		}
	} else if school.hasType(CARDTYPE_ACTION) {
		// TODO
	} else if school.hasType(CARDTYPE_LEADER) {
		if p.hasLeader() {
			// TODO: Homer
			p.removeUsableWhiteTokens(-1)

			csm.processRequest(&BanishCardRequest{
				position: CardPosition{
					stackId:  p.stacks[LEADER],
					position: 0,
				},
			})
		}
		csm.processRequest(&MoveCardRequest{
			sourcePosition: CardPosition{
				stackId:  p.stacks[HAND],
				position: index,
			},
			targetPosition: CardPosition{
				stackId:  p.stacks[LEADER],
				position: 0,
			},
		})
	}
}

func (p *PlayerBoard) playHand(index int, attachment interface{}) {
	csm := p.game.cardStackManager
	p.removeUsableWhiteTokens(1)
	card := csm.cardStacks[p.stacks[HAND]][index]
	p.playCard(card, index, attachment)
}
