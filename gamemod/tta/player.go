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
		stackId:  stacks[MILI_INFANTRY],
		schoolId: 25,
	})
	csm.processRequest(&AddCardRequest{
		stackId:  stacks[FARM],
		schoolId: 1,
	})
	csm.processRequest(&AddCardRequest{
		stackId:  stacks[MINE],
		schoolId: 5,
	})
	csm.processRequest(&AddCardRequest{
		stackId:  stacks[URBAN_TEMPLE],
		schoolId: 13,
	})
	csm.processRequest(&AddCardRequest{
		stackId:  stacks[URBAN_LAB],
		schoolId: 9,
	})
	return &PlayerBoard{
		stacks: stacks,
		yellow: 10,
		blue:   15,
	}
}
