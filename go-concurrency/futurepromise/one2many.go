package main

import (
	"fmt"
	"math/rand"
	"time"
)

// future promise的场景是发送请求后，resp并不立即返回，而是再未来的某个时间返回
// 请求响应的另一个场景是：
// 一个请求可以发送给多个接受者，而发送者只接收最近的一个返回，剩下不管，
// 例子中的channel被当作入参or返回值，来进行routine的通信

//请求接收者只负责写数据，
func receiver(ch chan<- int32, times int) {
	//ra, rb := rand.Int31(), rand.Intn(5) + 1
	ra := rand.Int31()
	// Sleep 1s/2s/3s.
	time.Sleep(time.Duration(times+1) * 3 * time.Second)
	ch <- ra
}

type msg struct {
	success bool
	info    int32
}

func receiverMore(ch chan<- msg) {
	//ra, rb := rand.Int31(), rand.Intn(5) + 1
	ra := rand.Intn(3)
	// Sleep 1s/2s/3s.
	time.Sleep(time.Duration(ra+1) * 3 * time.Second)
	//根据业务逻辑，写入msg以及成功与否标志
	ch <- msg{success: true, info: 123}
}

func main() {
	//要创建一个带有缓冲区的channel
	ch := make(chan int32, 5)
	//启动多个routine去请求
	for i := 0; i < cap(ch); i++ {
		go receiver(ch, i)
	}
	// 一旦接收，直接打印
	rnd := <-ch
	fmt.Println(rnd)
	//这是ch里的第二个值
	/*rnd2 := <- ch
	fmt.Println(rnd2)*/

}
