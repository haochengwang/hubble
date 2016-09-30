package main

import (
	"fmt"
)

type StateHolder interface {
	IsPending() bool
	IsMoveLegal(move interface{}) (legal bool, reason string)
	Resolve(move interface{})
}

type BaseStateHolder struct {
	game *TtaGame
}

type TurnStartStateHolder struct {
	base BaseStateHolder
}

func (h *TurnStartStateHolder) IsPending() bool {
	return false
}

func (h *TurnStartStateHolder) IsMoveLegal(move interface{}) (legal bool, reason string) {
	return true, ""
}

func (h *TurnStartStateHolder) Resolve(move interface{}) {
	fmt.Println("start resolve")
	g := h.base.game

	g.popStateHolder()
	g.pushStateHolder(&TurnEndStateHolder{
		base: BaseStateHolder {
			game: g,
		},
	})
	g.pushStateHolder(&DiscardMilitaryCardsStateHolder{
		base: BaseStateHolder {
			game: g,
		},
		player:    g.CurrentPlayer,
		toMaxHand: true,
	})
	g.pushStateHolder(&ProductionPhaseStateHolder{
		base: BaseStateHolder {
			game: g,
		},
	})
	g.pushStateHolder(&CivilStateHolder{
		base: BaseStateHolder {
			game: g,
		},
	})
}
type TurnEndStateHolder struct {
	base BaseStateHolder
}

func (h *TurnEndStateHolder) IsPending() bool {
	return false
}

func (h *TurnEndStateHolder) IsMoveLegal(move interface{}) (legal bool, reason string) {
	return true, ""
}

func (h *TurnEndStateHolder) Resolve(move interface{}) {
	fmt.Println("end resolve")
	g := h.base.game
	g.CurrentPlayer += 1
	if g.CurrentPlayer >= g.options.PlayerCount {
		g.CurrentPlayer = 0
		if g.getCurrentAge() == 0 {
			g.banishAgeACards()
		}
	}

	g.popStateHolder()
	g.pushStateHolder(&TurnStartStateHolder{
		base: BaseStateHolder {
			game: g,
		},
	})
}

type CivilStateHolder struct {
	base BaseStateHolder
	end  bool
}

func (h *CivilStateHolder) IsPending() bool {
	return !h.end
}

func (h *CivilStateHolder) IsMoveLegal(m interface{}) (legal bool, reason string) {
	move := m.(*Move)
	if move.FromPlayer != h.base.game.CurrentPlayer {
		return false, "Not current player."
	}
	p := h.base.game.players[h.base.game.CurrentPlayer]
	switch move.MoveType {
	case CIVIL_FETCH_CARD:
		if len(move.Data) != 1 {
			return false, "Invalid fetch command."
		}
		index := move.Data[0]
		if !p.canTakeCardFromWheel(index) {
			return false, "Invalid fetch command."
		}
		return true, ""
	case CIVIL_PLAY_CARD:
		if len(move.Data) < 1 {
			return false, "Invalid play command."
		}
		index := move.Data[0]
		var attachment interface{}
		if len(move.Data) > 1 {
			attachment = move.Data[1:]
		} else {
			attachment = nil
		}
		if !p.canPlayHand(index, attachment) {
			return false, "Invalid play command"
		}
		return true, ""
	case CIVIL_INC_POP:
		if !p.canIncreasePop() {
			return false, "Invalid incpop command"
		}
		p.increasePop()
	case CIVIL_BUILD:
		if len(move.Data) < 2 {
			return false, "Invalid build command."
		}
		stack := move.Data[0]
		index := move.Data[0]
		if !p.canBuild(stack, index, 0) {
			return false, "Invalid build command"
		}
		return true, ""
	case CIVIL_BUILD_WONDER:
		if len(move.Data) < 1 {
			return false, "Invalid buildwonder command."
		}
		step := move.Data[0]
		if !p.canBuildWonder(step, 0) {
			return false, "Invalid buildwonder command"
		}
		return true, ""
	case CIVIL_UPGRADE:
		if len(move.Data) < 3 {
			return false, "Invalid upgrade command."
		}
		stack := move.Data[0]
		index1 := move.Data[1]
		index2 := move.Data[2]
		if !p.canUpgrade(stack, index1, index2, 0) {
			return false, "Invalid upgrade command"
		}
		return true, ""
	case CIVIL_SPECIAL_ABILITY:
		if len(move.Data) < 1 {
			return false, "Invalid specialability command."
		}
		sa := move.Data[0]
		var attachment interface{}
		if len(move.Data) > 1 {
			attachment = move.Data[1:]
		} else {
			attachment = nil
		}
		if !p.canUseCivilSpecialAbility(sa, attachment) {
			return false, "Invalid specialability command"
		}
		return true, ""
	case CIVIL_END:
		return true, ""
	}
	return false, "Unknown command"
}

func (h *CivilStateHolder) Resolve(m interface{}) {
	if m == nil {
		return
	}
	move := m.(*Move)
	p := h.base.game.players[h.base.game.CurrentPlayer]
	switch move.MoveType {
	case CIVIL_FETCH_CARD:
		index := move.Data[0]
		p.takeCardFromWheel(index)
	case CIVIL_PLAY_CARD:
		index := move.Data[0]
		var attachment interface{}
		if len(move.Data) > 1 {
			attachment = move.Data[1:]
		} else {
			attachment = nil
		}
		p.playHand(index, attachment)
	case CIVIL_INC_POP:
		p.increasePop()
	case CIVIL_BUILD:
		stack := move.Data[0]
		index := move.Data[1]
		p.build(stack, index, 0)
	case CIVIL_BUILD_WONDER:
		step := move.Data[0]
		p.buildWonder(step, 0)
	case CIVIL_UPGRADE:
		stack := move.Data[0]
		index1 := move.Data[1]
		index2 := move.Data[2]
		p.upgrade(stack, index1, index2, 0)
	case CIVIL_SPECIAL_ABILITY:
		sa := move.Data[0]
		var attachment interface{}
		if len(move.Data) > 1 {
			attachment = move.Data[1:]
		} else {
			attachment = nil
		}
		p.useCivilSpecialAbility(sa, attachment)
	case CIVIL_END:
		h.base.game.popStateHolder()
	}
}

type ProductionPhaseStateHolder struct {
	base BaseStateHolder
}

func (h *ProductionPhaseStateHolder) IsPending() bool {
	return false
}

func (h *ProductionPhaseStateHolder) IsMoveLegal(m interface{}) (legal bool, reason string) {
	return true, ""
}

func (h *ProductionPhaseStateHolder) Resolve(m interface{}) {
	p := h.base.game.players[h.base.game.CurrentPlayer]
	p.doProductionPhase()
	h.base.game.popStateHolder()
}

type DiscardMilitaryCardsStateHolder struct {
	base      BaseStateHolder
	player    int
	toMaxHand bool
	toDiscard int
}

func (h *DiscardMilitaryCardsStateHolder) toDiscardMax() int {
	if h.toMaxHand {
		p := h.base.game.players[h.player]
		return p.getMilitaryHandSize() - p.getMaxMilitaryHandSize()
	} else {
		return h.toDiscard
	}
}

func (h *DiscardMilitaryCardsStateHolder) IsPending() bool {
	return h.toDiscardMax() > 0
}

func (h *DiscardMilitaryCardsStateHolder) IsMoveLegal(m interface{}) (legal bool, reason string) {
	if m == nil {
		return
	}
	move := m.(*Move)
	if len(move.Data) > h.toDiscardMax() {
		return false, "Too many cards."
	}

	p := h.base.game.players[h.player]
	return p.canDiscardMiliCards(move.Data), "Invalid card indexes or size."
}

func (h *DiscardMilitaryCardsStateHolder) Resolve(m interface{}) {
	if !h.IsPending() {
		h.base.game.popStateHolder()
		return
	}
	if m == nil {
		return
	}
	move := m.(*Move)
	p := h.base.game.players[h.player]
	p.discardMiliCards(move.Data)
	if !h.IsPending() {
		h.base.game.popStateHolder()
		return
	}
}
