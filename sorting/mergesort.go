package sorting

func MergeSort(arr []int) {
    if len(arr) > 1 {
        mid := len(arr) / 2
        left := make([]int, mid)
        right := make([]int, len(arr)-mid)

        copy(left, arr[:mid])
        copy(right, arr[mid:])

        MergeSort(left)
        MergeSort(right)

        merge(arr, left, right)
    }
}

func merge(arr, left, right []int) {
    var i, j, k int

    for i < len(left) && j < len(right) {
        if left[i] <= right[j] {
            arr[k] = left[i]
            i++
        } else {
            arr[k] = right[j]
            j++
        }
        k++
    }

    for i < len(left) {
        arr[k] = left[i]
        i++
        k++
    }

    for j < len(right) {
        arr[k] = right[j]
        j++
        k++
    }
}
