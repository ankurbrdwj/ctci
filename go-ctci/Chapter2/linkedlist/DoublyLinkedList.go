package linkedlist

import (
	"github.com/ankurbrdwj/ctci/go-ctci/Chapter2/node"
)

//DoublyLinkedList struct type
type DoublyLinkedList struct {
	head *node.Node
	tail *node.Node
}

//New Constructor with no params
func New() DoublyLinkedList {
	return DoublyLinkedList{
		head: nil,
		tail: nil,
	}
}

//NewNode Constructor with node param
func NewNode(n *node.Node) DoublyLinkedList {
	return DoublyLinkedList{
		head: n,
		tail: n,
	}
}

func (d *DoublyLinkedList) SetHead(n *node.Node) {
	d.head = n
}

func (d *DoublyLinkedList) SetTail(n *node.Node) {
	d.tail = n
}

func (d *DoublyLinkedList) InsertBefore(n *node.Node, nodeToInsert *node.Node) {
	if nodeToInsert == d.head && nodeToInsert == d.tail {
		return
	}
	d.Remove(nodeToInsert)
	nodeToInsert.SetPrev(*n.Prev())
	nodeToInsert.SetNext(*n)
	// if insertion is before head
	if n.Prev() == d.head {
		d.head = nodeToInsert
	} else {
		n.Prev().SetNext(*nodeToInsert)
	}
	n.SetPrev(*nodeToInsert)
}

func (d *DoublyLinkedList) InsertAfter(n *node.Node, nodeToInsert *node.Node) {
	// Write your code here.

}

func (d *DoublyLinkedList) InsertAtPosition(position int, nodeToInsert *node.Node) {
	// Write your code here.

}

func (d *DoublyLinkedList) RemoveNodesWithValue(value interface{}) {
	n := d.head
	for n != nil {
		nodeToRemove := n
		n = n.Next()
		if nodeToRemove.Value() == value {
			d.Remove(nodeToRemove)
		}

	}
}

func (d *DoublyLinkedList) Remove(n *node.Node) {
	if n == d.head {
		d.head = d.head.Next()
	}
	if n == d.tail {
		d.tail = d.tail.Prev()
	}
	d.RemoveNodeBindings(n)

}

func (d *DoublyLinkedList) RemoveNodeBindings(n *node.Node) {
	if n.Prev() != nil {
		n.Prev().SetNext(*n.Next())
	}
	if n.Next() != nil {
		n.Next().SetPrev(*n.Prev())
	}
	n.SetPrev(node.New())
	n.SetNext(node.New())
}

func (d *DoublyLinkedList) ContainsNodeWithValue(value interface{}) bool {
	n := d.head // runner for list
	for n != nil && n.Value() != value {
		n = n.Next()
	}
	return n != nil
}
