/* 
https://blog.csdn.net/longfeng995/article/details/140817438
 */

package sort

// import "fmt"

func partition(arr []int, low int, high int) int {
	flag := arr[high]
	slow := low

	for i := low; i < high; i++ {
		if arr[i] <= flag {
			arr[i], arr[slow] = arr[slow], arr[i]
			slow++
		}
	}

	arr[slow], arr[high] = arr[high], arr[slow]

	return slow
}

func QuickSort(arr []int, low, high int) {
	if low < high {
		index := partition(arr, low, high)
		QuickSort(arr, 0, index-1)
		QuickSort(arr, index+1, high)
	}
}

// func main() {
// 	arr := []int{5, 2, 3, 1}
// 	fmt.Println(arr)

// 	QuickSort(arr, 0, len(arr)-1)
// 	fmt.Println(arr)
// }
