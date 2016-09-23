package main

import (
	"fmt"
)

type ArrayList struct {
	data []interface{}
}

func NewEmptyArrayList() *ArrayList {
	return &ArrayList{
		data: make([]interface{}, 0),
	}
}

func NewArrayList(x []interface{}) *ArrayList {
	data := make([]interface{}, len(x))
	for i, _ := range x {
		data[i] = x[i]
	}
	return &ArrayList{
		data: data,
	}
}

func (l *ArrayList) Push(x interface{}) {
	l.data = append(l.data, x)
}

func (l *ArrayList) PushAt(index int, x interface{}) (err error) {
	if index >= 0 && index < len(l.data) {
		rear := append(make([]interface{}, 0), l.data[index:]...)
		l.data = append(l.data[:index], x)
		l.data = append(l.data, rear...)
	} else if index == len(l.data) {
		l.data = append(l.data, x)
	} else {
		return fmt.Errorf("Invalid PushAt, index out of range: ", index)
	}
	return nil
}

func (l *ArrayList) Pop() (x interface{}, err error) {
	size := len(l.data)
	if size <= 0 {
		return nil, fmt.Errorf("Invalid Pop, empty ArrayList.")
	}
	x = l.data[size-1]
	l.data = l.data[:size-1]
	return x, nil
}

func (l *ArrayList) PopAt(index int) (x interface{}, err error) {
	size := len(l.data)
	if index < 0 || index >= size {
		return nil, fmt.Errorf("Invalid Pop, index out of range.")
	}
	x = l.data[size-1]
	l.data = append(l.data[:index], l.data[index+1:]...)
	return x, nil
}

func (l *ArrayList) Set(index int, x interface{}) (err error) {
	size := len(l.data)
	if index < 0 || index >= size {
		return fmt.Errorf("Invalid At, index out of range.")
	}
	l.data[index] = x
	return nil
}

func (l *ArrayList) At(index int) (x interface{}, err error) {
	size := len(l.data)
	if index < 0 || index >= size {
		return nil, fmt.Errorf("Invalid At, index out of range.")
	}
	return l.data[index], nil
}

func (l *ArrayList) Length() int {
	return len(l.data)
}
