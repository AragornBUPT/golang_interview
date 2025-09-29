package main

import "fmt"

func heapify(arr []int, parent, lastIndex int) {
	max := parent
	left := max*2 + 1
	right := max*2 + 2

	if left <= lastIndex && arr[left] > arr[max] {
		max = left
	}

	if right <= lastIndex && arr[right] > arr[max] {
		max = right
	}

	if max != parent {
		arr[max], arr[parent] = arr[parent], arr[max]
		heapify(arr, max, lastIndex)
	}
}

func buildHeap(arr []int, lastIndex int) {
	parent := (lastIndex - 1) / 2
	for parent >= 0 {
		heapify(arr, parent, lastIndex)
		parent--
	}
}

func HeapSort(arr []int) {
	length := len(arr)
	for lastIndex := length - 1; lastIndex >= 1; lastIndex-- {
		buildHeap(arr, lastIndex)
		arr[lastIndex], arr[0] = arr[0], arr[lastIndex]
	}
}

func main() {
	arr := []int{5, 2, 3, 1, 9, 5, 3}
	fmt.Println(arr)

	HeapSort(arr)
	fmt.Println(arr)
	// fmt.Println(5/2)
}
