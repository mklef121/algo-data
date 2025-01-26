package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("Test Push and Top", func(t *testing.T) {
		stack := New[int]()
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)

		if top, _ := stack.Top(); top != 3 {
			t.Errorf("Expected top to be 3, but got %d", top)
		}
	})

	t.Run("Test Pop", func(t *testing.T) {
		stack := New[int]()
		stack.Push(1)
		stack.Push(2)
		stack.Pop()

		if top, _ := stack.Top(); top != 1 {
			t.Errorf("Expected top to be 1, but got %d", top)
		}
	})

	t.Run("Test IsEmpty", func(t *testing.T) {
		stack := New[int]()

		if !stack.IsEmpty() {
			t.Error("Expected stack to be empty")
		}

		stack.Push(1)

		if stack.IsEmpty() {
			t.Error("Expected stack to be non-empty")
		}
	})

	t.Run("Test Pop on empty stack", func(t *testing.T) {
		stack := New[int]()
		stack.Pop()

		if !stack.IsEmpty() {
			t.Error("Expected stack to be empty after Pop")
		}
	})

	t.Run("Test Top on empty stack", func(t *testing.T) {
		stack := New[int]()
		_, err := stack.Top()

		if err == nil {
			t.Error("Expected an error when calling Top on empty stack")
		}
	})
}
