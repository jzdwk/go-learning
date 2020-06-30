package main

import (
	"fmt"
	"math/rand"
	"time"
)

// future promise的场景是发送请求后，resp并不立即返回，而是再未来的某个时间返回
// 请求响应的其中一个场景是：
// 一个请求对应一个响应，例子中的channel被当作入参or返回值，来进行routine的通信

func longTimeRequest1() <-chan int32 {
	r := make(chan int32)
	//这个只读channel在3s后会被塞入一个值，因此，3s内的读入将阻塞
	go func() {
		// Simulate a workload.3s的工作负载
		time.Sleep(time.Second * 3)
		r <- rand.Int31n(100)
	}()
	return r
}

// 定义一个只写的channel
func longTimeRequest2(ch chan<- int32) {
	go func() {
		//模拟5s的工作负载
		time.Sleep(time.Second * 5)
		//写入channel
		ch <- rand.Int31n(100)
	}()
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//因为2个chan都是只读的，因此<-a和<-b都会阻塞至读到数据
	a, b := longTimeRequest1(), longTimeRequest1()
	fmt.Println(sumSquares(<-a, <-b))
	//另一种方式
	c, d := make(chan int32), make(chan int32)
	//并发写
	go longTimeRequest2(c)
	go longTimeRequest2(d)
	fmt.Println(sumSquares(<-c, <-d))
}
