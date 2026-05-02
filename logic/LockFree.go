package logic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int64 = 0

func add() {
	atomic.AddInt64(&count, 1)
}

func LockFree() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				add()
			}
		}()
	}
	wg.Wait()
	//输出最终结果
	fmt.Println("Final count:", count)
}
