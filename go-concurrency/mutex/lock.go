package main

import (
	"fmt"
	"time"
)

//使用channel作为mutex lock

var counter = 0

func doSum(c chan struct{}) {
	//写入，当多余一个写入，则阻塞，相当于lock
	c <- struct{}{}
	counter++
	//取出值,unlock
	<-c
}

func main() {
	var mutex = make(chan struct{}, 1)

	for i := 0; i < 1000; i++ {
		go doSum(mutex)
	}
	time.Sleep(time.Second * 3)
	fmt.Println("final is ", counter)
}
