package main

type StateHolder interface {
	IsPending() bool
	IsMoveLegal(move interface{}) (legal bool, reason string)
	Resolve(move interface{})
}

type BaseStateHolder struct {
	game *TtaGame
}

type CivilStateHolder struct {
	BaseStateHolder
	end bool
}

func (h *CivilStateHolder) IsPending() bool {
	return !h.end
}

func (h *CivilStateHolder) IsMoveLegal(m interface{}) (legal bool, reason string) {
	move := m.(*Move)
	if move.FromPlayer != h.game.CurrentPlayer {
		return false, "Not current player."
	}
	p := h.game.players[h.game.CurrentPlayer]
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
	p := h.game.players[h.game.CurrentPlayer]
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
		h.end = true
	}
}

type DiscardMilitaryCardStateHolder struct {
	BaseStateHolder
	player    int
	toDiscard int
}

func (h *DiscardMilitaryCardStateHolder) IsPending() bool {
	return h.toDiscard > 0
}

func (h *DiscardMilitaryCardStateHolder) IsMoveLegal(move interface{}) (legal bool, reason string) {
	return false, ""
}

func (h *DiscardMilitaryCardStateHolder) Resolve(move interface{}) {
}
