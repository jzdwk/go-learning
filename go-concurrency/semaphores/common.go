package main

import (
	"log"
	"math/rand"
	"time"
)

//semaphores和mutex的区别，mutex的上锁和解锁都是由同一个线程完成(看例子)，因此使用场景限制于临界区
//semaphores(信号量)可应用于生产-消费此类模型，并不局限于哪个线程去上锁/解锁

//下面例子展示了使用一个大小为n的channel作为维护资源数的semaphores

type Seat int
type Bar chan Seat

//
func ServeCustomerAtSeat(c int, seat Seat, bar *Bar) {
	log.Print("customer#", c, " drinks at seat#", seat)
	time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
	log.Print("<- customer#", c, " frees seat#", seat)
	*bar <- seat // free seat and leave the bar
}

//channel的
func main() {
	rand.Seed(time.Now().UnixNano())
	//容量10的chan
	bar24x7 := make(Bar, 10)
	//向bar内摆满座位
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		bar24x7 <- Seat(seatId)
	}
	//不停的上顾客，当bar满座后，新的顾客只能阻塞，等bar中有顾客离开，后者再填充
	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		// Need a seat to serve next customer.
		//如果没有座位（说明人满），则阻塞
		seat := <-bar24x7
		go ServeCustomerAtSeat(customerId, seat, &bar24x7)
	}
	for {
		time.Sleep(time.Second)
	}
}
