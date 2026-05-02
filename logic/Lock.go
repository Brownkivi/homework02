package logic

import (
	"fmt"
	"sync"
)

func make() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func MuteLock() {
	lock := make()
	var mu sync.Mutex
	mu = sync.Mutex{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				lock()
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	//输出最终结果
	fmt.Println("Final count:", lock())
}
