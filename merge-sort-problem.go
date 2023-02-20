package main

import "fmt"

func merge(left, right []int) []int {
    merged := make([]int, len(left)+len(right))
    i := 0
    j := 0
    k := 0

    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            merged[k] = left[i]
            i++
        } else {
            merged[k] = right[j]
            j++
        }
        k++
    }

    for i < len(left) {
        merged[k] = left[i]
        i++
        k++
    }

    for j < len(right) {
        merged[k] = right[j]
        j++
        k++
    }

    return merged
}

func mergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    mid := len(arr) / 2
    left := mergeSort(arr[:mid])
    right := mergeSort(arr[mid:])
    return merge(left, right)
}

func main() {
    arr := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
    fmt.Println("Unsorted array:", arr)

    arr = mergeSort(arr)
    fmt.Println("Sorted array:", arr)
}
