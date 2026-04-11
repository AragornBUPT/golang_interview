package goexpertprogramming

import (
	"fmt"
)

func SliceExtend() {
	var slice []int
	s1 := append(slice, 1, 2, 3)
	s2 := append(s1, 4)
	fmt.Println(&s1[0] == &s2[0])
}

func SliceExpress() {
	orderLen := 5
	order := make([]uint16, 2*orderLen)

	simpleorder := order[:orderLen]
	simpleorder2 := order[3 : 2*orderLen]
	pollorder := order[:orderLen:orderLen]
	lockorder := order[orderLen:][:orderLen:orderLen]

	fmt.Println("len(simpleorder) = ", len(simpleorder))	// 5
	fmt.Println("cap(simpleorder) = ", cap(simpleorder))	// 10
	fmt.Println("len(simpleorder2) = ", len(simpleorder2))	// 7
	fmt.Println("cap(simpleorder2) = ", cap(simpleorder2))	// 7
	fmt.Println("len(pollorder) = ", len(pollorder))	// 5
	fmt.Println("cap(pollorder) = ", cap(pollorder))	// 5
	fmt.Println("len(lockorder) = ", len(lockorder))	// 5
	fmt.Println("cap(lockorder) = ", cap(lockorder))	// 5
}
