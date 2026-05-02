package logic

import (
	"fmt"
	"sync"
	"time"
)

type DoTask interface {
	doTask(wg *sync.WaitGroup)
}

type Task struct {
	Id int
}

func (t Task) doTask(wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now() // 记录开始时间
	fmt.Printf("任务开始执行")
	time.Sleep(1 * time.Second)   // 模拟任务执行时间
	duration := time.Since(start) // 计算执行耗时
	fmt.Printf("任务执行完成，耗时：%v\n", duration)
}

func Scheduler2(tasks []DoTask) {
	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go task.doTask(&wg)
	}
	wg.Wait()
	fmt.Println("所有任务执行完成！")

}
