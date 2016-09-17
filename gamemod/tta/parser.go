package main

import (
	"fmt"
	"strconv"
	"strings"
)

func cardIdToIndex(p *PlayerBoard, cardId int) (stack int, index int, ok bool) {
	stackIds := []int{
		MILI_INFANTRY,
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
	}
	csm := p.game.cardStackManager
	for _, s := range stackIds {
		for i, c := range csm.cardStacks[p.stacks[s]] {
			if cardId == c.schoolId {
				return s, i, true
			}
		}
	}
	return 0, 0, false
}

func toAttachment(game *TtaGame, splitted []string) interface{} {
	result := make([]int, 0)
	for i := 2; i < len(splitted); i++ {
		r, err := strconv.Atoi(splitted[i])
		if err != nil {
			return nil
		}
		result = append(result, r)
	}

	if len(result) == 1 {
		return result[0]
	} else {
		return result
	}
}

func parseCommand(game *TtaGame, command string) {
	splitted := strings.Split(command, " ")
	if len(splitted) < 1 {
		fmt.Println("Unknown command")
	}

	switch splitted[0] {
	case "show", "s":
		if len(splitted) < 2 {
			fmt.Println("Unknown command")
		} else if splitted[1] == "0" {
			PrintUserBoard(game, game.players[0])
		} else if splitted[1] == "p" {
			PrintPublicArea(game)
		}
	case "end", "e":
		game.players[0].doProductionPhase()
		game.weedOut(2)
		game.refillWheels()
	case "fetch", "f":
		if len(splitted) < 2 {
			fmt.Println("Unknown command")
		} else {
			index, err := strconv.Atoi(splitted[1])

			if err != nil || index < 0 || index > 13 ||
				!game.players[0].canTakeCardFromWheel(index) {
				fmt.Println("Invliad fetch command ", err)
			} else {
				game.players[0].takeCardFromWheel(index)
			}
		}
	case "play", "p":
		if len(splitted) < 2 {
			fmt.Println("Unknown command")
		} else {
			index, err := strconv.Atoi(splitted[1])

			att := toAttachment(game, splitted)

			if err != nil || index < 0 || index > game.players[0].getHandSize() ||
				!game.players[0].canPlayHand(index, att) {
				fmt.Println("Invliad play command ", err)
			} else {
				game.players[0].playHand(index, att)
			}
		}
	case "incpop", "i":
		if !game.players[0].canIncreasePop() {
			fmt.Println("Invalid incpop command")
		} else {
			game.players[0].increasePop()
		}
	case "build", "b":
		if len(splitted) < 2 {
			fmt.Println("Unknown command")
		} else {
			cardId, err := strconv.Atoi(splitted[1])
			if err != nil {
				fmt.Println("Invalid build command")
				return
			}
			stack, index, ok := cardIdToIndex(game.players[0], cardId)
			if !ok || !game.players[0].canBuild(stack, index) {
				fmt.Println("Invalid build command")
			} else {
				game.players[0].build(stack, index)
			}
		}
	case "upgrade", "u":
		if len(splitted) < 3 {
			fmt.Println("Unknown command")
		} else {
			cardId1, err := strconv.Atoi(splitted[1])
			if err != nil {
				fmt.Println("Invalid upgrade command")
				return
			}
			cardId2, err := strconv.Atoi(splitted[2])
			if err != nil {
				fmt.Println("Invalid upgrade command")
				return
			}
			stack1, index1, ok := cardIdToIndex(game.players[0], cardId1)
			if !ok {
				fmt.Println("Invalid upgrade command")
				return
			}
			stack2, index2, ok := cardIdToIndex(game.players[0], cardId2)
			if !ok {
				fmt.Println("Invalid upgrade command")
				return
			}
			if stack1 != stack2 {
				fmt.Println("Invalid upgrade command")
				return
			}
			if !game.players[0].canUpgrade(stack1, index1, index2) {
				fmt.Println("Invalid build command")
			} else {
				fmt.Println("upgrade")
				game.players[0].upgrade(stack1, index1, index2)
			}
		}
	}
}