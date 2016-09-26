package main

import (
	"fmt"
	"math/rand"
)

type UserStackId int

type TtaGameOptions struct {
	PlayerCount  int
	LoveAndPeace bool
}

type PendingActionType int
const (
	CIVIL PendingActionType = 1
	REMOVE_YELLOW           = 2
	REMOVE_BLUE             = 3
)

type PendingAction struct {
	PendingActionType PendingActionType
}

type MoveType int
const (
	CIVIL_FETCH_CARD MoveType = iota
	CIVIL_PLAY_CARD
	CIVIL_INC_POP
	CIVIL_BUILD
	CIVIL_BUILD_WONDER
	CIVIL_UPGRADE
	CIVIL_SPECIAL_ABILITY
	CIVIL_END
	CHOOSE_YELLOW
	CHOOSE_BLUE
)

type Move struct {
	FromPlayer int
	MoveType   MoveType

	Data []int
}

type TtaGame struct {
	cardStackManager   *CardStackUniversalManager
	globalTokenManager *TokenBankUniversalManager
	cardTokenManager   *TokenBankUniversalManager

	// All card schools
	cardSchools map[int]*CardSchool

	greatWheel []int // 13 stacks
	ageStacks  []int // 4 stacks by age
	players    []*PlayerBoard

	// Pending action
	CurrentPlayer int
	PendingAction PendingAction
}

func NewTta(options *TtaGameOptions) (result *TtaGame) {
	if options.PlayerCount < 1 || options.PlayerCount > 4 {
		return nil
	}
	game := &TtaGame{
		cardStackManager:   NewCardStackUniversalManager(),
		globalTokenManager: NewTokenBankUniversalManager(),
		cardTokenManager:   NewTokenBankUniversalManager(),
		greatWheel:         make([]int, 13),
		ageStacks:          make([]int, 4),
		players:            make([]*PlayerBoard, 2),
	}
	game.cardSchools = InitBasicCardSchools()
	for i := 0; i < options.PlayerCount; i++ {
		game.players[i] = initPlayerBoard(game)
	}

	for i := 0; i < 13; i++ {
		game.greatWheel[i] = game.cardStackManager.newStack()
	}
	for i := 0; i < 4; i++ {
		game.ageStacks[i] = game.cardStackManager.newStack()
	}

	game.initBasicCards(options)
	game.refillWheels()
	game.banishAgeACards()
	return game
}

func (g *TtaGame) checkDecay() {

}

func (g *TtaGame) initBasicCards(options *TtaGameOptions) {
	// Fill all the civil cards
	csm := g.cardStackManager
	for id, school := range InitBasicCardSchools() {
		if school.schoolId == 1 ||
			school.schoolId == 5 ||
			school.schoolId == 9 ||
			school.schoolId == 13 ||
			school.schoolId == 25 ||
			school.schoolId == 35 {
			continue
		}
		if options.PlayerCount == 1 {  // Solo test mode
			for i := 0; i < school.cardCounts[0]; i++ {
				csm.processRequest(&AddCardToTopRequest{
					schoolId: id,
					stackId:  g.ageStacks[school.age],
				})
			}
		} else {
			for i := 0; i < school.cardCounts[options.PlayerCount - 2]; i++ {
				csm.processRequest(&AddCardToTopRequest{
					schoolId: id,
					stackId:  g.ageStacks[school.age],
				})
			}
		}
	}

	for i := 0; i <= 3; i++ {
		cardCount := csm.getStackSize(g.ageStacks[i])
		randomPerm := rand.Perm(cardCount)
		for j := 0; j < cardCount; j++ {
			csm.processRequest(&SwapCardRequest{
				sourcePosition: CardPosition{
					stackId:  g.ageStacks[i],
					position: j,
				},
				targetPosition: CardPosition{
					stackId:  g.ageStacks[i],
					position: randomPerm[j],
				},
			})
		}
	}
}

func (g *TtaGame) banishAgeACards() {
	csm := g.cardStackManager

	csm.processRequest(&BanishAllCardsInStackRequest{
		stackId: g.ageStacks[0],
	})
}

func (g *TtaGame) refillWheels() {
	csm := g.cardStackManager
	search := 0
	currentAge := 0
	for s := 0; s < 13; s++ {
		if len(csm.cardStacks[g.greatWheel[s]]) > 0 { // Have card at the position
			continue
		}

		if search <= s {
			search = s + 1
		}

		// Search for next position which have a card
		for {
			if search >= 13 || len(csm.cardStacks[g.greatWheel[search]]) > 0 {
				break
			}
			search++
		}

		if search < 13 {
			csm.processRequest(&MoveCardRequest{
				sourcePosition: CardPosition{
					stackId:  g.greatWheel[search],
					position: 0,
				},
				targetPosition: CardPosition{
					stackId:  g.greatWheel[s],
					position: 0,
				},
			})
		} else {
			for {
				if currentAge > 3 || len(csm.cardStacks[g.ageStacks[currentAge]]) > 0 {
					break
				}
				currentAge++
			}

			if currentAge <= 3 {
				csm.processRequest(&MoveCardRequest{
					sourcePosition: CardPosition{
						stackId:  g.ageStacks[currentAge],
						position: 0,
					},
					targetPosition: CardPosition{
						stackId:  g.greatWheel[s],
						position: 0,
					},
				})
			}
		}
	}
}

func (g *TtaGame) getCardOnGreatWheel(index int) *Card {
	csm := g.cardStackManager
	return csm.getFirstCard(g.greatWheel[index])
}

func (g *TtaGame) weedOut(position int) {
	csm := g.cardStackManager
	for i := 0; i < position; i++ {
		if csm.getStackSize(g.greatWheel[i]) > 0 {
			csm.processRequest(&BanishCardRequest{
				position: CardPosition{
					stackId:  g.greatWheel[i],
					position: 0,
				},
			})
		}
	}
}

func (g *TtaGame) processCivilMove(move *Move) (err error) {
	if move.FromPlayer != g.CurrentPlayer {
		return fmt.Errorf("Not current player.")
	}
	p := g.players[g.CurrentPlayer]
	switch move.MoveType {
	case CIVIL_FETCH_CARD:
		if len(move.Data) != 1 {
			return fmt.Errorf("Invalid fetch command.")
		}
		index := move.Data[0]
		if !p.canTakeCardFromWheel(index) {
			return fmt.Errorf("Invalid fetch command.")
		}
		p.takeCardFromWheel(index)
	case CIVIL_PLAY_CARD:
		if len(move.Data) < 1 {
			return fmt.Errorf("Invalid play command.")
		}
		index := move.Data[0]
		var attachment interface{}
		if len(move.Data) > 1 {
			attachment = move.Data[1:]
		} else {
			attachment = nil
		}
		if !p.canPlayHand(index, attachment) {
			return fmt.Errorf("Invalid play command")
		}
		p.playHand(index, attachment)
	case CIVIL_INC_POP:
		if !p.canIncreasePop() {
			return fmt.Errorf("Invalid incpop command")
		}
		p.increasePop()
	case CIVIL_BUILD:
		if len(move.Data) < 2 {
			return fmt.Errorf("Invalid build command.")
		}
		stack := move.Data[0]
		index := move.Data[0]
		if !p.canBuild(stack, index, 0) {
			return fmt.Errorf("Invalid build command")
		}
		p.build(stack, index, 0)
	case CIVIL_BUILD_WONDER:
		if len(move.Data) < 1 {
			return fmt.Errorf("Invalid buildwonder command.")
		}
		step := move.Data[0]
		if !p.canBuildWonder(step, 0) {
			return fmt.Errorf("Invalid buildwonder command")
		}
		p.buildWonder(step, 0)
	case CIVIL_UPGRADE:
		if len(move.Data) < 3 {
			return fmt.Errorf("Invalid upgrade command.")
		}
		stack := move.Data[0]
		index1 := move.Data[1]
		index2 := move.Data[2]
		if !p.canUpgrade(stack, index1, index2, 0) {
			return fmt.Errorf("Invalid upgrade command")
		}
		p.upgrade(stack, index1, index2, 0)
	case CIVIL_SPECIAL_ABILITY:
		if len(move.Data) < 1 {
			return fmt.Errorf("Invalid specialability command.")
		}
		sa := move.Data[0]
		var attachment interface{}
		if len(move.Data) > 1 {
			attachment = move.Data[1:]
		} else {
			attachment = nil
		}
		if !p.canUseCivilSpecialAbility(sa, attachment) {
			return fmt.Errorf("Invalid specialability command")
		}
		p.useCivilSpecialAbility(sa, attachment)
	case CIVIL_END:
	}
	return nil
}

func (g *TtaGame) ProcessMove(move *Move) (err error) {
	switch g.PendingAction.PendingActionType {
	case CIVIL:
		return g.processCivilMove(move)
	default:
		return fmt.Errorf("Invalid PendingAction")
	}
}
