package main

import "fmt"

type fibStt struct {
	first  int
	second int
}
type Fib1 chan fibStt
type Fib2 chan fibStt

func main() {
	//init
	fib1 := make(Fib1, 1)
	fib2 := make(Fib2, 1)
	fibStt := fibStt{1, 1}
	//写入channel1
	fib1 <- fibStt
	//调用两个routine去处理
	go fib1.doFib("goroutine_1:", &fib2)
	go fib2.doFib("goroutine_2:", &fib1)
}

func (f *Fib1) doFib(name string, fib2 *Fib2) {
	for {
		//从f1中取出值
		fibStt := <-*f
		tmp := fibStt.first
		first := fibStt.second
		second := tmp + first
		fmt.Println(name + " generate " + string(second))
		fibStt.first = first
		fibStt.second = second
		//写到f2
		*fib2 <- fibStt
	}

}

func (f *Fib2) doFib(name string, fib1 *Fib1) {
	for {
		//从f2中取出值
		fibStt := <-*f
		tmp := fibStt.first
		first := fibStt.second
		second := tmp + first
		fmt.Println(name + " generate " + string(second))
		fibStt.first = first
		fibStt.second = second
		//写到f1
		*fib1 <- fibStt
	}

}
