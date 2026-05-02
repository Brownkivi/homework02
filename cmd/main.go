package main

import (
	"fmt"
	"homework02/logic"
)

func main() {
	//testModifyInt()
	//testModifySlice()
	//testPrintNum()
	testScheduler()
}

func testModifyInt() {
	fmt.Println("Testing ModifyInt...")
	a := 5
	fmt.Println("Before modification:", a) // Output: 5
	logic.ModifyInt(&a)
	fmt.Println("After modification:", a) // Output: 15
}

func testModifySlice() {
	fmt.Println("Testing ModifySlice...")
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("Before modification:", s) // Output: [1 2 3 4 5]
	logic.ModifySlice(s)
	fmt.Println("After modification:", s) // Output: [2 4 6 8 10]
}

func testPrintNum() {
	fmt.Println("Testing PrintNum...")
	logic.PrintNum()
}

func testScheduler() {
	fmt.Println("Testing Scheduler...")
	tasks := logic.BuildTasks()
	logic.Scheduler(tasks)
}

func testScheduler2() {
	fmt.Println("Testing Scheduler2...")
	var tasks []logic.DoTask
	for i := 0; i < 5; i++ {
		task := new(logic.Task)
		task.Id = i
		tasks = append(tasks, *task)
	}
	logic.Scheduler2(tasks)

}
