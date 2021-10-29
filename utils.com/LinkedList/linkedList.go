package linkedlist

import "fmt"

type listNode struct {
	value    interface{}
	next     *listNode
	previous *listNode
}

type linkedList struct {
	head *listNode
}

func NewList() *linkedList {
	return new(linkedList)
}

func (l *linkedList) AddNode(value interface{}) *listNode {
	if l.head == nil {
		l.head = &listNode{value: value}
		return l.head
	}

	currentNode := l.head

	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	newNode := new(listNode)
	newNode.value = value
	newNode.previous = currentNode
	currentNode.next = newNode

	return currentNode.next
}

func (l *linkedList) Print() {
	current := l.head

	for current != nil {
		if current.next != nil {
			fmt.Printf("%v ->", current.value)
		} else {
			fmt.Printf("%v \n", current.value)
		}
		current = current.next
	}
}
