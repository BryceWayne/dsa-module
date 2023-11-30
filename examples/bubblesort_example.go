package main

import (
    "dsa/sorting"
    "fmt"
)

func main() {
    arr := []int{22, 45, 12, 8, 10, 6}
    fmt.Println("Original array:", arr)
    sorting.BubbleSort(arr)
    fmt.Println("Sorted array:", arr)
}
