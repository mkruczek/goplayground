package linkedlist_vs_slice

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (list *LinkedList) Add(data int) {
	newNode := &Node{Data: data}

	if list.Head == nil {
		list.Head = newNode
	} else {
		list.Tail.Next = newNode
	}

	list.Tail = newNode
}

func (list *LinkedList) Display() {
	for node := list.Head; node != nil; node = node.Next {
		fmt.Println(node.Data)
	}
}

func (list *LinkedList) Size() int {
	count := 0
	for node := list.Head; node != nil; node = node.Next {
		count++
	}
	return count
}

func (list *LinkedList) IsEmpty() bool {
	return list.Head == nil
}

func (list *LinkedList) Remove(data int) {
	if list.IsEmpty() {
		return
	}

	if list.Head.Data == data {
		list.Head = list.Head.Next
		if list.Head == nil {
			list.Tail = nil
		}
		return
	}

	prev := list.Head
	for prev.Next != nil && prev.Next.Data != data {
		prev = prev.Next
	}

	if prev.Next != nil {
		prev.Next = prev.Next.Next
		if prev.Next == nil {
			list.Tail = prev
		}
	}
}

func (list *LinkedList) Search(data int) (*Node, bool) {
	for node := list.Head; node != nil; node = node.Next {
		if node.Data == data {
			return node, true
		}
	}
	return nil, false
}
