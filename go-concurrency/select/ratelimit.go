package main

import (
	"fmt"
	"time"
)

type request interface {
}

func handle(r request) {
	fmt.Println(r.(int))
}

//200 request in one minute
const rateLimitPeriod = time.Minute
const rateLimit = 200

func handleRequests(requests <-chan request) {
	quotas := make(chan time.Time, rateLimit)

	go func() {
		//core code by use ticker, when ticker, put v to chan
		tick := time.NewTicker(rateLimitPeriod / rateLimit)
		defer tick.Stop()
		for t := range tick.C {
			select {
			case quotas <- t:
			default:
			}
		}
	}()
	//read from chan and do handle
	for r := range requests {
		<-quotas
		go handle(r)
	}
}

func main() {
	//go timerTest()
	//go tickerTest()

	requests := make(chan request)
	go handleRequests(requests)
	// time.Sleep(time.Minute)
	for i := 0; ; i++ {
		requests <- i
	}
}

func timerTest() {
	t := time.NewTimer(time.Second * 2)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			fmt.Println("test timer")
			//if no reset, timer will put v to t.C once and the goroutine will be asleep
			t.Reset(time.Second * 2)
		}
	}
}

func tickerTest() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fmt.Println("test ticker")
		default:
		}
	}
}
