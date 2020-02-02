package com.java.ctci.lists;

public class DoublyLinkedList {
    public Node head;
    public Node tail;

    public void setHead(Node node) {
        // Write your code here.
    }

    public void setTail(Node node) {
        // Write your code here.
    }

    public void insertBefore(Node node, Node nodeToInsert) {
        // Write your code here.
    }

    public void insertAfter(Node node, Node nodeToInsert) {
        // Write your code here.
    }

    public void insertAtPosition(int position, Node nodeToInsert) {
        // Write your code here.
    }

    public void removeNodesWithValue(int value) {
        // Write your code here.
    }

    public void remove(Node node) {
        // Write your code here.
    }

    public boolean containsNodeWithValue(int value) {
        // Write your code here.
        Node node = this.head;
        while (node.getData() != value) {
            node = node.getNext();
        }
        return node.getData() == value;
    }
}