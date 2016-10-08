package main

import (
	"fmt"
)

type ColonizeDetail struct {
	sacrificedStacks  []int
	sacrificedIndexes []int
	sacrificedCount   []int

	discardedMiliHands []int

	confirmed     bool
	giveUp        bool
	colonizePower int
}

type ColonizeStateHolder struct {
	base          BaseStateHolder
	currentWinner int
	currentPlayer int
	currentMax    int
	details       []ColonizeDetail
}

func NewColonizeStateHolder(game *TtaGame) *ColonizeStateHolder {
	return &ColonizeStateHolder{
		base: BaseStateHolder{
			game: game,
		},
		currentPlayer: game.CurrentPlayer,
		currentWinner: -1,
		details:       make([]ColonizeDetail, game.options.PlayerCount),
	}
}

func (h *ColonizeStateHolder) TryParseColonizeMove(move *Move) (detail ColonizeDetail, err error) {
	g := h.base.game
	detail = ColonizeDetail{}
	fmt.Println("Trying to parse colonize command: ", move.Data)
	index := 0
	if len(move.Data) <= index {
		return detail, fmt.Errorf("Invalid data")
	}

	l := move.Data[index]
	detail.sacrificedStacks = make([]int, l)
	detail.sacrificedIndexes = make([]int, l)
	detail.sacrificedCount = make([]int, l)
	for i := 0; i < l; i++ {
		index++
		if len(move.Data) <= index {
			return detail, fmt.Errorf("Invalid data")
		}
		detail.sacrificedStacks[i] = move.Data[index]
		index++
		if len(move.Data) <= index {
			return detail, fmt.Errorf("Invalid data")
		}
		detail.sacrificedIndexes[i] = move.Data[index]
		index++
		if len(move.Data) <= index {
			return detail, fmt.Errorf("Invalid data")
		}
		detail.sacrificedCount[i] = move.Data[index]
	}

	index++
	if len(move.Data) <= index {
		return detail, fmt.Errorf("Invalid data")
	}
	l = move.Data[index]
	detail.discardedMiliHands = make([]int, l)
	for i := 0; i < l; i++ {
		index++
		if len(move.Data) <= index {
			return detail, fmt.Errorf("Invalid data")
		}
		detail.discardedMiliHands[i] = move.Data[index]
	}

	detail.colonizePower = g.players[h.currentPlayer].calcColonizePower(&detail)
	if detail.colonizePower < 0 {
		return detail, fmt.Errorf("Invalid colonize command")
	}
	fmt.Println(detail)
	return detail, nil
}

func (h *ColonizeStateHolder) IsPending() bool {
	return true
}

func (h *ColonizeStateHolder) IsMoveLegal(move interface{}) (legal bool, reason string) {
	m := move.(*Move)
	if m.FromPlayer != h.currentPlayer {
		return false, "Not current player"
	}
	switch m.MoveType {
	case MOVE_COLONIZE:
		var detail ColonizeDetail
		var err error
		if detail, err = h.TryParseColonizeMove(m); err != nil {
			return false, err.Error()
		}

		if detail.colonizePower <= h.currentMax {
			return false, "Colonize power not enough"
		}
		return true, ""
	case MOVE_END:
		return true, ""
	}
	return false, ""
}

func (h *ColonizeStateHolder) Resolve(move interface{}) {
	fmt.Println("ColonizeStateHolder.Resolve")
	m := move.(*Move)
	g := h.base.game
	csm := g.cardStackManager
	switch m.MoveType {
	case MOVE_COLONIZE:
		var detail ColonizeDetail
		var err error
		if detail, err = h.TryParseColonizeMove(m); err != nil {
			return
		}

		h.currentMax = detail.colonizePower
		h.currentWinner = h.currentPlayer
		h.details[h.currentPlayer] = detail
		h.details[h.currentPlayer].confirmed = true
	case MOVE_END:
		h.details[h.currentPlayer].confirmed = true
		h.details[h.currentPlayer].giveUp = true
	}

	confirmedPlayers := 0
	giveUpPlayers := 0
	for _, detail := range h.details {
		if detail.confirmed {
			confirmedPlayers++
		}
		if detail.giveUp {
			giveUpPlayers++
		}
	}

	// All confirmed, at least playerCount - 1 players gave up
	if confirmedPlayers == g.options.PlayerCount &&
		giveUpPlayers >= g.options.PlayerCount-1 {
		g.popStateHolder()
		card := csm.cardStacks[g.pastEventsDeck][0]
		school := g.cardSchools[card.schoolId]
		if h.currentWinner >= 0 {
			csm.processRequest(&MoveCardRequest{
				sourcePosition: CardPosition{
					stackId:  g.pastEventsDeck,
					position: 0,
				},
				targetPosition: CardPosition{
					stackId:  g.players[h.currentWinner].stacks[COLONY],
					position: csm.getStackSize(g.players[h.currentWinner].stacks[COLONY]),
				},
			})
		}

		winner := g.players[h.currentWinner]
		winner.colonizeSuccess(&h.details[h.currentWinner])
		winner.gainColony(card)
		if school.hasTrait(TRAIT_DEVELOPED_TERRITORY) {
			winner.gainTech(school.actionBonus)
		} else if school.hasTrait(TRAIT_HISTORIC_TERRITORY) {
			winner.gainCulture(school.actionBonus)
		} else if school.hasTrait(TRAIT_INHABITED_TERRITORY) {
			for i := 0; i < school.actionBonus; i++ {
				if winner.canIncreasePop(0) {
					winner.increasePop(0)
				} else {
					break
				}
			}
		} else if school.hasTrait(TRAIT_STRATEGIC_TERRITORY) {
			winner.drawMiliCards(school.actionBonus)
		} else if school.hasTrait(TRAIT_VAST_TERRITORY) {
			winner.gainCrop(school.actionBonus)
		} else if school.hasTrait(TRAIT_WEALTHLY_TERRITORY) {
			winner.gainResource(school.actionBonus)
		}
		return
	} else {
		for {
			h.currentPlayer = (h.currentPlayer + 1) % g.options.PlayerCount
			if !h.details[h.currentPlayer].giveUp {
				break
			}
		}
	}
}
