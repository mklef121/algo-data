package main

import (
	"errors"
	"fmt"
)

// "container/heap"

type Stack[T any] struct {
	items []T
	top   int
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
	s.top++
}

func (s *Stack[T]) Pop() {

	if s.IsEmpty() {
		return
	}

	// set the last top to the zero value of the stack data type to free up memory
	var zeroVal T
	s.items[s.top] = zeroVal

	s.top--
}

func (s *Stack[T]) Top() (T, error) {
	var zeroVal T // Declare a variable of type T, which will have its zero value
	if s.IsEmpty() {
		return zeroVal, errors.New("stack is empty")
	}

	return s.items[s.top], nil
}

func (s *Stack[T]) IsEmpty() bool {
	// if the top is negative, then the stack is empty
	if s.top <= -1 {
		return true
	}
	return false
}

func (s *Stack[T]) Print() {
	for _, item := range s.items {
		fmt.Print(item, " ")
	}
}

func New[T any]() Stack[T] {
	return Stack[T]{
		items: []T{},
		top:   -1,
	}
}
