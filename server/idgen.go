package main

import (
	"sync"
)

type IdGenerator struct {
	lock   sync.Mutex
	nextId int64
}

func NewIdGenerator(startId int64) *IdGenerator {
	return &IdGenerator{
		nextId: startId,
	}
}

func (g *IdGenerator) GetNextId() int64 {
	g.lock.Lock()
	defer g.lock.Unlock()

	ret := g.nextId
	g.nextId += 1
	return ret
}
