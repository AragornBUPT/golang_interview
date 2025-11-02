package leetcode

func MinCost(nums []int) int {
	length := len(nums)
	if length < 3 {
		return max(nums[0], nums[1])
	}

	costSum := 0
	rest := nums[0]
	i := 1
	for ; i <= length-2; i += 2 {
		a := rest
		b := nums[i]
		c := nums[i+1]
		rest_tmp, cost := getCostAndRest(a, b, c)
		rest = rest_tmp
		costSum += cost
	}

	costSum += rest
	if i == length-1 {
		costSum += nums[i]
		return costSum
	}

	return costSum
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func getCostAndRest(a, b, c int) (int, int) {
	if a >= b && b >= c {
		return a, c
	} else if a >= c && c >= b {
		return a, b
	} else if b >= a && a >= c {
		return b, c
	} else if b >= c && c >= a {
		return b, a
	} else if c >= a && a >= b {
		return c, b
	} else {
		return c, a
	}
}