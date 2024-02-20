package main

import (
	"fmt"
	"testing"
)

type Node struct {
	data int
	next *Node
}

var genius *Node

func NewList(data int) *Node {
	genius = &Node{data: data}
	return genius
}

func (n *Node) Delete() string {
	for i := genius; i != nil; i = i.next {
		if i.data == n.data {
			i.next = i.next.next
		}
	}
	return fmt.Sprintf("Deleting %d", n.data)
}

// Insert items at a give position
func (n *Node) Insert(data int, position int) string {
	newNode := &Node{data: data}
	for i := genius; i != nil; i = i.next {
		if i.data == position {
			tmp := i.next
			i.next = newNode
			newNode.next = tmp
		}
	}
	return fmt.Sprintf("Inserting %d at position %d", data, position)
}

func Test_LinkedList(t *testing.T) {
	fmt.Println("Testing LinkedList")

	genius := NewList(1)

	second := &Node{data: 2}
	third := &Node{data: 3}

	genius.next = second
	genius.next.next = third

	for n := genius; n != nil; n = n.next {
		fmt.Println(n.data)
	}

	second.Delete()

	for n := genius; n != nil; n = n.next {
		fmt.Println(n.data)
	}

	genius.Insert(3, 2)
	genius.Insert(4, 3)
	genius.Insert(5, 4)
	genius.Insert(6, 5)
	genius.Insert(99, 3)

	// create func to loop over the list and print
	for n := genius; n != nil; n = n.next {
		fmt.Println("Afer insert: ", n.data)
	}

}
