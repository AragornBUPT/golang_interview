package sort

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
		heapify(arr, max, lastIndex)	// 继续向下整理顺序
	}
}

// 构建大顶堆
func buildHeap(arr []int, lastIndex int) {
	parent := (lastIndex - 1) / 2	// 从堆底元素的父节点开始，逐个调整顺序，使得parent下的所有顺序都是对的
	for parent >= 0 {
		heapify(arr, parent, lastIndex)
		parent--
	}
}

func HeapSort(arr []int) {
	length := len(arr)
	for lastIndex := length - 1; lastIndex >= 1; lastIndex-- {
		buildHeap(arr, lastIndex)
		arr[lastIndex], arr[0] = arr[0], arr[lastIndex]	// 排序，将最小的元素放在最后，然后对剩余元素继续构造大顶堆
	}
}

func main() {
	arr := []int{5, 2, 3, 1, 9, 5, 3}
	fmt.Println(arr)

	HeapSort(arr)
	fmt.Println(arr)
	// fmt.Println(5/2)
}
