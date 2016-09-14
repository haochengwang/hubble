package main

import (
	"fmt"
	"strconv"
)

// Unicode is not considered

func toRunes(str []string) (result [][]rune) {
	result = make([][]rune, len(str))
	for i := 0; i < len(str); i++ {
		result[i] = []rune(str[i])
	}
	return result
}

func toStrings(runes [][]rune) (result []string) {
	result = make([]string, len(runes))
	for i := 0; i < len(runes); i++ {
		result[i] = string(runes[i])
	}
	return result
}

func prepareBoard(width, height int) (result [][]rune) {
	backGround := []string{
		"+---------------------------------------------------+",
		"|                                                   |",
		"|                                                   |",
		"|                                                   |",
		"|                                                   |",
		"|                                                   |",
		"|                                                   |",
		"|                                                   |",
		"|                                       [1]         |",
		"|                                                   |",
		"|                                                   |",
		"|                                                   |",
		"+---------+  +----------+                           |",
		"|Despotism|  | Aristotle|                           |",
		"|         |  |          |                           |",
		"|  [4]  2 |  |          |                           |",
		"+---------+  +----------+                           |",
		"+---------+                                         |",
		"|  Pyramid|                                         |",
		"|         |                                         |",
		"|   * 2 1 |                                         |",
		"+---------+                                         |",
		"|                                                   |",
		"|                                                   |",
		"|                                                   |",
		"+---------------------------------------------------+",
	}

	return toRunes(backGround)
}

func printUpon(source [][]rune, str [][]rune, x, y int) [][]rune {
	for i := 0; i < len(str); i++ {
		for {
			if i+y < len(source) {
				break
			}
			source = append(source, []rune{})
		}
		for j := 0; j < len(str[i]); j++ {
			for {
				if j+x < len(source[i+y]) {
					break
				}
				source[i+y] = append(source[i+y], ' ')
			}
			source[i+y][j+x] = str[i][j]
		}
	}
	return source
}

func paintBlueBank(player *PlayerBoard) ([][]rune, int) {
	backGround := []string{
		"   o o   o o   o o  ",
		" 6     4     2         Blue",
		"   o o   o o   o o o   bank",
	}

	bgRunes := toRunes(backGround)

	type Point struct {
		x, y int
	}

	points := []Point{
		Point{x: 1, y: 1},
		Point{x: 3, y: 2},
		Point{x: 3, y: 0},
		Point{x: 5, y: 2},
		Point{x: 5, y: 0},
		Point{x: 7, y: 1},
		Point{x: 9, y: 2},
		Point{x: 9, y: 0},
		Point{x: 11, y: 2},
		Point{x: 11, y: 0},
		Point{x: 13, y: 1},
		Point{x: 15, y: 2},
		Point{x: 15, y: 0},
		Point{x: 17, y: 2},
		Point{x: 17, y: 0},
		Point{x: 19, y: 0},
	}

	for i := 0; i < len(points); i++ {
		if player.blue > i {
			bgRunes[points[i].y][points[i].x] = '*'
		} else {
			break
		}
	}

	return bgRunes, 3
}

func paintYellowBank(player *PlayerBoard) ([][]rune, int) {
	backGround := []string{
		" /8\\ /7\\ /6\\ /5\\ /4\\ /3\\ /--2--\\ /1\\ 0",
		" 5 o o o 4 o o o 3 o o o 2 o o o 1 o    Yellow bank",
		" +--7--+ +--5--+ +--4--+ +--3--+ +2+",
	}
	bgRunes := toRunes(backGround)
	for i := 0; i < player.yellow; i++ {
		bgRunes[1][i*2+1] = '*'
	}
	return bgRunes, 3
}

func paintCardOnWheel(index, age, id int, name string) ([][]rune, int) {
	backGround := []string{
		"       [                            ]",
	}
	ageStr := "???"
	switch age {
	case 0:
		ageStr = "A"
	case 1:
		ageStr = "I"
	case 2:
		ageStr = "II"
	case 3:
		ageStr = "III"
	}

	costStr := "???"
	switch index {
	case 1, 2, 3, 4, 5:
		costStr = "*"
	case 6, 7, 8, 9, 10:
		costStr = "**"
	case 11, 12, 13, 14:
		costStr = "***"
	}
	result := printUpon(toRunes(backGround),
		toRunes([]string{strconv.Itoa(index)}), 1, 0)
	result = printUpon(result,
		toRunes([]string{costStr}), 4, 0)
	result = printUpon(result,
		toRunes([]string{ageStr}), 8, 0)
	result = printUpon(result,
		toRunes([]string{strconv.Itoa(id)}), 11, 0)
	result = printUpon(result,
		toRunes([]string{name}), 15, 0)

	return result, 5
}

func paintSingleStructure(age, schoolId int, name string, token1, token2 int) ([][]rune, int) {
	backGround := []string{
		"+--------+",
		"|        |",
		"|        |",
		"|        |",
		"+--------+",
	}

	ageStr := "???"
	switch age {
	case 0:
		ageStr = "A"
	case 1:
		ageStr = "I"
	case 2:
		ageStr = "II"
	case 3:
		ageStr = "III"
	}
	result := printUpon(toRunes(backGround),
		toRunes([]string{ageStr}), 2, 1)
	result = printUpon(result,
		toRunes([]string{strconv.Itoa(schoolId)}), 6, 1)
	result = printUpon(result,
		toRunes([]string{name}), 1, 2)

	return result, 5
}

func paintStructures(game *TtaGame, player *PlayerBoard) ([][]rune, int) {
	csm := game.cardStackManager
	schools := InitBasicCardSchools()
	//	maxDepth := 1

	infantryStack := csm.cardStacks[player.stacks[MILI_INFANTRY]]
	farmStack := csm.cardStacks[player.stacks[FARM]]
	mineStack := csm.cardStacks[player.stacks[MINE]]
	templeStack := csm.cardStacks[player.stacks[URBAN_TEMPLE]]
	labStack := csm.cardStacks[player.stacks[URBAN_LAB]]
	cardSchool := schools[infantryStack[0].schoolId]

	result, _ := paintSingleStructure(cardSchool.age, cardSchool.schoolId, cardSchool.shortName, 0, 0)

	cardSchool = schools[farmStack[0].schoolId]
	fmt.Println(farmStack[0].schoolId)
	r2, _ := paintSingleStructure(cardSchool.age, cardSchool.schoolId, cardSchool.shortName, 0, 0)
	result = printUpon(result, r2, 10, 0)

	cardSchool = schools[mineStack[0].schoolId]
	r3, _ := paintSingleStructure(cardSchool.age, cardSchool.schoolId, cardSchool.shortName, 0, 0)
	result = printUpon(result, r3, 20, 0)

	cardSchool = schools[templeStack[0].schoolId]
	r4, _ := paintSingleStructure(cardSchool.age, cardSchool.schoolId, cardSchool.shortName, 0, 0)
	result = printUpon(result, r4, 30, 0)

	cardSchool = schools[labStack[0].schoolId]
	r5, _ := paintSingleStructure(cardSchool.age, cardSchool.schoolId, cardSchool.shortName, 0, 0)
	result = printUpon(result, r5, 40, 0)

	return result, 5
}

func PrintGreatWheels(game *TtaGame) ([][]rune, int) {
	csm := game.cardStackManager
	schools := InitBasicCardSchools()

	result := make([][]rune, 0)
	for i := 0; i < 14; i++ {
		stack := csm.cardStacks[game.greatWheel[i]]
		if len(stack) > 0 {
			cardSchool := schools[stack[0].schoolId]
			r, _ := paintCardOnWheel(i+1, cardSchool.age, cardSchool.schoolId, cardSchool.schoolName)
			result = printUpon(result, r, 0, i)
		} else {
			r, _ := paintCardOnWheel(i+1, -1, 0, "")
			result = printUpon(result, r, 0, i)
		}

	}
	return result, 1
}

func PrintAll(str []string) {
	for _, s := range str {
		fmt.Println(s)
	}
}

func PrintPublicArea(game *TtaGame) {
	runes, _ := PrintGreatWheels(game)
	PrintAll(toStrings(runes))
}

func PrintUserBoard(game *TtaGame, player *PlayerBoard) {
	//csm := game.cardStackManager
	//schools := InitBasicCardSchools()

	//infantryStack := csm.cardStacks[player.stacks[MILI_INFANTRY]]
	//fmt.Println(schools[infantryStack[0].schoolId].schoolName)

	runes := prepareBoard(80, 20)
	var h int
	h = 1
	structures, height := paintStructures(game, player)
	printUpon(runes, structures, 1, h)
	h += height
	blue, height := paintBlueBank(player)
	printUpon(runes, blue, 1, h)
	h += height
	yellow, height := paintYellowBank(player)
	printUpon(runes, yellow, 1, 9)
	PrintAll(toStrings(runes))
}

func PrintGame(game *TtaGame) {
	PrintUserBoard(game, game.players[0])
	PrintPublicArea(game)
}

func main() {
	game := NewTta()
	for {
		PrintGame(game)
		var a string
		fmt.Scanln(&a)

		game.weedOut(1)
		game.refillWheels()
	}
}
