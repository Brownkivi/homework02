package logic

import (
	"fmt"
	"sync"
)

func Send1(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		fmt.Printf("Sent: %d\n", i)
		ch <- i
	}
	close(ch)
}

func Receive1(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("Received: %d\n", num)
	}
}
