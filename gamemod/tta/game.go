package main

type UserStackId int

type TtaGame struct {
	cardStackManager *CardStackUniversalManager

	greatWheels []int // 14 stacks
	ageStacks   []int // 4 stacks by age
	players     []*PlayerBoard
}

func NewTta() (result *TtaGame) {
	game := &TtaGame{
		cardStackManager: NewCardStackUniversalManager(),
		greatWheels:      make([]int, 0),
		ageStacks:        make([]int, 0),
		players:          make([]*PlayerBoard, 2),
	}
	for i := 0; i < 2; i++ {
		game.players[i] = initPlayerBoard(game)
	}
	return game
}
