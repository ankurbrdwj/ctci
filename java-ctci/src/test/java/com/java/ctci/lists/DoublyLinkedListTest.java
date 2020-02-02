package com.java.ctci.lists;

import static org.junit.jupiter.api.Assertions.*;

import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;

class DoublyLinkedListTest {

    private static DoublyLinkedList list = new DoublyLinkedList();

    @BeforeAll
    public static void setUp() {
        list.head = new Node(1);
        Node runner =list.head ;
        for (int i = 1; i < 11; i++) {
            Node node = new Node(i);
			runner.setNext(node);
			node.setPrev(runner);
			runner = runner.getNext();
        }
        list.tail =runner;
    }

    @Test
    void testContainsNodeWithValue() {
        assertTrue(list.containsNodeWithValue(4));
    }

}
