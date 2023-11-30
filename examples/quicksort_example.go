package main

import (
    "dsa/sorting"
    "fmt"
)

func main() {
    arr := []int{10, 5, 3, 8, 6, 7}
    fmt.Println("Original array:", arr)
    sorting.QuickSort(arr)
    fmt.Println("Sorted array:", arr)
}
