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
	WONDER_UNDER_CONSTRUCTION
	WONDER
	TECH_SPECIAL
	HAND
	USER_STACK_SIZE
)

type PlayerBoard struct {
	stacks []int

	yellow int
	blue   int
}

func initPlayerBoard(game *TtaGame) (result *PlayerBoard) {
	csm := game.cardStackManager
	stacks := make([]int, USER_STACK_SIZE)
	for i := 0; i < USER_STACK_SIZE; i++ {
		stacks[i] = csm.newStack()
	}
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
	return &PlayerBoard{
		stacks: stacks,
		yellow: 15,
		blue:   15,
	}
}
