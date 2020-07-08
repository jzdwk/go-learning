package main

// 通知的另一个场景为：
// 可以通过一个无缓冲channel，来控制多个routine的交互
// 一个routine完成后，通知另一个routine

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	// The capacity of the signal channel can
	// also be one. If this is true, then a
	// value must be sent to the channel before
	// creating the following goroutine.

	go func() {
		fmt.Print("Hello")
		// Simulate a workload.
		time.Sleep(time.Second * 2)

		// Receive a value from the done
		// channel, to unblock the second
		// send in main goroutine.
		//这个channel被读取后，main routine 继续执行
		<-done
	}()

	//注意，无缓冲的channel当写入后，则一直阻塞，直到有一个routine去读这个
	done <- struct{}{}
	fmt.Println(" world!")
}
