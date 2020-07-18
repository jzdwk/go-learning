/*
@Time : 20-7-8
@Author : jzd
@Project: go-learning
*/
package main

import (
	"fmt"
	"os"
	"time"
)

type Ball uint64

func main() {
	table := make(chan Ball, 1)
	go func() {
		table <- 1
	}()
	go Play("A:", table)
	Play("B:", table)

}

func Play(playerName string, table chan Ball) {
	var lastValue Ball = 1
	for {
		ball := <-table // get the ball
		fmt.Println(playerName, ball)
		ball += lastValue
		if ball < lastValue { // overflow
			os.Exit(0)
		}
		lastValue = ball
		time.Sleep(time.Second)
		table <- ball // bat back the ball
	}
}
