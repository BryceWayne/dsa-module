package main

import (
    "dsa/sorting"
    "fmt"
)

func main() {
    arr := []int{32, 45, 67, 2, 7}
    fmt.Println("Original array:", arr)
    sorting.MergeSort(arr)
    fmt.Println("Sorted array:", arr)
}
