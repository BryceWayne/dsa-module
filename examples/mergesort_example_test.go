package main

import (
    "dsa/sorting"
    "reflect"
    "testing"
)

func TestMergeSort(t *testing.T) {
    arr := []int{32, 45, 67, 2, 7}
    sortedArr := []int{2, 7, 32, 45, 67}
    sorting.MergeSort(arr)
    if !reflect.DeepEqual(arr, sortedArr) {
        t.Errorf("MergeSort failed, got: %v, want: %v.", arr, sortedArr)
    }
}
