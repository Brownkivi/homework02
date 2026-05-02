package logic

import (
	"fmt"
	"time"
)

func PrintNum() {
	fmt.Println("Printing numbers from 1 to 10...")
	go printOdd()
	go printEven()
	time.Sleep(2 * time.Second) // 等待 goroutine 完成
	fmt.Println("Done printing numbers.")

}

func printOdd() {
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println("奇数: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func printEven() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("偶数: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
