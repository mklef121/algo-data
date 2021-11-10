package main

import "fmt"

type Node struct {
	next  *Node
	value int
}

func (node *Node) setNext(next *Node) {
	node.next = next
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (list *LinkedList) prepend(node *Node) *LinkedList {
	node.next = list.head
	list.head = node
	list.length++

	return list
}

func (list *LinkedList) append(node *Node) *LinkedList {
	list.tail.next = node
	list.tail = node
	list.length++

	return list
}

func (list *LinkedList) insert(index int, value int) *LinkedList {
	node := &Node{value: value}
	if index == 0 {
		list.prepend(node)
		return list
	}

	if index >= list.length {
		list.append(node)
		return list
	}

	current := list.head.next
	previous := list.head
	var ite int = 1
	for current != nil {
		if ite == index {
			previous.next = node
			node.next = current
			list.length++
			break
		}

		current = current.next
		previous = current
		ite++
	}

	return list
}

func (list LinkedList) printList() {
	current := list.head
	// fmt.Printf("%d ", current.value)
	for list.length != 0 {
		fmt.Printf("%d ", current.value)
		current = current.next
		list.length--
	}

	fmt.Println("")
}
func CreateList(value int) *LinkedList {
	head := Node{
		value: value,
		next:  nil,
	}

	list := LinkedList{
		head:   &head,
		tail:   &head,
		length: 1,
	}

	return &list

}

func (list *LinkedList) reverse() *LinkedList {
	if list.head.next == nil {
		return list
	}

	first := list.head
	second := first.next
	//This will be the tail now
	first.next = nil

	list.tail = first

	count := 1

	for second != nil {
		next := second.next
		second.next = first

		count++

		if count == list.length {
			list.head = second
		}

		first = second
		second = next
	}

	return list

}

func main() {
	mylist := CreateList(34)

	mylist.prepend(&Node{value: 45}).append(&Node{value: 77}).prepend(&Node{value: 50}).prepend(&Node{value: 60}).insert(1, 69684)
	// .prepend(&Node{value: 50}).prepend(&Node{value: 60}).insert(1, 69684)

	mylist.printList()

	fmt.Println(mylist)

	mylist.reverse()

	mylist.printList()

}
