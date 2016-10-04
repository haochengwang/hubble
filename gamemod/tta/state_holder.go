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
	fmt.Println("TurnStartStateHolder.Resolve")
	g := h.base.game
	// Only rotate the wheel when age is not A
	if g.RoundCount > 0 {
		g.weedOut(3)
		g.refillWheels()
	}

	if g.RoundCount > 0 && g.CurrentPlayer == 0 {
		g.banishAgeACards()
	}

	// Put tactic to shared zone
	g.shareTactic(g.CurrentPlayer)
	g.popStateHolder()
	g.pushStateHolder(&TurnEndStateHolder{
		base: BaseStateHolder{
			game: g,
		},
	})
	g.pushStateHolder(&DrawMilitaryCardsStateHolder{
		base: BaseStateHolder{
			game: g,
		},
		player:      g.CurrentPlayer,
		toRedTokens: true,
	})
	g.pushStateHolder(&DiscardMilitaryCardsStateHolder{
		base: BaseStateHolder{
			game: g,
		},
		player:    g.CurrentPlayer,
		toMaxHand: true,
	})
	g.pushStateHolder(&ProductionPhaseStateHolder{
		base: BaseStateHolder{
			game: g,
		},
	})
	g.pushStateHolder(&CivilStateHolder{
		base: BaseStateHolder{
			game: g,
		},
	})
	g.pushStateHolder(&PoliticalStateHolder{
		base: BaseStateHolder{
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
	fmt.Println("TurnEndStateHolder.Resolve")
	g := h.base.game
	p := g.players[g.CurrentPlayer]
	p.refillWhiteRedTokens()
	p.clearupTurn()
	g.CurrentPlayer += 1
	if g.CurrentPlayer >= g.options.PlayerCount {
		g.CurrentPlayer = 0
		g.RoundCount++
		fmt.Println(g.RoundCount)
	}

	g.popStateHolder()
	g.pushStateHolder(&TurnStartStateHolder{
		base: BaseStateHolder{
			game: g,
		},
	})
}

type CivilStateHolder struct {
	base BaseStateHolder
}

func (h *CivilStateHolder) IsPending() bool {
	return true
}

func (h *CivilStateHolder) IsMoveLegal(m interface{}) (legal bool, reason string) {
	move := m.(*Move)
	if move.FromPlayer != h.base.game.CurrentPlayer {
		return false, "Not current player."
	}
	p := h.base.game.players[h.base.game.CurrentPlayer]
	if h.base.game.getCurrentAge() == 0 { // Age A only fetch allowed
		if move.MoveType != MOVE_FETCH_CARD && move.MoveType != MOVE_END {
			return false, "Only fetch card is allowed in first round"
		}
	}
	switch move.MoveType {
	case MOVE_FETCH_CARD:
		if len(move.Data) != 1 {
			return false, "Invalid fetch command."
		}
		index := move.Data[0]
		if !p.canTakeCardFromWheel(index) {
			return false, "Invalid fetch command."
		}
		return true, ""
	case MOVE_PLAY_CIVIL_CARD:
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
	case MOVE_PLAY_MILITARY_CARD:
		if len(move.Data) != 1 {
			return false, "Invalid playmilitary command."
		}
		index := move.Data[0]
		if !p.civilPlayTacticLegal(index) {
			return false, "Invalid play command"
		}
		return true, ""
	case MOVE_INC_POP:
		if !p.canIncreasePop() {
			return false, "Invalid incpop command"
		}
		return true, ""
	case MOVE_BUILD:
		if len(move.Data) < 2 {
			return false, "Invalid build command."
		}
		stack := move.Data[0]
		index := move.Data[1]
		if !p.canBuild(stack, index, 0) {
			return false, "Invalid build command"
		}
		return true, ""
	case MOVE_BUILD_WONDER:
		if len(move.Data) < 1 {
			return false, "Invalid buildwonder command."
		}
		step := move.Data[0]
		if !p.canBuildWonder(step, 0) {
			return false, "Invalid buildwonder command"
		}
		return true, ""
	case MOVE_UPGRADE:
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
	case MOVE_DISBAND:
		if len(move.Data) < 2 {
			return false, "Invalid disband command."
		}
		stack := move.Data[0]
		index := move.Data[1]
		if !p.civilDisbandLegal(stack, index) {
			return false, "Invalid disband command"
		}
		return true, ""
	case MOVE_SPECIAL_ABILITY:
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
	case MOVE_LEARN_TACTIC:
		if len(move.Data) != 1 {
			return false, "Invalid learntactic command."
		}
		index := move.Data[0]
		if !p.civilLearnTacticLegal(index) {
			return false, "Invalid learntactic command"
		}
		return true, ""
	case MOVE_END:
		return true, ""
	}
	return false, "Unknown command"
}

func (h *CivilStateHolder) Resolve(m interface{}) {
	fmt.Println("CivilStateHolder.Resolve")
	if m == nil {
		return
	}
	move := m.(*Move)
	p := h.base.game.players[h.base.game.CurrentPlayer]

	switch move.MoveType {
	case MOVE_FETCH_CARD:
		index := move.Data[0]
		p.takeCardFromWheel(index)
	case MOVE_PLAY_CIVIL_CARD:
		index := move.Data[0]
		var attachment interface{}
		if len(move.Data) > 1 {
			attachment = move.Data[1:]
		} else {
			attachment = nil
		}
		p.playHand(index, attachment)
	case MOVE_PLAY_MILITARY_CARD:
		index := move.Data[0]
		p.civilPlayTactic(index)
	case MOVE_INC_POP:
		p.increasePop()
	case MOVE_BUILD:
		stack := move.Data[0]
		index := move.Data[1]
		p.build(stack, index, 0)
	case MOVE_BUILD_WONDER:
		step := move.Data[0]
		p.buildWonder(step, 0)
	case MOVE_UPGRADE:
		stack := move.Data[0]
		index1 := move.Data[1]
		index2 := move.Data[2]
		p.upgrade(stack, index1, index2, 0)
	case MOVE_DISBAND:
		stack := move.Data[0]
		index := move.Data[1]
		p.civilDisband(stack, index)
	case MOVE_SPECIAL_ABILITY:
		sa := move.Data[0]
		var attachment interface{}
		if len(move.Data) > 1 {
			attachment = move.Data[1:]
		} else {
			attachment = nil
		}
		p.useCivilSpecialAbility(sa, attachment)
	case MOVE_LEARN_TACTIC:
		index := move.Data[0]
		p.civilLearnTactic(index)
	case MOVE_END:
		h.base.game.popStateHolder()
	}
}

type PoliticalStateHolder struct {
	base BaseStateHolder
}

func (h *PoliticalStateHolder) IsPending() bool {
	return h.base.game.getCurrentAge() != 0
}

func (h *PoliticalStateHolder) IsMoveLegal(m interface{}) (legal bool, reason string) {
	move := m.(*Move)
	if move.FromPlayer != h.base.game.CurrentPlayer {
		return false, "Not current player."
	}
	p := h.base.game.players[h.base.game.CurrentPlayer]
	switch move.MoveType {
	case MOVE_PLAY_MILITARY_CARD:
		if len(move.Data) < 1 {
			return false, "Invalid playmilitary command."
		}
		index := move.Data[0]
		var attachment interface{}
		if len(move.Data) > 1 {
			attachment = move.Data[1:]
		} else {
			attachment = nil
		}
		if !p.politicalPlayMilitaryHandLegal(index, attachment) {
			return false, "Invalid playmilitary command"
		}
		return true, ""
	case MOVE_BREAK_PACT:
		if len(move.Data) != 1 {
			return false, "Invalid breakpact command."
		}
		pid := move.Data[0]
		if !p.politicalBreakPactLegal(pid) {
			return false, "Invalid breakpact command"
		}
		return true, ""
	case MOVE_END:
		return true, ""
	}
	return false, "Unknown command"
}

func (h *PoliticalStateHolder) Resolve(m interface{}) {
	fmt.Println("PoliticalStateHolder.Resolve")
	g := h.base.game
	g.popStateHolder()
	if m == nil {
		return
	}
	move := m.(*Move)
	p := g.players[h.base.game.CurrentPlayer]
	switch move.MoveType {
	case MOVE_PLAY_MILITARY_CARD:
		index := move.Data[0]
		var attachment interface{}
		if len(move.Data) > 1 {
			attachment = move.Data[1:]
		} else {
			attachment = nil
		}
		p.politicalPlayMilitaryHand(index, attachment)

		// Check if aggression or pact is played
		pp, pendingCard := g.getPendingAggressionOrPact()
		if pendingCard != nil {
			school := g.cardSchools[pendingCard.schoolId]
			if school.hasType(CARDTYPE_AGGRESSION) {
				g.pushStateHolder(&DefenseAggressionStateHolder{
					base: BaseStateHolder{
						game: g,
					},
					sourcePlayer: g.CurrentPlayer,
					sourcePower:  p.calcPower(),
					player:       pp,
				})
			} else if school.hasType(CARDTYPE_PACT) {
				aSideSelected := g.cardTokenManager.getTokenCount(
					pendingCard.id, PACT_A) > 0
				g.pushStateHolder(&ConfirmPactStateHolder{
					base: BaseStateHolder{
						game: g,
					},
					sourcePlayer:  g.CurrentPlayer,
					player:        pp,
					aSideSelected: aSideSelected,
				})
			} else {
			}
		} else {
			g.pushStateHolder(NewColonizeStateHolder(g))

		}
	case MOVE_BREAK_PACT:
		pid := move.Data[0]
		p.politicalBreakPact(pid)
	case MOVE_END:
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
	fmt.Println("ProductionPhaseStateHolder.Resolve")
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
	fmt.Println("DiscardMilitaryCardsStateHolder.Resolve")
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
	if move.FromPlayer != h.base.game.CurrentPlayer {
		return false, "Not current player."
	}
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

type DrawMilitaryCardsStateHolder struct {
	base        BaseStateHolder
	player      int
	toRedTokens bool
	toDraw      int
}

func (h *DrawMilitaryCardsStateHolder) drawCount() int {
	if h.toRedTokens {
		p := h.base.game.players[h.player]
		result := p.getUsableRedTokens()
		if result > 3 {
			return 3
		} else {
			return result
		}
	} else {
		return h.toDraw
	}
}

func (h *DrawMilitaryCardsStateHolder) IsPending() bool {
	return false
}

func (h *DrawMilitaryCardsStateHolder) IsMoveLegal(m interface{}) (legal bool, reason string) {
	return false, ""
}

func (h *DrawMilitaryCardsStateHolder) Resolve(m interface{}) {
	fmt.Println("DrawMilitaryCardsStateHolder.Resolve")
	p := h.base.game.players[h.player]
	p.drawMiliCards(h.drawCount())
	h.base.game.popStateHolder()
}

type DefenseAggressionStateHolder struct {
	base         BaseStateHolder
	sourcePlayer int
	sourcePower  int
	player       int
}

func (h *DefenseAggressionStateHolder) IsPending() bool {
	return true
}

func (h *DefenseAggressionStateHolder) IsMoveLegal(m interface{}) (legal bool, reason string) {
	move := m.(*Move)
	p := h.base.game.players[h.player]
	switch move.MoveType {
	case MOVE_PLAY_MILITARY_CARD:
		if len(move.Data) < 1 {
			return false, "Invalid playmilitary command."
		}
		if !p.defenseAggressionLegal(move.Data) {
			return false, "Invalid playmilitary command"
		}
		return true, ""
	case MOVE_END:
		return true, ""
	}
	return false, "Invalid command"
}

func (h *DefenseAggressionStateHolder) Resolve(m interface{}) {
	move := m.(*Move)
	g := h.base.game
	p := g.players[h.player]
	powerBonus := p.defenseAggressionPowerBonus(move.Data)
	p.defenseAggression(move.Data)
	g.popStateHolder()

	// Check if the aggression is successful
	if h.sourcePower > p.calcPower()+powerBonus {
		fmt.Println("Aggression is sucgcessful")
		fmt.Println(h.sourcePower, "vs", p.calcPower(), "+", powerBonus)
		_, pendingCard := g.popPendingAggressionOrPact()
		school := g.cardSchools[pendingCard.schoolId]
		resolveSuccessfulAggression(h.base.game, h.sourcePlayer, h.player, school)
	} else {
		fmt.Println("Aggression has failed ")
		fmt.Println(h.sourcePower, "vs", p.calcPower(), "+", powerBonus)
		g.popPendingAggressionOrPact()
	}
}

type ConfirmPactStateHolder struct {
	base          BaseStateHolder
	sourcePlayer  int
	player        int
	aSideSelected bool
}

func (h *ConfirmPactStateHolder) IsPending() bool {
	return true
}

func (h *ConfirmPactStateHolder) IsMoveLegal(m interface{}) (legal bool, reason string) {
	move := m.(*Move)
	switch move.MoveType {
	case MOVE_GENERAL_OP:
		if len(move.Data) != 1 {
			return false, "Invalid operation command."
		}
		return true, ""
	case MOVE_END:
		return true, ""
	}
	return false, "Invalid command"
}

func (h *ConfirmPactStateHolder) Resolve(m interface{}) {
	move := m.(*Move)
	g := h.base.game
	g.popStateHolder()

	accepted := false
	switch move.MoveType {
	case MOVE_GENERAL_OP:
		if move.Data[0] > 0 {
			accepted = true
		}
	}
	// Check if the pact is accepted
	if accepted {
		fmt.Println("Pact is accpeted")
		g.acceptPendingPact()
	} else {
		fmt.Println("Pact is rejected")
		g.rejectPendingPact()
	}
}

type LosePopulationStateHolder struct {
	base      BaseStateHolder
	popToLose []int
}

func (h *LosePopulationStateHolder) IsPending() bool {
	g := h.base.game
	for i, p := range g.players {
		if p.getFreeWorkers() < h.popToLose[i] {
			return true
		}
	}
	return false
}

func (h *LosePopulationStateHolder) IsMoveLegal(m interface{}) (legal bool, reason string) {
	move := m.(*Move)
	g := h.base.game
	if move.FromPlayer < 0 || move.FromPlayer >= g.options.PlayerCount {
		return false, "Not valid user."
	}
	if h.popToLose[move.FromPlayer] <= g.players[move.FromPlayer].getFreeWorkers() {
		return false, "Not valid user, no need to lose pop"
	}
	p := h.base.game.players[move.FromPlayer]
	switch move.MoveType {
	case MOVE_DISBAND:
		if len(move.Data) < 2 {
			return false, "Invalid disband command."
		}
		stack := move.Data[0]
		index := move.Data[1]
		return p.canDisband(stack, index), ""
	}
	return false, ""
}

func (h *LosePopulationStateHolder) Resolve(m interface{}) {
	if !h.IsPending() {
		h.base.game.popStateHolder()
		for i, p := range h.base.game.players {
			p.removeFreeWorkers(h.popToLose[i])
		}
		return
	}
	if m == nil {
		return
	}
	move := m.(*Move)
	p := h.base.game.players[move.FromPlayer]
	switch move.MoveType {
	case MOVE_DISBAND:
		stack := move.Data[0]
		index := move.Data[1]
		p.canDisband(stack, index)
	}

	if !h.IsPending() {
		h.base.game.popStateHolder()
	}
}

type GainCropResourceStateHolder struct {
	base     BaseStateHolder
	crop     []int
	resource []int
}

func (h *GainCropResourceStateHolder) IsPending() bool {
	return false
}

func (h *GainCropResourceStateHolder) IsMoveLegal(m interface{}) (legal bool, reason string) {
	return false, ""
}

func (h *GainCropResourceStateHolder) Resolve(m interface{}) {
	fmt.Println("GainCropResourceStateHolder.Resolve", h.crop, h.resource)
	for i, r := range h.resource {
		p := h.base.game.players[i]
		p.gainResource(r)
	}
	for i, c := range h.crop {
		p := h.base.game.players[i]
		p.gainCrop(c)
	}
	h.base.game.popStateHolder()
}
