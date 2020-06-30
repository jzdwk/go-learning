package main

import "log"
import "time"

//一个多goroutine和主goroutine的例子，当每一个goroutine
//在生产环境中，常用sync.WaitGroup来完成，而不是本例子

type T = struct{}

//两个入参，第一个ready是只读入参，没有值则阻塞；第二个done是只写，
func worker(id int, ready <-chan T, done chan<- T) {
	<-ready // block here and wait a notification
	log.Print("Worker#", id, " starts.")
	// Simulate a workload.
	time.Sleep(time.Second * time.Duration(id+1))
	log.Print("Worker#", id, " job done.")
	// Notify the main goroutine (N-to-1),
	done <- T{}
}

func main() {
	log.SetFlags(0)

	ready, done := make(chan T), make(chan T)
	go worker(0, ready, done)
	go worker(1, ready, done)
	go worker(2, ready, done)

	// Simulate an initialization phase.
	time.Sleep(time.Second * 3 / 2)
	// 1-to-N notifications.
	ready <- T{}
	ready <- T{}
	ready <- T{}
	// Being N-to-1 notified.
	<-done
	<-done
	<-done
}
