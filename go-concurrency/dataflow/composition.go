package main

import (
	"fmt"
	"math/rand"
	"time"
)

func composer(inA, inB <-chan int64) <-chan int64 {
	output := make(chan int64)
	go func() {
		for {
			a1, b, a2 := <-inA, <-inB, <-inA
			fmt.Println(fmt.Sprintf("a1: %d, b1: %d, a2: %d", a1, b, a2))
			output <- a1 + a2 - b
		}
	}()
	return output
}

func main() {
	inA, inB := make(chan int64), make(chan int64)
	//gen random A
	go func() {
		for {
			inA <- rand.Int63n(10)
			time.Sleep(time.Second * 1)
			inA <- rand.Int63n(10)
		}
	}()
	//gen random B
	go func() {
		for {
			inB <- rand.Int63n(10)
			time.Sleep(time.Second * 1)
		}
	}()
	go func() {
		out := composer(inA, inB)
		for {
			select {
			case v := <-out:
				fmt.Println("result is ", v)
			default:
			}
		}
	}()
	select {}
}
