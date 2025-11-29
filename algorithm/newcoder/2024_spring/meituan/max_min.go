package main

import (
	"fmt"
)

func main() {
	n := 0
	q := 0

	fmt.Scan(&n, &q)
	sum := 0
	zero_count := 0
	for i := 0; i < n; i++ {
		num := 0
		fmt.Scan(&num)

		if num == 0 {
			zero_count++
		}
		sum += num
	}

	l, r := 0, 0
	for i := 0; i < q; i++ {
		fmt.Scan(&l, &r)
		min := zero_count*l + sum
		max := zero_count*r + sum
		fmt.Println(min, max)
	}
}
