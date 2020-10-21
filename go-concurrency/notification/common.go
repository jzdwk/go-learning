package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
	_ "time"
)

// 通知可以被看作是特殊的请求/回应用例。
// 主逻辑通过声明一个无缓冲channel，并将其作为参数传给某个routine，然后继续执行自身逻辑
// routine中完成自身逻辑后向channel写入
// 主逻辑最后通过读取通知来执行routine结果

func main4() {
	values := make([]byte, 32*1024*1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	done := make(chan struct{}) // 也可以是缓冲的

	// 排序协程
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		done <- struct{}{} // 通知排序已完成
	}()

	// 并发地做一些其它事情...

	<-done // 等待通知
	fmt.Println(values[0], values[len(values)-1])
}

func main() {
	done := make(chan int, 1)
	handler(func() {
		defer close(done)
		fmt.Println("main sleep...")
		time.Sleep(time.Second * 3)
	})
	go func() {
		<-done
		fmt.Println("routine get...")
	}()
	time.Sleep(time.Second * 10)
}

func handler(handle func()) {
	handle()
}
