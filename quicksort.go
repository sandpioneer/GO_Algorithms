package main

import (
	"fmt"
)

func partition(arr []int, left, right int) int {
	pivot := arr[left]
	for left < right {
		for left < right && arr[right] >= pivot {
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] <= pivot {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = pivot
	return left
}

func quick_sort(arr []int, left, right int) {
	if left >= right {
		return
	} else {
		position := partition(arr, left, len(arr)-1)
		quick_sort(arr, left, position-1)
		quick_sort(arr, position+1, right)
	}

}

func main() {
	arr := []int{2, 0, 1, 3, 2, 4}
	quick_sort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
