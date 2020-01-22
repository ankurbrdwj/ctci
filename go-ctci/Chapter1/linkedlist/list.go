package linkedlist

type List struct {
	head *Node
	tail *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (d *List) SetHead(n *Node) {
	d.head = n
}

func (d *List) SetTail(n *Node) {
	d.tail = n
}

func (d *List) InsertBefore(n *Node, nodeToInsert *Node) {
	if nodeToInsert == d.head && nodeToInsert == d.tail {
		return
	}
	d.Remove(nodeToInsert)
	nodeToInsert.next = n
	nodeToInsert.prev = n.Prev()

	// if insertion is before head
	if n.Prev() == d.head {
		d.head = nodeToInsert
	} else {
		n.Prev().next = nodeToInsert
	}
	n.prev = nodeToInsert
}

func (d *List) InsertAfter(n *Node, nodeToInsert *Node) {
	// Write your code here.

}

func (d *List) InsertAtPosition(position int, nodeToInsert *Node) {
	n := d.head
	for i := 1; i < position; i++ {
		n = n.Next()
	}
	d.InsertBefore(n, nodeToInsert)
}

func (d *List) RemoveNodesWithValue(value interface{}) {
	n := d.head
	for n != nil {
		nodeToRemove := n // refer before removing
		n = n.Next()      // reassign runner
		if nodeToRemove.value == value {
			d.Remove(nodeToRemove)
		}

	}
}

func (d *List) Remove(n *Node) {
	if n == d.head {
		d.head = d.head.Next()
	}
	if n == d.tail {
		d.tail = d.tail.Prev()
	}
	d.RemoveNodeBindings(n)

}

func (d *List) RemoveNodeBindings(n *Node) {
	if n.Prev() != nil {
		n.Prev().next = n.next
	}
	if n.Next() != nil {
		n.Next().prev = n.prev
	}
	n.prev = nil
	n.next = nil
}

func (d *List) ContainsNodeWithValue(value interface{}) bool {
	n := d.head // runner for list
	for n != nil && n.value != value {
		n = n.Next()
	}
	return n != nil
}
