package logic

import (
	"fmt"
	"sync"
	"time"
)

type task func(idx int, wg *sync.WaitGroup)

func Scheduler(tasks []task) {
	var wg sync.WaitGroup
	for idx, task := range tasks {
		wg.Add(1)
		go task(idx, &wg)
	}
	wg.Wait()
	fmt.Println("所有任务执行完成！")
}

func BuildTasks() []task {
	var tasks []task
	for i := 0; i < 5; i++ {
		t := func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			start := time.Now() // 记录开始时间
			fmt.Printf("第%d任务开始执行\n", i)
			time.Sleep(1 * time.Second)   // 模拟任务执行时间
			duration := time.Since(start) // 计算执行耗时
			fmt.Printf("任务执行完成，耗时：%v\n", duration)
		}
		tasks = append(tasks, t)
	}
	return tasks
}
