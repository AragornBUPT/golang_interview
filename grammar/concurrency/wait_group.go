package concurrency

import (
	"fmt"
	"sync"
)

// TestWaitGroup 演示WaitGroup的使用以及Go语言中for循环变量闭包的问题
func TestWaitGroup() {
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
