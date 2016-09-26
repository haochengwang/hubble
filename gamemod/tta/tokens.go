package main

type TokenBank map[int]int

type AddTokenRequest struct {
	bankId     int
	tokenType  int
	tokenCount int
}

type RemoveTokenRequest struct {
	bankId     int
	tokenType  int
	tokenCount int
}

type SetTokenRequest struct {
	bankId     int
	tokenType  int
	tokenCount int
}

type MoveTokenRequest struct {
	sourceBankId int
	targetBankId int
	tokenType    int
	tokenCount   int
}

type ClearTokenRequest struct {
	bankId    int
	tokenType int
}

type TokenBankUniversalManager struct {
	tokenBanks map[int]TokenBank
}

func NewTokenBankUniversalManager() *TokenBankUniversalManager {
	return &TokenBankUniversalManager{
		tokenBanks: make(map[int]TokenBank),
	}
}

func (m *TokenBankUniversalManager) getTokenCount(bankId, tokenType int) int {
	if bank, ok := m.tokenBanks[bankId]; ok {
		if tokenCount, ok := bank[tokenType]; ok {
			return tokenCount
		}
		return 0
	}
	return 0
}

func (m *TokenBankUniversalManager) setTokenCount(bankId, tokenType, tokenCount int) {
	if _, ok := m.tokenBanks[bankId]; !ok {
		m.tokenBanks[bankId] = TokenBank(make(map[int]int))
	}

	m.tokenBanks[bankId][tokenType] = tokenCount
}

func (m *TokenBankUniversalManager) modifyToken(bankId, tokenType, tokenDiff int) {
	if _, ok := m.tokenBanks[bankId]; !ok {
		m.tokenBanks[bankId] = TokenBank(make(map[int]int))
	}

	bank := m.tokenBanks[bankId]
	if _, ok := bank[tokenType]; !ok {
		if tokenDiff > 0 {
			bank[tokenType] = tokenDiff
		} else {
			bank[tokenType] = 0
		}
	} else {
		bank[tokenType] += tokenDiff
		if bank[tokenType] < 0 {
			bank[tokenType] = 0
		}
	}
}

func (m *TokenBankUniversalManager) processRequest(request interface{}) {
	switch request := request.(type) {
	case *AddTokenRequest:
		m.modifyToken(request.bankId, request.tokenType, request.tokenCount)
	case *RemoveTokenRequest:
		m.modifyToken(request.bankId, request.tokenType, -request.tokenCount)
	case *SetTokenRequest:
		m.setTokenCount(request.bankId, request.tokenType, request.tokenCount)
	case *MoveTokenRequest:
		m.modifyToken(request.sourceBankId, request.tokenType, -request.tokenCount)
		m.modifyToken(request.targetBankId, request.tokenType, request.tokenCount)
	}
}
