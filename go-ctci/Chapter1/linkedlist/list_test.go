package linkedlist

import (
	"testing"
)

func TestContainsWithValue(t *testing.T) {
	list := fillLinkedList(List{}, 11)
	result := list.ContainsNodeWithValue(10)
	if !result {
		t.Errorf("TestContainsWithValue Failed expected %t but found %t", true, result)
	}

}
func TestRemoveNodeBindings(t *testing.T) {
	list := fillLinkedList(List{}, 11)
	item := list.head.next
	list.RemoveNodeBindings(item)
	if item.prev != nil {
		t.Errorf("Test RemoveNodeBindings Failed expected %v but found %v", nil, item.prev)
	}
	if item.next != nil {
		t.Errorf("Test RemoveNodeBindings Failed expected %v but found %v", nil, item.next)
	}
}

func TestRemove(t *testing.T) {
	list := fillLinkedList(List{}, 11)
	value := list.head.next.value
	item := list.head.next
	list.Remove(item)
	if list.ContainsNodeWithValue(value) {
		t.Errorf("TestRemove Failed expected %v but found %v", false, value)
	}
}
func TestRemoveNodesWithValue(t *testing.T) {
	list := fillLinkedList(List{}, 11)
	value := list.head.next.value
	list.RemoveNodesWithValue(value)
	if list.ContainsNodeWithValue(value) {
		t.Errorf("TestRemoveNodesWithValue Failed expected %v but found %v", false, value)
	}
}

func TestInsertAtPosition(t *testing.T) {
	list := fillLinkedList(List{}, 11)
	item := list.head.next
	list.InsertAtPosition(8, item)
	runner := list.head
	for i := 1; i < 7; i++ {
		runner = runner.next
	}
	if item.value != runner.value {
		t.Errorf("TestInsertAtPosition Failed expected %v but found %v", item.value, runner.value)
	}
}

func TestInsertAfter(t *testing.T) {
	list := fillLinkedList(List{}, 11)
	item := list.head.next
	list.InsertAfter(list.tail, item)
	runner := list.head
	for i := 1; i < 11; i++ {
		runner = runner.next
	}
	if item.value != runner.value {
		t.Errorf("TestInsertAtPosition Failed expected %v but found %v", item.value, runner.value)
	}
}

func fillLinkedList(l List, length int) List {
	for i := 1; i < 11; i++ {
		node := &Node{value: i}
		if l.head == nil {
			l.head = node
		} else {
			l.tail.next = node
			node.prev = l.tail
		}

		l.tail = node
	}
	return l
}
