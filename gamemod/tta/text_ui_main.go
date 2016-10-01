package main

import (
	"bufio"
	"fmt"
	"os"
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

func ageToString(age int) string {
	switch age {
	case 0:
		return " A "
	case 1:
		return " I "
	case 2:
		return "II "
	case 3:
		return "III"
	}
	return "???"
}
func prepareBoard(width, height int) (result [][]rune) {
	backGround := []string{"a"}

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

func paintCounters(player *PlayerBoard) ([][]rune, int) {
	backGround := []string{
		"================================================",
		" CULTURE:                                       ",
		"    TECH:                                       ",
		"    PWER:                                       ",
	}
	result := toRunes(backGround)

	result = printUpon(result,
		toRunes([]string{strconv.Itoa(player.getCultureTotal()) + "/+" +
			strconv.Itoa(player.calcCultureInc())}), 14, 1)
	result = printUpon(result,
		toRunes([]string{strconv.Itoa(player.getTechTotal()) + "/+" +
			strconv.Itoa(player.calcTechInc())}), 14, 2)
	result = printUpon(result,
		toRunes([]string{strconv.Itoa(player.calcPower())}), 14, 3)

	return result, len(result)
}

func paintBlueBank(player *PlayerBoard) ([][]rune, int) {
	backGround := []string{
		"   o o   o o   o o           +--------------+",
		" 6     4     2         Blue  |FREE W.       |",
		"   o o   o o   o o o   bank  +--------------+",
	}

	result := toRunes(backGround)

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
		Point{x: 19, y: 2},
	}

	for i := 0; i < len(points); i++ {
		if player.specialTokenManager.getTokenCount(FREE_BLUE, TOKEN_BLUE) > i {
			result[points[i].y][points[i].x] = '*'
		} else {
			break
		}
	}

	freeWorkers := player.getFreeWorkers()
	if freeWorkers > 0 {
		result = printUpon(result,
			toRunes([]string{strconv.Itoa(freeWorkers)}), 38, 1)
	}

	return result, 3
}

func paintYellowBank(player *PlayerBoard) ([][]rune, int) {
	backGround := []string{
		" /8\\ /7\\ /6\\ /5\\ /4\\ /3\\ /--2--\\ /1\\ 0",
		" 6 o o o 4 o o o 3 o o o 2 o o o 1 o    Yellow",
		" +--7--+ +--5--+ +--4--+ +--3--+ +2+     bank",
	}
	bgRunes := toRunes(backGround)
	for i := 0; i < player.getFreeYellowTokens(); i++ {
		bgRunes[1][i*2+1] = '*'
	}

	// Happiness
	type Point struct {
		x, y int
	}

	points := []Point{
		Point{x: 37, y: 0},
		Point{x: 34, y: 0},
		Point{x: 28, y: 0},
		Point{x: 22, y: 0},
		Point{x: 18, y: 0},
		Point{x: 14, y: 0},
		Point{x: 10, y: 0},
		Point{x: 6, y: 0},
		Point{x: 2, y: 0},
	}

	for i := 0; i <= 8; i++ {
		if i <= player.calcHappiness() {
			bgRunes[points[i].y][points[i].x] = '@'
		} else if i <= player.getNeededHappiness() && i <= player.calcHappiness()+player.getFreeWorkers() {
			bgRunes[points[i].y][points[i].x] = 'W'
		} else if i <= player.getNeededHappiness() {
			bgRunes[points[i].y][points[i].x] = '!'
		}
	}
	return bgRunes, 3
}

func paintGovernmentAndLeader(player *PlayerBoard) ([][]rune, int) {
	csm := player.game.cardStackManager
	backGround := []string{
		"================================================",
		" GOVERNMENT:                                    ",
		" +--------------------+                         ",
		" |                    |                         ",
		" |                    |                         ",
		" +--------------------+                         ",
	}
	governmentCard := csm.getFirstCard(player.stacks[GOVERNMENT])
	govSchool := player.game.cardSchools[governmentCard.schoolId]
	result := printUpon(toRunes(backGround),
		toRunes([]string{ageToString(govSchool.age)}), 3, 3)
	result = printUpon(result,
		toRunes([]string{govSchool.schoolName}), 7, 3)
	result = printUpon(result,
		toRunes([]string{"[" + strconv.Itoa(govSchool.schoolId) + "]"}), 3, 4)
	result = printUpon(result,
		toRunes([]string{strconv.Itoa(player.getUsableWhiteTokens()) + "/" +
			strconv.Itoa(player.getUsableRedTokens())}), 10, 4)

	leaderCard := csm.getFirstCard(player.stacks[LEADER])
	if leaderCard != nil {
		leaderBg := []string{
			"LEADER:     ",
			"+---------------+",
			"|               |",
			"|               |",
			"+---------------+",
		}
		leaderSchool := player.game.cardSchools[leaderCard.schoolId]
		result = printUpon(result,
			toRunes(leaderBg), 30, 1)
		result = printUpon(result,
			toRunes([]string{leaderSchool.schoolName}), 32, 3)
		result = printUpon(result,
			toRunes([]string{"[" + strconv.Itoa(leaderSchool.schoolId) + "]"}), 32, 4)
	}
	return result, len(result)
}

func paintSpecialTechs(player *PlayerBoard) ([][]rune, int) {
	csm := player.game.cardStackManager
	need := false
	backGround := []string{
		"================================================",
		" SPEC TECHS:                                    ",
	}
	result := toRunes(backGround)
	count := 0
	for _, c := range csm.cardStacks[player.stacks[TECH_SPECIAL_CIVIL]] {
		need = true
		school := player.game.cardSchools[c.schoolId]
		result = printUpon(result,
			toRunes([]string{"[" + strconv.Itoa(school.schoolId) + "] " +
				school.schoolName}), 16, 1+count)
		count++
	}
	for _, c := range csm.cardStacks[player.stacks[TECH_SPECIAL_WARFARE]] {
		need = true
		school := player.game.cardSchools[c.schoolId]
		result = printUpon(result,
			toRunes([]string{"[" + strconv.Itoa(school.schoolId) + "] " +
				school.schoolName}), 16, 1+count)
		count++
	}
	for _, c := range csm.cardStacks[player.stacks[TECH_SPECIAL_COLONIZE]] {
		need = true
		school := player.game.cardSchools[c.schoolId]
		result = printUpon(result,
			toRunes([]string{"[" + strconv.Itoa(school.schoolId) + "] " +
				school.schoolName}), 16, 1+count)
		count++
	}
	for _, c := range csm.cardStacks[player.stacks[TECH_SPECIAL_CONSTRUCTION]] {
		need = true
		school := player.game.cardSchools[c.schoolId]
		result = printUpon(result,
			toRunes([]string{"[" + strconv.Itoa(school.schoolId) + "] " +
				school.schoolName}), 16, 1+count)
		count++
	}
	if need {
		return result, len(result)
	} else {
		return make([][]rune, 0), 0
	}
}

func paintWonders(player *PlayerBoard) ([][]rune, int) {
	csm := player.game.cardStackManager
	need := false
	backGround := []string{
		"================================================",
		" WONDERS:                                       ",
	}
	result := toRunes(backGround)
	wonderCard := csm.getFirstCard(player.stacks[WONDER_NOT_COMPLETED])
	if wonderCard != nil {
		need = true
		constructingWonderBg := []string{
			" +-------------------+                          ",
			" |                   |                          ",
			" |                   |                          ",
			" |                   |                          ",
			" +-------------------+                          ",
		}
		result = printUpon(result,
			toRunes(constructingWonderBg), 0, 2)

		wonderSchool := player.game.cardSchools[wonderCard.schoolId]
		result = printUpon(result,
			toRunes([]string{wonderSchool.schoolName}), 8, 3)
		result = printUpon(result,
			toRunes([]string{"[" + strconv.Itoa(wonderSchool.schoolId) + "]"}), 3, 3)
		for i, s := range wonderSchool.wonderBuildCosts {
			result = printUpon(result,
				toRunes([]string{strconv.Itoa(s)}), 4+i*2, 4)
		}
		for i := 0; i < player.getBlueTokensOnCurrentWonder(); i++ {
			result = printUpon(result,
				toRunes([]string{"*"}), 4+i*2, 5)
		}
	}

	for i, wonderCard := range csm.cardStacks[player.stacks[WONDER_COMPLETED]] {
		need = true
		wonderSchool := player.game.cardSchools[wonderCard.schoolId]
		result = printUpon(result,
			toRunes([]string{"[" + strconv.Itoa(wonderSchool.schoolId) + "] " +
				wonderSchool.schoolName}), 25, 1+i)
	}
	if need {
		return result, len(result)
	} else {
		return make([][]rune, 0), 0
	}
}

func paintHands(player *PlayerBoard) ([][]rune, int) {
	csm := player.game.cardStackManager
	need := false
	backGround := []string{
		"================================================",
		" HANDS:                                       ",
	}
	result := toRunes(backGround)
	h := 1
	for _, handCard := range csm.cardStacks[player.stacks[HAND]] {
		need = true
		handSchool := player.game.cardSchools[handCard.schoolId]
		result = printUpon(result,
			toRunes([]string{"[" + ageToString(handSchool.age) +
				strconv.Itoa(handSchool.schoolId) + "] " +
				handSchool.schoolName}), 15, h)
		h += 1
	}
	for _, handCard := range csm.cardStacks[player.stacks[MILI_HAND]] {
		need = true
		handSchool := player.game.cardSchools[handCard.schoolId]
		result = printUpon(result,
			toRunes([]string{"[" + ageToString(handSchool.age) +
				strconv.Itoa(handSchool.schoolId) + "] " +
				handSchool.schoolName}), 15, h)
		h += 1
	}
	if need {
		return result, len(result)
	} else {
		return make([][]rune, 0), 0
	}
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
	case 6, 7, 8, 9:
		costStr = "**"
	case 10, 11, 12, 13:
		costStr = "***"
	}
	result := printUpon(toRunes(backGround),
		toRunes([]string{strconv.Itoa(index - 1)}), 1, 0)
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

	result := printUpon(toRunes(backGround),
		toRunes([]string{ageToString(age)}), 2, 1)
	result = printUpon(result,
		toRunes([]string{strconv.Itoa(schoolId)}), 6, 1)
	result = printUpon(result,
		toRunes([]string{name}), 1, 2)
	if token1 > 0 {
		result = printUpon(result,
			toRunes([]string{strconv.Itoa(token1)}), 2, 3)
	}
	if token2 > 0 {
		result = printUpon(result,
			toRunes([]string{strconv.Itoa(token2)}), 6, 3)
	}

	return result, 5
}

func paintStructures(game *TtaGame, player *PlayerBoard) ([][]rune, int) {
	csm := game.cardStackManager
	schools := InitBasicCardSchools()
	maxDepth := 0
	maxWidth := 0

	toPaint := []int{
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

	for _, p := range toPaint {
		stack := csm.cardStacks[player.stacks[p]]
		if len(stack) > 0 {
			maxWidth++
			if len(stack) > maxDepth {
				maxDepth = len(stack)
			}
		}
	}

	x := 0
	result := make([][]rune, 0)
	for _, school := range toPaint {
		stack := csm.cardStacks[player.stacks[school]]
		if len(stack) > 0 {
			for y, card := range stack {
				cardSchool := schools[card.schoolId]
				r, _ := paintSingleStructure(
					cardSchool.age,
					cardSchool.schoolId,
					cardSchool.shortName,
					game.cardTokenManager.getTokenCount(card.id, TOKEN_YELLOW),
					game.cardTokenManager.getTokenCount(card.id, TOKEN_BLUE))
				result = printUpon(result, r, x*9, (maxDepth-y-1)*4)
			}
			x += 1
		}
	}

	return result, maxDepth*4 + 1
}

func paintEventDecks(game *TtaGame) ([][]rune, int) {
	csm := game.cardStackManager
	backGround := []string{
		"+------+    +------+",
		"|      | -> |      |",
		"+------+    +------+",
		"+------------------+",
		"|                  |",
		"+------------------+",
	}
	result := toRunes(backGround)
	if csm.getStackSize(game.futureEventsDeck) > 0 {
		eventCard := csm.getFirstCard(game.futureEventsDeck)
		eventSchool := game.cardSchools[eventCard.schoolId]
		result = printUpon(result,
			toRunes([]string{ageToString(eventSchool.age)}), 1, 1)
		result = printUpon(result, toRunes([]string{strconv.Itoa(
			csm.getStackSize(game.futureEventsDeck))}), 5, 1)
	}

	if csm.getStackSize(game.nowEventsDeck) > 0 {
		eventCard := csm.getFirstCard(game.nowEventsDeck)
		eventSchool := game.cardSchools[eventCard.schoolId]
		result = printUpon(result,
			toRunes([]string{ageToString(eventSchool.age)}), 13, 1)
		result = printUpon(result, toRunes([]string{strconv.Itoa(
			csm.getStackSize(game.nowEventsDeck))}), 17, 1)
	}

	if csm.getStackSize(game.pastEventsDeck) > 0 {
		eventCard := csm.getFirstCard(game.pastEventsDeck)
		eventSchool := game.cardSchools[eventCard.schoolId]
		result = printUpon(result,
			toRunes([]string{"[" + ageToString(eventSchool.age) +
				strconv.Itoa(eventSchool.schoolId) + "]" +
				eventSchool.schoolName}), 1, 4)
	}
	return result, len(result)
}

func PrintGreatWheels(game *TtaGame) ([][]rune, int) {
	csm := game.cardStackManager
	schools := InitBasicCardSchools()

	result := make([][]rune, 0)
	for i := 0; i < 13; i++ {
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

	events, _ := paintEventDecks(game)
	result = printUpon(result, events, 40, 0)
	return result, len(result)
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
	runes := prepareBoard(80, 20)
	var h int
	h = 1
	counters, height := paintCounters(player)
	runes = printUpon(runes, counters, 1, h)
	h += height
	structures, height := paintStructures(game, player)
	runes = printUpon(runes, structures, 1, h)
	h += height
	blue, height := paintBlueBank(player)
	runes = printUpon(runes, blue, 1, h)
	h += height
	yellow, height := paintYellowBank(player)
	runes = printUpon(runes, yellow, 1, h)
	h += height
	govLeader, height := paintGovernmentAndLeader(player)
	runes = printUpon(runes, govLeader, 1, h)

	h += height
	specTech, height := paintSpecialTechs(player)
	runes = printUpon(runes, specTech, 1, h)
	h += height
	wonders, height := paintWonders(player)
	runes = printUpon(runes, wonders, 1, h)
	h += height
	hands, height := paintHands(player)
	runes = printUpon(runes, hands, 1, h)
	PrintAll(toStrings(runes))
}

func GetCurrentPendingPlayer(game *TtaGame) int {
	stateHolder := game.peekStateHolder()
	switch h := stateHolder.(type) {
	case *CivilStateHolder:
		return game.CurrentPlayer
	case *DiscardMilitaryCardsStateHolder:
		return h.player
	case *PoliticalStateHolder:
		return game.CurrentPlayer
	case *DefenseAggressionStateHolder:
		return h.player
	case *LosePopulationStateHolder:
		for i := 0; i < len(game.players); i++ {
			if h.popToLose[i] <= game.players[i].getFreeWorkers() {
				return i
			}
		}
		return -1
	case *PlunderStateHolder:
		return game.CurrentPlayer
	case *RaidStateHolder:
		return game.CurrentPlayer
	default:
		return -1
	}
}

func GetCurrentPendingPlayerForPlayAttachment(game *TtaGame) int {
	stateHolder := game.peekStateHolder()
	switch h := stateHolder.(type) {
	case *RaidStateHolder:
		return h.targetPlayer
	default:
		return GetCurrentPendingPlayer(game)
	}
}

func PrintCurrentState(game *TtaGame) {
	stateHolder := game.peekStateHolder()
	switch h := stateHolder.(type) {
	case *CivilStateHolder:
		fmt.Println("[PENDING]Waiting for player ", GetCurrentPendingPlayer(game), " for civil actions")
	case *DiscardMilitaryCardsStateHolder:
		fmt.Println("[PENDING]Waiting for player", GetCurrentPendingPlayer(game), "for discarding military cards")
	case *PoliticalStateHolder:
		fmt.Println("[PENDING]Waiting for player ", GetCurrentPendingPlayer(game), " for political actions")
	case *DefenseAggressionStateHolder:
		fmt.Println("[PENDING]Waiting for player ", GetCurrentPendingPlayer(game), " for playing defensive cards")
		fmt.Println(" Aggression under player ", h.sourcePlayer, ", strength is ", h.sourcePower)
	case *LosePopulationStateHolder:
		fmt.Println("[PENDING]Waiting for player ", GetCurrentPendingPlayer(game), " to choose population to lose[use d command]")
	case *PlunderStateHolder:
		fmt.Println("[PENDING]Waiting for player ", GetCurrentPendingPlayer(game), " for choosing crop/resource to plunder[use o command]")
	case *RaidStateHolder:
		fmt.Println("[PENDING]Waiting for player ", GetCurrentPendingPlayer(game), " for choosing structure of player ", h.targetPlayer, "to destroy[use d command]")
	default:
		fmt.Println(h)
	}
}

func PrintGame(game *TtaGame) {
	PrintUserBoard(game, game.players[0])
	PrintPublicArea(game)
}

func main() {
	options := &TtaGameOptions{
		PlayerCount: 2,
	}
	game := NewTta(options)
	PrintGame(game)
	for {
		bio := bufio.NewReader(os.Stdin)
		line, _, err := bio.ReadLine()

		if err != nil {
			continue
		}

		parseCommand(game, string(line))
	}
}
