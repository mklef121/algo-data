package main

import (
	"fmt"
)

type any interface{}

type Element struct {
	// Next and previous pointers in the doubly-linked list of elements.
	// To simplify the implementation, internally a list l is implemented
	// as a ring, such that &l.root is both the next element of the last
	// list element (l.Back()) and the previous element of the first list
	// element (l.Front()).
	Left, Right *Element

	// The list to which this element belongs.
	// list *List

	// The value stored with this element.
	Value any
}

func PrintLevelOrder(root Element) {
	var queue = make([]Element, 0)

	queue = append(queue, root)

	for len(queue) > 0 {
		first := queue[0]
		queue = queue[1:]

		//This is the operation that happens to the element
		fmt.Println(first.Value)

		if first.Left != nil {
			queue = append(queue, *first.Left)
		}

		if first.Right != nil {
			queue = append(queue, *first.Right)
		}

	}
}

func main() {
	// var  root = new Node(1);
	//     root.left = new Node(2);
	//     root.right = new Node(3);
	//     root.left.left = new Node(4);
	//     root.left.right = new Node(5);
	var root = Element{
		Value: 1,
		Left: &Element{
			Value: 2,
			Left: &Element{
				Value: 4,
			},
			Right: &Element{
				Value: 5,
			},
		},
		Right: &Element{
			Value: 3,
		},
	}

	// root.Left =

	// root.se
	fmt.Println(root)

	PrintLevelOrder(root)
}
