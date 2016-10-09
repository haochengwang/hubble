package main

import (
	"fmt"
)

type GeneralEventStateHolder struct {
	base BaseStateHolder
}

func NewEventStateHolder(game *TtaGame, school *CardSchool) StateHolder {
	if school.hasType(CARDTYPE_EVENT) {
		return &GeneralEventStateHolder{
			base: BaseStateHolder{
				game: game,
			},
		}

	} else if school.hasType(CARDTYPE_TERRITORY) {
		return &ColonizeStateHolder{
			base: BaseStateHolder{
				game: game,
			},
		}
	} else {
		fmt.Println("NewEventStateHolder not valid event")
		return nil
	}
}

func (h *GeneralEventStateHolder) IsPending() bool {
	return false
}

func (h *GeneralEventStateHolder) IsMoveLegal(move interface{}) (legal bool, reason string) {
	return false, ""
}

func (h *GeneralEventStateHolder) Resolve(move interface{}) {
	fmt.Println("GeneralEventStateHolder.Resolve")
	g := h.base.game
	csm := g.cardStackManager

	card := csm.getFirstCard(g.pastEventsDeck)
	school := g.cardSchools[card.schoolId]

	g.popStateHolder()
	if school.hasTrait(TRAIT_DEVELOPMENT_OF_AGRICULTURE) {
		// TODO
	}
}
