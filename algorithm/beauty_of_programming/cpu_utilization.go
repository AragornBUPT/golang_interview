package beauty_of_programming

import (
	"runtime"
	"sync"
	"time"
)

func UtilCpu() {
	numCpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numCpu)

	var wg sync.WaitGroup
	wg.Add(numCpu)

	times := 0
	for i := 0; i < numCpu; i++ {
		go func() {
			defer wg.Done()
			for {
				for i := 0; i <= 100000; i++ {
					times++
				}
				time.Sleep(10 * time.Microsecond)
			}
		}()
	}
	wg.Wait()
}
