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
	CIVIL            PendingActionType = 1
	DISCARD_MILITARY                   = 2
	REMOVE_YELLOW                      = 3
	REMOVE_BLUE                        = 4
)

type PendingAction struct {
	Type PendingActionType
	Data []int
}

type MoveType int

const (
	MOVE_FETCH_CARD MoveType = iota
	MOVE_PLAY_CIVIL_CARD
	MOVE_PLAY_MILITARY_CARD
	MOVE_INC_POP
	MOVE_BUILD
	MOVE_BUILD_WONDER
	MOVE_UPGRADE
	MOVE_SPECIAL_ABILITY
	MOVE_END
	MOVE_DISCARD_MILITARY_CARDS
	CHOOSE_YELLOW
	CHOOSE_BLUE
)

type Move struct {
	FromPlayer int
	MoveType   MoveType

	Data []int
}

type TtaGame struct {
	// Game options
	options *TtaGameOptions

	cardStackManager   *CardStackUniversalManager
	globalTokenManager *TokenBankUniversalManager
	cardTokenManager   *TokenBankUniversalManager

	// All card schools
	cardSchools map[int]*CardSchool

	greatWheel       []int // 13 stacks
	ageStacks        []int // 4 stacks by age
	miliDecks        []int // 4 stacks by age, first element is zero
	miliDiscardDecks []int
	futureEventsDeck int
	nowEventsDeck    int
	pastEventsDeck   int
	publicTacticDeck int
	players          []*PlayerBoard

	// Pending action
	CurrentPlayer int
	RoundCount    int
	StateStack    []StateHolder
}

func NewTta(options *TtaGameOptions) (result *TtaGame) {
	if options.PlayerCount < 1 || options.PlayerCount > 4 {
		return nil
	}
	game := &TtaGame{
		options:            options,
		cardStackManager:   NewCardStackUniversalManager(),
		globalTokenManager: NewTokenBankUniversalManager(),
		cardTokenManager:   NewTokenBankUniversalManager(),
		greatWheel:         make([]int, 13),
		ageStacks:          make([]int, 4),
		miliDecks:          make([]int, 4),
		miliDiscardDecks:   make([]int, 4),
		players:            make([]*PlayerBoard, 2),

		CurrentPlayer: 0,
		RoundCount:    0,
		StateStack:    make([]StateHolder, 0),
	}
	game.cardSchools = InitBasicCardSchools()
	for i := 0; i < options.PlayerCount; i++ {
		game.players[i] = initPlayerBoard(game)
		game.players[i].setUsableWhiteTokens(i + 1)
	}

	for i := 0; i < 13; i++ {
		game.greatWheel[i] = game.cardStackManager.newStack()
	}
	for i := 0; i < 4; i++ {
		game.ageStacks[i] = game.cardStackManager.newStack()
		game.miliDecks[i] = game.cardStackManager.newStack()
		game.miliDiscardDecks[i] = game.cardStackManager.newStack()
	}
	game.futureEventsDeck = game.cardStackManager.newStack()
	game.nowEventsDeck = game.cardStackManager.newStack()
	game.pastEventsDeck = game.cardStackManager.newStack()
	game.publicTacticDeck = game.cardStackManager.newStack()

	game.initBasicCards(options)
	game.refillWheels()

	game.StateStack = []StateHolder{
		&TurnStartStateHolder{
			base: BaseStateHolder{
				game: game,
			},
		},
	}
	game.Initialize()
	return game
}

func (g *TtaGame) getCurrentAge() int {
	csm := g.cardStackManager
	for i := 0; i <= 3; i++ {
		if csm.getStackSize(g.ageStacks[i]) > 0 {
			return i
		}
	}
	return 4
}

func (g *TtaGame) reshuffleMilitaryDeck() {
	csm := g.cardStackManager
	for {
		if csm.getStackSize(g.miliDiscardDecks[g.getCurrentAge()]) <= 0 {
			break
		}
		csm.processRequest(&MoveCardRequest{
			sourcePosition: CardPosition{
				stackId:  g.miliDiscardDecks[g.getCurrentAge()],
				position: 0,
			},
			targetPosition: CardPosition{
				stackId:  g.miliDecks[g.getCurrentAge()],
				position: 0,
			},
		})
	}

	cardCount := csm.getStackSize(g.miliDecks[g.getCurrentAge()])
	randomPerm := rand.Perm(cardCount)
	for j := 0; j < cardCount; j++ {
		csm.processRequest(&SwapCardRequest{
			sourcePosition: CardPosition{
				stackId:  g.miliDecks[g.getCurrentAge()],
				position: j,
			},
			targetPosition: CardPosition{
				stackId:  g.miliDecks[g.getCurrentAge()],
				position: randomPerm[j],
			},
		})
	}
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
		var cardCountIdx int
		if options.PlayerCount == 1 { // Solo test mode
			cardCountIdx = 0
		} else {
			cardCountIdx = options.PlayerCount - 2
		}
		for i := 0; i < school.cardCounts[cardCountIdx]; i++ {
			if school.isCivilCard() {
				csm.processRequest(&AddCardToTopRequest{
					schoolId: id,
					stackId:  g.ageStacks[school.age],
				})
			} else {
				csm.processRequest(&AddCardToTopRequest{
					schoolId: id,
					stackId:  g.miliDecks[school.age],
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

	for i := 0; i <= 3; i++ {
		cardCount := csm.getStackSize(g.miliDecks[i])
		randomPerm := rand.Perm(cardCount)
		for j := 0; j < cardCount; j++ {
			csm.processRequest(&SwapCardRequest{
				sourcePosition: CardPosition{
					stackId:  g.miliDecks[i],
					position: j,
				},
				targetPosition: CardPosition{
					stackId:  g.miliDecks[i],
					position: randomPerm[j],
				},
			})
		}
	}

	for i := 0; i < options.PlayerCount+2; i++ {
		csm.processRequest(&MoveCardRequest{
			sourcePosition: CardPosition{
				stackId:  g.miliDecks[0],
				position: csm.getStackSize(g.miliDecks[0]) - 1,
			},
			targetPosition: CardPosition{
				stackId:  g.nowEventsDeck,
				position: i,
			},
		})
	}
	g.banishAgeAMilitaryCards()
}

func (g *TtaGame) banishAgeACards() {
	fmt.Println("TtaGame.banishAgeACards")
	csm := g.cardStackManager

	csm.processRequest(&BanishAllCardsInStackRequest{
		stackId: g.ageStacks[0],
	})
}

func (g *TtaGame) banishAgeAMilitaryCards() {
	fmt.Println("TtaGame.banishAgeAMilitaryCards")
	csm := g.cardStackManager

	csm.processRequest(&BanishAllCardsInStackRequest{
		stackId: g.miliDecks[0],
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

func (g *TtaGame) nextEventHappen() {
	csm := g.cardStackManager
	csm.processRequest(&MoveCardRequest{
		sourcePosition: CardPosition{
			stackId:  g.nowEventsDeck,
			position: csm.getStackSize(g.nowEventsDeck) - 1,
		},
		targetPosition: CardPosition{
			stackId:  g.pastEventsDeck,
			position: 0,
		},
	})

	if csm.getStackSize(g.nowEventsDeck) <= 0 {
		// Move future events to now
		for {
			if csm.getStackSize(g.futureEventsDeck) <= 0 {
				return
			}
			csm.processRequest(&SwapCardRequest{
				sourcePosition: CardPosition{
					stackId:  g.futureEventsDeck,
					position: 0,
				},
				targetPosition: CardPosition{
					stackId:  g.nowEventsDeck,
					position: 0,
				},
			})
		}
		cardCount := csm.getStackSize(g.nowEventsDeck)
		randomPerm := rand.Perm(cardCount)
		for i := 0; i < cardCount; i++ {
			csm.processRequest(&SwapCardRequest{
				sourcePosition: CardPosition{
					stackId:  g.nowEventsDeck,
					position: i,
				},
				targetPosition: CardPosition{
					stackId:  g.nowEventsDeck,
					position: randomPerm[i],
				},
			})
		}
	}
}

func (g *TtaGame) getCardOnGreatWheel(index int) *Card {
	if index < 0 || index >= 13 {
		return nil
	}
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

func (g *TtaGame) countPlayersWithPowerMoreThan(power int) int {
	result := 0
	for _, player := range g.players {
		if player.calcPower() > power {
			result++
		}
	}
	return result
}

func (g *TtaGame) processCivilMove(move *Move) (err error) {
	if move.FromPlayer != g.CurrentPlayer {
		return fmt.Errorf("Not current player.")
	}
	p := g.players[g.CurrentPlayer]
	switch move.MoveType {
	case MOVE_FETCH_CARD:
		if len(move.Data) != 1 {
			return fmt.Errorf("Invalid fetch command.")
		}
		index := move.Data[0]
		if !p.canTakeCardFromWheel(index) {
			return fmt.Errorf("Invalid fetch command.")
		}
		p.takeCardFromWheel(index)
	case MOVE_PLAY_CIVIL_CARD:
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
	case MOVE_INC_POP:
		if !p.canIncreasePop() {
			return fmt.Errorf("Invalid incpop command")
		}
		p.increasePop()
	case MOVE_BUILD:
		if len(move.Data) < 2 {
			return fmt.Errorf("Invalid build command.")
		}
		stack := move.Data[0]
		index := move.Data[1]
		if !p.canBuild(stack, index, 0) {
			return fmt.Errorf("Invalid build command")
		}
		p.build(stack, index, 0)
	case MOVE_BUILD_WONDER:
		if len(move.Data) < 1 {
			return fmt.Errorf("Invalid buildwonder command.")
		}
		step := move.Data[0]
		if !p.canBuildWonder(step, 0) {
			return fmt.Errorf("Invalid buildwonder command")
		}
		p.buildWonder(step, 0)
	case MOVE_UPGRADE:
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
	case MOVE_SPECIAL_ABILITY:
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
	case MOVE_END:
	}
	return nil
}

func (g *TtaGame) processDiscardMilitaryMove(move *Move) (err error) {
	return nil
}

func (g *TtaGame) pushStateHolder(stateHolder StateHolder) {
	g.StateStack = append(g.StateStack, stateHolder)
}

func (g *TtaGame) popStateHolder() StateHolder {
	result := g.StateStack[len(g.StateStack)-1]
	g.StateStack = g.StateStack[:len(g.StateStack)-1]
	return result
}

func (g *TtaGame) peekStateHolder() StateHolder {
	return g.StateStack[len(g.StateStack)-1]
}

func (g *TtaGame) Initialize() (err error) {
	for {
		stateHolder := g.peekStateHolder()
		if stateHolder.IsPending() {
			return
		} else {
			stateHolder.Resolve(nil)
		}
	}
}

func (g *TtaGame) TryResolveMove(move *Move) (err error) {
	if len(g.StateStack) == 0 {
		// Game end
		return
	}
	stateHolder := g.peekStateHolder()
	if !stateHolder.IsPending() {
		panic(stateHolder)
	}
	if legal, reason := stateHolder.IsMoveLegal(move); !legal {
		return fmt.Errorf(reason)
	}

	stateHolder.Resolve(move)
	for {
		if len(g.StateStack) == 0 {
			// Game end
			return
		}
		stateHolder = g.peekStateHolder()
		fmt.Println("loop", stateHolder)
		if stateHolder.IsPending() {
			return nil
		} else {
			stateHolder.Resolve(nil)
		}
	}
}
