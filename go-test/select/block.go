package main

import (
	"fmt"
	"time"
)

func blockTest() {
	start := time.Now()
	c := make(chan interface{})
	ch1 := make(chan int)
	ch2 := make(chan int)
	//4s close
	go func() {

		time.Sleep(5 * time.Second)
		close(c)
	}()
	//3s ch1填值
	go func() {
		time.Sleep(4 * time.Second)
		ch1 <- 4
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 3
	}()

	fmt.Println("Blocking on read...")
	select {
	//如果一个或者多个channel可用，则随机选取执行，否则看是否有default
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	case <-ch1:
		fmt.Printf("ch1 case...")
	case <-ch2:
		fmt.Printf("ch2 case...")
		//注释掉default,case将阻塞，直到执行最先可用的case（3s后的ch2可读）
		/*default:
		fmt.Printf("default go...")*/
	}
}
