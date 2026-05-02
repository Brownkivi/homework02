package main

import (
	"fmt"
	"homework02/logic"
	"sync"
	"time"
)

func main() {
	//testModifyInt()
	//testModifySlice()
	//testPrintNum()
	//testScheduler()
	//testShape()
	//testPerson()
	//testChannel()
	testBufferChannel()
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

func testShape() {
	fmt.Println("Testing Shape...")
	circle := logic.Circle{Radius: 5}
	rectangle := logic.Rectangle{Rength: 4, Width: 6}
	fmt.Printf("Circle Area: %.2f\n", circle.Area(circle.Radius, 0)) // Output: Circle Area: 78.50
	fmt.Printf("Circle Perimeter: %.2f\n", circle.Perimeter(circle.Radius, 0))
	fmt.Printf("Rectangle Area: %.2f\n", rectangle.Area(rectangle.Rength, rectangle.Width)) // Output: Rectangle Area: 24.00
	fmt.Printf("Rectangle Perimeter: %.2f\n", rectangle.Perimeter(rectangle.Rength, rectangle.Width))

}

func testPerson() {
	fmt.Println("Testing Person...")
	employee := logic.Employee{
		Person:     logic.Person{Name: "Alice", Age: 30},
		EmployeeID: 12345,
	}
	employee.PrintInfo()
}

func testChannel() {
	fmt.Println("Testing Channel...")
	ch := make(chan int)

	// Wait for the receiver to start
	go logic.Send(ch)
	go logic.Receive(ch)
	time.Sleep(5 * 1000)

	// Wait for a moment to allow goroutines to finish
}

func testBufferChannel() {
	fmt.Println("Testing Buffer Channel...")
	ch := make(chan int, 10)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go logic.Send1(ch, wg)
	go logic.Receive1(ch, wg)
	wg.Wait()
}
