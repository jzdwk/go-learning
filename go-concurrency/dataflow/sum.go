package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// sum 1-100

type task struct {
	begin int
	end   int
}

func main() {
	//allot task
	total := 10000
	var taskArr []task
	for i := 0; i < total/10; i++ {
		t := new(task)
		//task{begin:i*10,end:(i+1)*10-1}
		t.begin = i * 10
		t.end = (i+1)*10 - 1
		taskArr = append(taskArr, *t)
	}
	taskArr = append(taskArr, task{100000, 100000})
	wg := sync.WaitGroup{}
	wg.Add(len(taskArr))
	rst := make([]int, len(taskArr))
	begin := time.Now()
	for i := 0; i < len(taskArr); i++ {
		go doTask(&wg, &taskArr[i], rst, i)
	}
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
	since := time.Since(begin)
	sum := 0
	for _, v := range rst {
		sum += v
	}
	fmt.Println(sum)
	fmt.Println(since)
}

func doTask(wg *sync.WaitGroup, t *task, rst []int, index int) {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}
	rst[index] = sum
	wg.Done()
}
