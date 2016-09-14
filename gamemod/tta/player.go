package main

const (
	MILI_INFANTRY int = iota
	MILI_CAVALRY
	MILI_ARTILERY
	FARM
	MINE
	URBAN_TEMPLE
	URBAN_LAB
	URBAN_ARENA
	URBAN_LIBRARY
	URBAN_THEATER
	LEADER
	WONDER_NOT_COMPLETED
	WONDER_COMPLETED
	TECH_SPECIAL
	HAND
	USER_STACK_SIZE
)

const (
	FREE_YELLOW = iota
	FREE_BLUE
	FREE_WORKER
	FARM_A
	FARM_I
	FARM_II
	FARM_III
	MINE_A
	MINE_I
	MINE_II
	MINE_III
	WONDER
	POLITICAL_UNUSED
	POLITICAL_USED

	CULTURE_COUNTER
	TEC_COUNTER
)

const (
	TOKEN_YELLOW = iota
	TOKEN_BLUE
	TOKEN_WHITE
	TOKEN_RED
)

type PlayerBoard struct {
	stacks           []int
	handTokenManager *TokenBankUniversalManager
	userTokenManager *TokenBankUniversalManager
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

	// Prepare token bank manager
	handTokenManager := NewTokenBankUniversalManager()
	userTokenManager := NewTokenBankUniversalManager()
	userTokenManager.setTokenCount(FARM_A, TOKEN_YELLOW, 2)
	userTokenManager.setTokenCount(MINE_A, TOKEN_YELLOW, 2)
	userTokenManager.setTokenCount(FREE_WORKER, TOKEN_YELLOW, 1)
	userTokenManager.setTokenCount(FREE_YELLOW, TOKEN_YELLOW, 15)
	userTokenManager.setTokenCount(FREE_BLUE, TOKEN_BLUE, 15)

	// Finish
	return &PlayerBoard{
		stacks:           stacks,
		handTokenManager: handTokenManager,
		userTokenManager: userTokenManager,
	}
}
