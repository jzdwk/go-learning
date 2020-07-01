package main

import (
	"fmt"
	"time"
)

//定时调度的通知模型

//返回一个只读channel
func AfterDuration(d time.Duration) <-chan struct{} {
	c := make(chan struct{}, 1)
	go func() {
		time.Sleep(d)
		c <- struct{}{}
	}()
	return c
}

func main1() {
	fmt.Println("Hi!")
	//阻塞直至传入的time.Second，schedule的控制由main goroutine控制
	<-AfterDuration(time.Second)
	fmt.Println("Hello!")
	<-AfterDuration(time.Second)
	fmt.Println("Bye!")
}

func main2() {
	fmt.Println("Hi!")
	//阻塞直至传入的time.Second，schedule的控制由main goroutine控制
	ch := make(chan struct{}, 1)
	d := time.Second
	fmt.Println("Hello!")
	go func(ch chan struct{}, d time.Duration) {
		//sleep duration
		time.Sleep(d)
		//write
		ch <- struct{}{}
	}(ch, d)
	<-ch
	fmt.Println("Bye!")
}
