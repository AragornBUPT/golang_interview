package math_app

func CountGoodTriplets(arr []int, a int, b int, c int) int {
	length := len(arr)
	if length <= 2 {
		return 0
	}

	res := 0
	for i := 0; i <= length-3; i++ {
		for j := i + 1; j <= length-2; j++ {
			for k := j + 1; k <= length-1; k++ {
				if abs(arr[i], arr[j]) <= a && abs(arr[j], arr[k]) <= b && abs(arr[i], arr[k]) <= c {
					res++
				}
			}
		}
	}

	return res
}

// abs 计算两个整数之差的绝对值
func abs(a, b int) int {
	res := a - b
	if res < 0 {
		return -res
	}

	return res
}
