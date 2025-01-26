package main

import (
	"testing"
)

func TestAddFirst(t *testing.T) {
	list := NewList(1).(*List[int])
	list.AddFirst(2)
	list.AddFirst(3)

	if list.GetFirst().Value != 3 {
		t.Errorf("Expected first element to be 3, got %v", list.GetFirst().Value)
	}

	if list.GetLast().Value != 1 {
		t.Errorf("Expected last element to be 1, got %v", list.GetLast().Value)
	}

	if list.GetLength() != 3 {
		t.Errorf("Expected length to be 3, got %d", list.GetLength())
	}
}

func TestAddLast(t *testing.T) {
	list := NewList(1).(*List[int])
	list.AddLast(2)
	list.AddLast(3)

	if list.GetFirst().Value != 1 {
		t.Errorf("Expected first element to be 1, got %v", list.GetFirst().Value)
	}

	if list.GetLast().Value != 3 {
		t.Errorf("Expected last element to be 3, got %v", list.GetLast().Value)
	}

	if list.GetLength() != 3 {
		t.Errorf("Expected length to be 3, got %d", list.GetLength())
	}
}

func TestInsert(t *testing.T) {
	list := NewList(1).(*List[int])
	list.AddLast(3)
	list.Insert(1, 2)

	if list.Get(1).Value != 2 {
		t.Errorf("Expected element at index 1 to be 2, got %v", list.Get(1).Value)
	}

	// fmt.Printf("%#v \n\n", list)
	// fmt.Println(list.Get(0).Value, list.Get(1).Value, list.Get(2).Value)
	if list.Get(2).Value != 1 {
		t.Errorf("Expected element at index 2 to be 1, got %v", list.Get(2).Value)
	}

	if list.GetLength() != 3 {
		t.Errorf("Expected length to be 3, got %d", list.GetLength())
	}
}

func TestInsertBefore(t *testing.T) {
	list := NewList(1).(*List[int])
	element := list.AddLast(2)
	list.InsertBefore(0, element)

	if list.GetLast().Value != 0 {
		t.Errorf("Expected first element to be 0, got %v", list.GetFirst().Value)
	}
}

func TestInsertAfter(t *testing.T) {
	list := NewList(1).(*List[int])
	element := list.AddLast(2)
	list.InsertAfter(3, element)

	if list.Get(1).Value != 3 {
		t.Errorf("Expected last element to be 3, got %v", list.GetLast().Value)
	}
}

func TestSet(t *testing.T) {
	list := NewList(1).(*List[int])
	list.AddLast(2)
	oldLength := list.length
	list.Set(1, 42)

	if list.Get(1).Value != 42 {
		t.Errorf("Expected element at index 1 to be 42, got %v", list.Get(1).Value)
	}

	if oldLength != list.length {
		t.Errorf("Expected length of the list not to change. Initial length %v, current length: %v", oldLength, list.length)
	}
}

func TestContains(t *testing.T) {
	list := NewList(1).(*List[int])
	element := list.AddLast(2)

	if !list.Contains(element) {
		t.Errorf("Expected list to contain the element")
	}

	nonExistentElement := &Element[int]{Value: 3}
	if list.Contains(nonExistentElement) {
		t.Errorf("Expected list not to contain the element")
	}
}

func TestGet(t *testing.T) {
	list := NewList(1).(*List[int])
	list.AddLast(2)
	list.AddLast(3)

	if list.Get(1).Value != 2 {
		t.Errorf("Expected element at index 1 to be 2, got %v", list.Get(1).Value)
	}

	if list.Get(2).Value != 1 {
		t.Errorf("Expected element at index 2 to be 1, got %v", list.Get(2).Value)
	}

	if list.Get(0).Value != 3 {
		t.Errorf("Expected element at index 0 to be 3, got %v", list.Get(0).Value)
	}
}

func TestGetFirstAndLast(t *testing.T) {
	list := NewList(1).(*List[int])
	list.AddLast(2)
	list.AddLast(3)

	if list.GetFirst().Value != 1 {
		t.Errorf("Expected first element to be 1, got %v", list.GetFirst().Value)
	}

	if list.GetLast().Value != 3 {
		t.Errorf("Expected last element to be 3, got %v", list.GetLast().Value)
	}
}

func TestGetLength(t *testing.T) {
	list := NewList(1).(*List[int])
	list.AddLast(2)
	list.AddLast(3)

	if list.GetLength() != 3 {
		t.Errorf("Expected length to be 3, got %d", list.GetLength())
	}
}

func TestIterator(t *testing.T) {
	// Create a new list with some elements
	list := NewList(1).(*List[int])
	list.AddLast(2)
	list.AddLast(3)
	list.AddLast(4)

	expectedValues := []int{4, 3, 2, 1}
	expectedIndexes := []int{0, 1, 2, 3}
	index := 0

	// Use a custom range implementation for the iterator
	for i, el := range list.Iterator() {
		// Check the index
		if i != expectedIndexes[index] {
			t.Errorf("Expected index %d, got %d", expectedIndexes[index], i)
		}

		// Check the value of the element
		if el.Value != expectedValues[index] {
			t.Errorf("Expected value %d, got %v", expectedValues[index], el.Value)
		}

		index++
	}

	// Check if all elements were iterated
	if index != len(expectedValues) {
		t.Errorf("Iterator did not iterate over all elements; expected %d, got %d", len(expectedValues), index)
	}
}

func TestIteratorEarlyExit(t *testing.T) {
	// Create a new list with some elements
	list := NewList(1).(*List[int])
	list.AddLast(2)
	list.AddLast(3)
	list.AddLast(4) // this becomes the latest head

	expectedValues := []int{4, 3}
	expectedIndexes := []int{0, 1}
	index := 0

	// Use the Iterator to traverse the list and exit early
	list.Iterator()(func(i int, el *Element[int]) bool {
		// Check the index
		if i != expectedIndexes[index] {
			t.Errorf("Expected index %d, got %d", expectedIndexes[index], i)
		}

		// Check the value of the element
		if el.Value != expectedValues[index] {
			t.Errorf("Expected value %d, got %v", expectedValues[index], el.Value)
		}

		index++
		if index == len(expectedValues) {
			return false // Stop iteration early
		}

		return true
	})

	// Ensure iteration stopped early
	if index != len(expectedValues) {
		t.Errorf("Iterator did not exit early as expected; stopped at index %d", index)
	}
}
