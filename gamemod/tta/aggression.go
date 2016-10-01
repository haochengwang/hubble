package main

import (
	"fmt"
)

func resolveSuccessfulAggression(game *TtaGame, sourcePlayer, targetPlayer int,
	school *CardSchool) {
	fmt.Println("resolveSuccessfulAggression")
	if school.hasTrait(TRAIT_ENSLAVE) {
		popToLose := make([]int, game.options.PlayerCount)
		popToLose[targetPlayer] = 1
		cropOrResourceGain := make([]int, game.options.PlayerCount)
		cropOrResourceGain[sourcePlayer] = 2
		game.pushStateHolder(&GainCropResourceStateHolder{
			base: BaseStateHolder{
				game: game,
			},
			crop:     cropOrResourceGain,
			resource: cropOrResourceGain,
		})
		game.pushStateHolder(&LosePopulationStateHolder{
			base: BaseStateHolder{
				game: game,
			},
			popToLose: popToLose,
		})
	} else if school.hasTrait(TRAIT_PLUNDER) {
		game.pushStateHolder(&PlunderStateHolder{
			base: BaseStateHolder{
				game: game,
			},
			sourcePlayer: sourcePlayer,
			targetPlayer: targetPlayer,
			maxAmount:    school.actionBonus,
		})
	} else if school.hasTrait(TRAIT_RAID) {
		game.pushStateHolder(&RaidStateHolder{
			base: BaseStateHolder{
				game: game,
			},
			sourcePlayer:  sourcePlayer,
			targetPlayer:  targetPlayer,
			destroyCount:  school.actionBonus,
			destroyMaxAge: school.age,
		})
	} else if school.hasTrait(TRAIT_SPY) {
		totalTech := game.players[targetPlayer].getTechTotal()
		amount := school.actionBonus
		if amount > totalTech {
			amount = totalTech
		}
		game.players[targetPlayer].loseTech(amount)
		game.players[sourcePlayer].gainTech(amount)
	} else if school.hasTrait(TRAIT_ARMED_INTERVENTION) {
		game.players[targetPlayer].loseCulture(school.actionBonus)
		game.players[sourcePlayer].gainCulture(school.actionBonus)
	}
}

type PlunderStateHolder struct {
	base         BaseStateHolder
	sourcePlayer int
	targetPlayer int
	maxAmount    int
}

func (h *PlunderStateHolder) IsPending() bool {
	return true
}

func (h *PlunderStateHolder) IsMoveLegal(m interface{}) (legal bool, reason string) {
	move := m.(*Move)
	g := h.base.game
	if move.FromPlayer != h.sourcePlayer {
		return false, "Not valid user."
	}
	p := g.players[h.targetPlayer]
	switch move.MoveType {
	case MOVE_GENERAL_OP:
		if len(move.Data) != 2 {
			return false, "Invalid operation command"
		}
		crop := move.Data[0]
		resource := move.Data[1]
		if crop < 0 || crop > p.getCropTotal() {
			return false, "Invalid crop amount"
		}
		if resource < 0 || resource > p.getResourceTotal() {
			return false, "Invalid crop amount"
		}
		if crop+resource > h.maxAmount {
			return false, "Too many to plunder"
		}
		return true, ""
	}
	return false, ""
}

func (h *PlunderStateHolder) Resolve(m interface{}) {
	move := m.(*Move)
	g := h.base.game
	p := g.players[h.targetPlayer]
	switch move.MoveType {
	case MOVE_GENERAL_OP:
		crop := move.Data[0]
		resource := move.Data[1]
		p.spendCrop(crop)
		p.spendResource(resource)
		g.popStateHolder()
		// Gain crop and resource
		cropGain := make([]int, g.options.PlayerCount)
		cropGain[h.sourcePlayer] = crop
		resourceGain := make([]int, g.options.PlayerCount)
		resourceGain[h.sourcePlayer] = resource
		g.pushStateHolder(&GainCropResourceStateHolder{
			base: BaseStateHolder{
				game: g,
			},
			crop:     cropGain,
			resource: resourceGain,
		})
	}
}

type RaidStateHolder struct {
	base          BaseStateHolder
	sourcePlayer  int
	targetPlayer  int
	destroyCount  int
	destroyMaxAge int
	destroyed     []int
	destroyedCost int
}

func (h *RaidStateHolder) IsPending() bool {
	return true
}

func (h *RaidStateHolder) IsMoveLegal(m interface{}) (legal bool, reason string) {
	move := m.(*Move)
	g := h.base.game
	if move.FromPlayer != h.sourcePlayer {
		return false, "Not valid user."
	}
	p := g.players[h.targetPlayer]
	switch move.MoveType {
	case MOVE_DISBAND:
		if len(move.Data) != 2 {
			return false, "Invalid disband command"
		}
		stack := move.Data[0]
		index := move.Data[1]
		if !p.canDisband(stack, index) {
			return false, "Invalid disband command"
		}

		school := p.getCardSchool(stack, index)
		if !school.hasType(CARDTYPE_TECH_URBAN) {
			return false, "Invalid structure type"
		}
		if school.age > h.destroyMaxAge {
			return false, "structure too high level"
		}
		for _, d := range h.destroyed {
			if d >= h.destroyMaxAge && school.age == h.destroyMaxAge {
				return false, "structure too high level"
			}
		}
		return true, ""
	}
	return false, ""
}

func (h *RaidStateHolder) Resolve(m interface{}) {
	move := m.(*Move)
	g := h.base.game
	p := g.players[h.targetPlayer]
	switch move.MoveType {
	case MOVE_DISBAND:
		// Remove the building
		stack := move.Data[0]
		index := move.Data[1]
		p.disband(stack, index, true)

		// Record destroy
		school := p.getCardSchool(stack, index)
		h.destroyed = append(h.destroyed, school.age)
		h.destroyedCost += school.buildCost

		// Complete the resolve
		if len(h.destroyed) >= h.destroyCount {
			g.popStateHolder()
			cropGain := make([]int, g.options.PlayerCount)
			resourceGain := make([]int, g.options.PlayerCount)
			resourceGain[h.sourcePlayer] = (h.destroyedCost + 1) / 2
			g.pushStateHolder(&GainCropResourceStateHolder{
				base: BaseStateHolder{
					game: g,
				},
				crop:     cropGain,
				resource: resourceGain,
			})
		}
	case MOVE_END:
		// Complete the resolve
		if len(h.destroyed) >= h.destroyCount {
			g.popStateHolder()
			cropGain := make([]int, g.options.PlayerCount)
			resourceGain := make([]int, g.options.PlayerCount)
			resourceGain[h.sourcePlayer] = (h.destroyedCost + 1) / 2
			g.pushStateHolder(&GainCropResourceStateHolder{
				base: BaseStateHolder{
					game: g,
				},
				crop:     cropGain,
				resource: resourceGain,
			})
		}
	}
}
