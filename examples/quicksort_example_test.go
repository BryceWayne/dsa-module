package main

import (
    "dsa/sorting"
    "reflect"
    "testing"
)

func TestQuickSort(t *testing.T) {
    arr := []int{10, 5, 3, 8, 6, 7}
    sortedArr := []int{3, 5, 6, 7, 8, 10}
    sorting.QuickSort(arr)
    if !reflect.DeepEqual(arr, sortedArr) {
        t.Errorf("QuickSort failed, got: %v, want: %v.", arr, sortedArr)
    }
}
