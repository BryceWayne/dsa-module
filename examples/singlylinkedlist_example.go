package main

import (
    "dsa/linkedlists"
    "fmt"
)

func main() {
    list := linkedlists.NewSinglyLinkedList()

    list.Append(1)
    list.Append(2)
    list.Prepend(0)
    fmt.Println("After adding elements:")
    list.Print() // Output should be 0 1 2

    list.DeleteWithValue(1)
    fmt.Println("After deleting an element:")
    list.Print() // Output should be 0 2
}
