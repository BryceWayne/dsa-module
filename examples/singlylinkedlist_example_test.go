package main

import (
    "dsa/linkedlists"
    "testing"
)

func TestSinglyLinkedList(t *testing.T) {
    list := linkedlists.NewSinglyLinkedList()

    list.Append(1)
    list.Append(2)
    list.Prepend(0)

    if list.Head.Value != 0 || list.Head.Next.Value != 1 || list.Head.Next.Next.Value != 2 {
        t.Errorf("SinglyLinkedList Append/Prepend failed")
    }

    list.DeleteWithValue(1)
    if list.Head.Next.Value != 2 {
        t.Errorf("SinglyLinkedList DeleteWithValue failed")
    }
}
