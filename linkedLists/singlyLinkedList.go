package linkedlists

import "fmt"

// Node represents a node in the singly linked list.
type Node struct {
    Value int
    Next  *Node
}

// SinglyLinkedList represents a singly linked list.
type SinglyLinkedList struct {
    Head *Node
}

// NewSinglyLinkedList creates a new empty singly linked list.
func NewSinglyLinkedList() *SinglyLinkedList {
    return &SinglyLinkedList{}
}

// Append adds a new node with the given value at the end of the list.
func (sll *SinglyLinkedList) Append(value int) {
    newNode := &Node{Value: value}
    if sll.Head == nil {
        sll.Head = newNode
        return
    }
    current := sll.Head
    for current.Next != nil {
        current = current.Next
    }
    current.Next = newNode
}

// Prepend adds a new node with the given value at the beginning of the list.
func (sll *SinglyLinkedList) Prepend(value int) {
    newNode := &Node{Value: value, Next: sll.Head}
    sll.Head = newNode
}

// DeleteWithValue deletes the first node with the given value.
func (sll *SinglyLinkedList) DeleteWithValue(value int) {
    if sll.Head == nil {
        return
    }
    if sll.Head.Value == value {
        sll.Head = sll.Head.Next
        return
    }
    current := sll.Head
    for current.Next != nil {
        if current.Next.Value == value {
            current.Next = current.Next.Next
            return
        }
        current = current.Next
    }
}

// Print prints all the values in the linked list.
func (sll *SinglyLinkedList) Print() {
    current := sll.Head
    for current != nil {
        fmt.Print(current.Value, " ")
        current = current.Next
    }
    fmt.Println()
}
