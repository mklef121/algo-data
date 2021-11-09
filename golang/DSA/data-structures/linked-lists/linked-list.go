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

func (list LinkedList) printList() {
	current := list.head
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

func main() {
	mylist := CreateList(34)

	mylist.prepend(&Node{value: 45}).prepend(&Node{value: 50}).prepend(&Node{value: 60})

	mylist.printList()

	fmt.Println(mylist)

}
