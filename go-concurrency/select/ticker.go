package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	ch := ticker(time.Second)
	for range ch {
		fmt.Println(time.Since(now))
	}
}

//ticker implement by select, return read only chan
func ticker(d time.Duration) <-chan int {
	ch := make(chan int, 1)
	go func() {
		for {
			//write a value per second, key code, no select is ok
			time.Sleep(d)
			//select {
			//case ch <- 1:
			ch <- 1
			//default:
			//}
		}
	}()
	return ch
}
