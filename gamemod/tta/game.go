package main

import (
	"math/rand"
)

type UserStackId int

type TtaGame struct {
	cardStackManager   *CardStackUniversalManager
	globalTokenManager *TokenBankUniversalManager
	cardTokenManager   *TokenBankUniversalManager

	// All card schools
	cardSchools map[int]*CardSchool

	greatWheel []int // 13 stacks
	ageStacks  []int // 4 stacks by age
	players    []*PlayerBoard
}

func NewTta() (result *TtaGame) {
	game := &TtaGame{
		cardStackManager:   NewCardStackUniversalManager(),
		globalTokenManager: NewTokenBankUniversalManager(),
		cardTokenManager:   NewTokenBankUniversalManager(),
		greatWheel:         make([]int, 13),
		ageStacks:          make([]int, 4),
		players:            make([]*PlayerBoard, 2),
	}
	for i := 0; i < 2; i++ {
		game.players[i] = initPlayerBoard(game)
	}

	for i := 0; i < 13; i++ {
		game.greatWheel[i] = game.cardStackManager.newStack()
	}
	for i := 0; i < 4; i++ {
		game.ageStacks[i] = game.cardStackManager.newStack()
	}

	game.cardSchools = InitBasicCardSchools()

	game.initBasicCards()
	game.refillWheels()
	game.banishAgeACards()
	return game
}

func (g *TtaGame) checkDecay() {

}

func (g *TtaGame) initBasicCards() {
	// Fill all the civil cards
	csm := g.cardStackManager
	for id, school := range InitBasicCardSchools() {
		if school.schoolId == 1 ||
			school.schoolId == 5 ||
			school.schoolId == 9 ||
			school.schoolId == 13 ||
			school.schoolId == 25 {
			continue
		}
		for i := 0; i < school.cardCounts[0]; i++ {
			csm.processRequest(&AddCardToTopRequest{
				schoolId: id,
				stackId:  g.ageStacks[school.age],
			})
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
