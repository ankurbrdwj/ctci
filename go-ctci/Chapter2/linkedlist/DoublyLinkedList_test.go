package linkedlist

import (
	"testing"

	"github.com/ankurbrdwj/ctci/go-ctci/Chapter2/node"
)

var data int = 10

func TestContainsWithValue(t *testing.T) {
	var list DoublyLinkedList
	head := node.NewNode(data)
	// runner
	runner := head
	for i := 0; i < 10; i++ {
		temp := node.NewNode(i)
		list.InsertBefore((&runner), &temp)
		next := runner.Next()
		runner = *next
	}
	result := list.ContainsNodeWithValue(10)
	if !result {
		t.Errorf("Test Failed expected %t but found %t", true, result)
	}
}
