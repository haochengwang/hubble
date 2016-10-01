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

func getIthHandCardSchool(game *TtaGame, index int) *CardSchool {
	csm := game.cardStackManager
	if index < 0 || index >= csm.getStackSize(game.players[GetCurrentPendingPlayer(game)].stacks[HAND]) {
		return nil
	}
	card := csm.cardStacks[game.players[GetCurrentPendingPlayer(game)].stacks[HAND]][index]
	return game.cardSchools[card.schoolId]
}

func toAttachment(game *TtaGame, start int, splitted []string) []int {
	result := make([]int, 0)
	for i := start; i < len(splitted); i++ {
		r, err := strconv.Atoi(splitted[i])
		if err != nil {
			return nil
		}
		result = append(result, r)
	}

	return result
}

func toPlayAttachment(game *TtaGame, splitted []string) []int {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	result := make([]int, 0)
	c := -1
	for i := 1; i < len(splitted); i++ {
		r, err := strconv.Atoi(splitted[i])
		if err != nil {
			return nil
		}
		if i == 1 {
			c = r
		} else {
			result = append(result, r)
		}
	}

	cp := GetCurrentPendingPlayer(game)
	if len(result) == 2 && getIthHandCardSchool(game, c).hasType(CARDTYPE_ACTION_EFFICIENT_UPGRADE) {
		stack1, index1, ok := cardIdToIndex(game.players[cp], result[0])
		if !ok {
			return nil
		}
		stack2, index2, ok := cardIdToIndex(game.players[cp], result[1])
		if !ok {
			return nil
		}
		if stack1 != stack2 {
			return nil
		}

		return []int{stack1, index1, index2}
	} else if len(result) == 1 && getIthHandCardSchool(game, c).hasType(CARDTYPE_ACTION_RICH_LAND) {
		stack, index, ok := cardIdToIndex(game.players[cp], result[0])
		if !ok {
			return nil
		}

		return []int{stack, index}
	} else if len(result) == 2 && getIthHandCardSchool(game, c).hasType(CARDTYPE_ACTION_RICH_LAND) {
		stack1, index1, ok := cardIdToIndex(game.players[cp], result[0])
		if !ok {
			return nil
		}
		stack2, index2, ok := cardIdToIndex(game.players[cp], result[1])
		if !ok {
			return nil
		}
		if stack1 != stack2 {
			return nil
		}

		return []int{stack1, index1, index2}
	} else if len(result) == 1 && getIthHandCardSchool(game, c).hasType(CARDTYPE_ACTION_URBAN_GROWTH) {
		stack, index, ok := cardIdToIndex(game.players[cp], result[0])
		if !ok {
			return nil
		}

		return []int{stack, index}
	} else if len(result) == 2 && getIthHandCardSchool(game, c).hasType(CARDTYPE_ACTION_URBAN_GROWTH) {
		stack1, index1, ok := cardIdToIndex(game.players[cp], result[0])
		if !ok {
			return nil
		}
		stack2, index2, ok := cardIdToIndex(game.players[cp], result[1])
		if !ok {
			return nil
		}
		if stack1 != stack2 {
			return nil
		}

		return []int{stack1, index1, index2}
	} else {
		return result
	}
}

func parseCommand(game *TtaGame, command string) {
	splitted := strings.Split(command, " ")
	if len(splitted) < 1 {
		fmt.Println("Unknown command")
	}

	cp := GetCurrentPendingPlayer(game)
	switch splitted[0] {
	case "show", "s":
		if len(splitted) < 2 {
			fmt.Println("Unknown command")
		} else if splitted[1] == "0" {
			PrintUserBoard(game, game.players[0])
		} else if splitted[1] == "1" && game.options.PlayerCount >= 2 {
			PrintUserBoard(game, game.players[1])
		} else if splitted[1] == "2" && game.options.PlayerCount >= 3 {
			PrintUserBoard(game, game.players[2])
		} else if splitted[1] == "3" && game.options.PlayerCount >= 4 {
			PrintUserBoard(game, game.players[3])
		} else if splitted[1] == "p" {
			PrintPublicArea(game)
			PrintCurrentState(game)
		} else {
			fmt.Println("Unknown command")
		}
	case "end", "e":
		fmt.Println("OK")
		game.TryResolveMove(&Move{
			FromPlayer: cp,
			MoveType:   MOVE_END,
		})
		/*game.players[0].doProductionPhase()
		game.players[0].clearupTurn()
		game.players[0].drawMiliCards(2)
		game.weedOut(3)
		game.refillWheels()*/
	case "fetch", "f":
		if len(splitted) < 2 {
			fmt.Println("Unknown command")
		} else {
			index, err := strconv.Atoi(splitted[1])
			err = game.TryResolveMove(&Move{
				FromPlayer: cp,
				MoveType:   MOVE_FETCH_CARD,
				Data:       []int{index},
			})
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("OK")
			}
		}
	case "play", "p":
		if len(splitted) < 2 {
			fmt.Println("Unknown command")
		} else {
			index, err := strconv.Atoi(splitted[1])
			att := toPlayAttachment(game, splitted)
			err = game.TryResolveMove(&Move{
				FromPlayer: cp,
				MoveType:   MOVE_PLAY_CIVIL_CARD,
				Data:       append([]int{index}, att...),
			})
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("OK")
			}
		}
	case "playmili", "pm":
		if len(splitted) < 2 {
			fmt.Println("Unknown command")
		} else {
			index, err := strconv.Atoi(splitted[1])
			att := toAttachment(game, 2, splitted)
			err = game.TryResolveMove(&Move{
				FromPlayer: cp,
				MoveType:   MOVE_PLAY_MILITARY_CARD,
				Data:       append([]int{index}, att...),
			})
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("OK")
			}
		}
	case "incpop", "i":
		err := game.TryResolveMove(&Move{
			FromPlayer: cp,
			MoveType:   MOVE_INC_POP,
			Data:       []int{},
		})
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("OK")
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
			stack, index, ok := cardIdToIndex(game.players[cp], cardId)
			if !ok {
				fmt.Println("Invalid build command")
				return
			}
			err = game.TryResolveMove(&Move{
				FromPlayer: cp,
				MoveType:   MOVE_BUILD,
				Data:       append([]int{stack, index}),
			})
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("OK")
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
			stack1, index1, ok := cardIdToIndex(game.players[cp], cardId1)
			if !ok {
				fmt.Println("Invalid upgrade command")
				return
			}
			stack2, index2, ok := cardIdToIndex(game.players[cp], cardId2)
			if !ok {
				fmt.Println("Invalid upgrade command")
				return
			}
			if stack1 != stack2 {
				fmt.Println("Invalid upgrade command")
				return
			}
			err = game.TryResolveMove(&Move{
				FromPlayer: cp,
				MoveType:   MOVE_UPGRADE,
				Data:       append([]int{stack1, index1, index2}),
			})
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("OK")
			}
		}
	case "disband", "d":
		if len(splitted) < 2 {
			fmt.Println("Unknown command")
		} else {
			cardId, err := strconv.Atoi(splitted[1])
			if err != nil {
				fmt.Println("Invalid disband command")
				return
			}
			stack, index, ok := cardIdToIndex(game.players[cp], cardId)
			if !ok {
				fmt.Println("Invalid disband command")
				return
			}
			err = game.TryResolveMove(&Move{
				FromPlayer: cp,
				MoveType:   MOVE_DISBAND,
				Data:       append([]int{stack, index}),
			})
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("OK")
			}
		}
	case "buildwonder", "bw":
		step := 1
		if len(splitted) >= 2 {
			var err error
			step, err = strconv.Atoi(splitted[1])
			if err != nil {
				fmt.Println("Invalid buildwonder command")
				return
			}
		}
		err := game.TryResolveMove(&Move{
			FromPlayer: cp,
			MoveType:   MOVE_BUILD_WONDER,
			Data:       append([]int{step}),
		})
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("OK")
		}
	case "specialability", "sa":
		if len(splitted) < 2 {
			fmt.Println("Unknown command")
		} else {
			sa, err := strconv.Atoi(splitted[1])

			att := toAttachment(game, 2, splitted)

			// Ugly, temporary code
			if sa == 74 { // Hammurabi
				sa = 14
			} else if sa == 66 { // Ocean Liner Service
				sa = 5
			} else if sa == 72 { // Homer
				sa = 12
			} else if sa == 82 { // Barbarosa
				sa = 22
			} else if sa == 86 { // Robespierre
				sa = 26
			} else if sa == 93 { // Churchill
				sa = 33
			}

			err = game.TryResolveMove(&Move{
				FromPlayer: cp,
				MoveType:   MOVE_SPECIAL_ABILITY,
				Data:       append([]int{sa}, att...),
			})
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("OK")
			}
		}
	case "discardmili", "dm":
		if len(splitted) < 2 {
			fmt.Println("Unknown command")
		} else {
			att := toAttachment(game, 1, splitted)

			err := game.TryResolveMove(&Move{
				FromPlayer: cp,
				MoveType:   MOVE_DISCARD_MILITARY_CARDS,
				Data:       att,
			})
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("OK")
			}
		}
	case "operation", "o":
		att := toAttachment(game, 1, splitted)

		fmt.Println(att)
		err := game.TryResolveMove(&Move{
			FromPlayer: cp,
			MoveType:   MOVE_GENERAL_OP,
			Data:       att,
		})
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("OK")
		}
	default:
		fmt.Println("Unknown command")
	}
}
