package main

import (
    "dsa/sorting"
    "reflect"
    "testing"
)

func TestBubbleSort(t *testing.T) {
    arr := []int{22, 45, 12, 8, 10, 6}
    sortedArr := []int{6, 8, 10, 12, 22, 45}
    sorting.BubbleSort(arr)
    if !reflect.DeepEqual(arr, sortedArr) {
        t.Errorf("BubbleSort failed, got: %v, want: %v.", arr, sortedArr)
    }
}
