package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseCommand(game *TtaGame, command string) {
	splitted := strings.Split(command, " ")
	if len(splitted) < 1 {
		fmt.Println("Unknown command")
	}

	fmt.Println(command)
	fmt.Println(splitted)
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

			if err != nil || index < 0 || index > game.players[0].getHandSize() {
				//!game.players[0].canPlayHand(index) {
				fmt.Println("Invliad play command ", err)
			} else {
				game.players[0].playHand(index, nil)
			}
		}
	}
}
