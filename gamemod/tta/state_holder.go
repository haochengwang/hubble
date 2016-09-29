package main

type StateHolder interface {
	IsPending() bool
	IsMoveLegal(move interface{}) (legal bool, reason string)
	Resolve(move interface{})
}

type CivilStateHolder struct {
	end bool
}

func (h *CivilStateHolder) IsPending() bool {
	return !h.end
}

func (h *CivilStateHolder) IsMoveLegal(move interface{}) (legal bool, reason string) {
	return false, ""
}

func (h *CivilStateHolder) Resolve(move interface{}) {
}

type DiscardMilitaryCardStateHolder struct {
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
