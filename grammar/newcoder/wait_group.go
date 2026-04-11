package newcoder

import (
	"fmt"
	"sync"
)

func WaitGroupTest() {
	values := []int{10, 20, 30}
	var wg sync.WaitGroup
	wg.Add(len(values))
	for _, val := range values {
		go func() {
			fmt.Printf("%d ", val)
			wg.Done()
		}()
	}
	wg.Wait()
}
