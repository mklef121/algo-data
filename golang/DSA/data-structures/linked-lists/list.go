package main

import "iter"

type LinkedListInterface[T any] interface {
	// Appends the specified element to the end of this list(the head).
	AddLast(value T) *Element[T]

	// Inserts the specified element at the beginning of this list(the tail).
	AddFirst(value T) *Element[T]

	// Inserts the specified element at the specified position in this list.
	Insert(index int, value T) *Element[T]

	// InsertBefore inserts a new element with value v immediately before mark and returns element inserted.
	InsertBefore(value T, mark *Element[T]) *Element[T]

	// InsertAfter inserts a new element with value v immediately after mark and returns the inserted element.
	InsertAfter(value T, mark *Element[T]) *Element[T]

	// Replaces the element at the specified position in this list with the specified element.
	Set(index int, value T) *Element[T]

	// Returns true if this list contains the specified element.
	Contains(Object *Element[T]) bool

	//  Returns the element at the specified position in this list.
	Get(index int) *Element[T]

	// Returns the first element in this list.
	GetFirst() *Element[T]

	// Returns the last element in this list.
	GetLast() *Element[T]

	// Returns the number of items on the linked list.
	GetLength() int

	// This provides the ability to loop through the elements of the list using a range.
	Iterator() iter.Seq2[int, *Element[T]]
}

// Element is an element of a linked list.
type Element[T any] struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	next, prev *Element[T]

	// The list to which this element belongs.
	list *List[T]

	// The value stored with this element.
	Value any
}

type List[T any] struct {
	head   *Element[T]
	tail   *Element[T]
	length int
}

func (l *List[T]) createElement(val T, next, prev *Element[T]) *Element[T] {
	return &Element[T]{
		Value: val,
		next:  next,
		prev:  prev,
		list:  l,
	}
}

func (l *List[T]) GetFirst() *Element[T] {
	return l.tail
}

func (l *List[T]) GetLast() *Element[T] {
	return l.head
}

func (l *List[T]) Get(index int) *Element[T] {
	current := l.head

	for i := 0; i < l.length; i++ {
		if i == index {
			return current
		}

		current = current.next
	}

	return nil
}

func (l *List[T]) Contains(el *Element[T]) bool {
	if el != nil && el.list == l {
		return true
	}

	return false
}

// Inserts the specified element at the beginning of this list, making it the new tail
func (l *List[T]) AddFirst(element T) *Element[T] {
	node := l.createElement(element, nil, l.tail)

	l.tail.next = node
	l.tail = node
	l.length++

	return node
}

// Appends the specified element to the end of this list, making it the new head
func (l *List[T]) AddLast(element T) *Element[T] {
	node := l.createElement(element, l.head, nil)

	l.head.prev = node
	l.head = node
	l.length++

	return node
}

func (l *List[T]) GetLength() int {
	return l.length
}

// Set replaces the element at the specified position in this list with the specified element.
func (l *List[T]) Set(index int, value T) *Element[T] {
	current := l.head

	for i := 0; i < l.length; i++ {
		if i == index {
			node := l.createElement(value, nil, nil)

			node.prev = current.prev
			node.next = current.next

			if current.prev != nil {
				current.prev.next = node
			}

			if current.next != nil {
				current.next.prev = node
			}
			return node
		}

		current = current.next
	}
	return nil
}

func (l *List[T]) InsertBefore(value T, mark *Element[T]) (e *Element[T]) {

	current := l.head
	// TODO: use binary search to find this
	for i := 0; i < l.length; i++ {
		if current == mark {
			// the default behaviour of the Insert method is to perform insert Before
			return l.insertAt(i, value, true)
		}

		current = current.next
	}

	return nil
}

func (l *List[T]) InsertAfter(value T, mark *Element[T]) (e *Element[T]) {

	current := l.head
	// TODO: use binary search to find this
	for i := 0; i < l.length; i++ {
		if current == mark {
			// the default behaviour of the Insert method is to perform insert Before
			return l.insertAt(i, value, false)
		}

		current = current.next
	}

	return nil
}

// Inserts the specified element at the specified position in this list.
func (l *List[T]) Insert(index int, value T) *Element[T] {
	return l.insertAt(index, value, true)
}

// insertAt will insert an element at a particular position in the list, it uses the
// before flag to know if it should put the element  behind or in front of the
// element occupying the current position.
func (l *List[T]) insertAt(index int, value T, before bool) (e *Element[T]) {

	if index < 0 || l.length == 0 || index >= l.length {
		return
	}

	if index == 0 && before {
		return l.AddLast(value)
	}

	// add the element after the tail(top of the list) if before is false
	// use this to ensure list is not iterated fully
	if index == l.length-1 && !before {
		return l.AddFirst(value)
	}

	node := l.createElement(value, nil, nil)

	// so we start counting from the Node after the head to find the provided index
	// TODO: use binary search to make this faster
	current := l.head
	currentIndex := 0
	for current != nil {

		if currentIndex == index {
			// element to be added before the found item
			if before {
				if current.prev != nil {
					node.prev = current.prev
					current.prev.next = node
				}
				node.next = current
				current.prev = node
			} else {

				if current.next != nil {
					node.next = current.next
					current.next.prev = node
				}
				node.prev = current
				current.next = node
			}

			l.length++
			return node
		}

		current = current.next
		currentIndex++
	}
	return nil
}

func (l *List[T]) Iterator() iter.Seq2[int, *Element[T]] {
	return func(yield func(int, *Element[T]) bool) {
		current := l.head
		for i := 0; i < l.length; i++ {
			// once the loop breaks or the loop function returns, this yield will become false
			if !yield(i, current) {
				return
			}

			current = current.next
		}
	}
}

func NewList[T any](val T) LinkedListInterface[T] {

	node := Element[T]{
		next:  nil,
		prev:  nil,
		Value: val,
	}

	return &List[T]{
		head:   &node,
		tail:   &node,
		length: 1,
	}
}
