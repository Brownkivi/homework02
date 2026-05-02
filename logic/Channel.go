package logic

import (
	"fmt"
)

func Send(ch chan int) {

	for i := 0; i < 10; i++ {
		fmt.Printf("Sent: %d\n", i)
		ch <- i
	}
	close(ch)
}

func Receive(ch chan int) {

	for num := range ch {
		fmt.Printf("Received: %d\n", num)
	}

}
