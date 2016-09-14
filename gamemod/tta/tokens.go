package main

type TokenBank map[int]int

type TokenBankUniversalManager struct {
	nextStackId int
	tokenBanks  map[int]*TokenBank
}
