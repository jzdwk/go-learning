/*
@Time : 21-1-4
@Author : jzd
@Project: go-learning
*/
package main

import (
	"context"
	"fmt"
)

func main() {
	p1 := "this is p1"
	p2 := 2021

	list := []funcValue{
		createFuncValue1(p1, p2),
		createFuncValue2(p1, p2),
	}

	for _, o := range list {
		o(context.Background(), p1, p2)
	}

}

type funcValue func(ctx context.Context, p1 string, p2 int)

func createFuncValue1(p1 string, p2 int) funcValue {
	return func(ctx context.Context, subP1 string, subP2 int) {
		fmt.Println(fmt.Sprintf("func1   subP1: %v, subP2: %v, p1 %v, p2 %v", subP1, subP2, p1, p2))

	}
}

func createFuncValue2(p1 string, p2 int) funcValue {
	return func(ctx context.Context, subP1 string, subP2 int) {
		fmt.Println(fmt.Sprintf("func2	subP1: %v, subP2: %v, p1 %v, p2 %v", subP1, subP2, p1, p2))

	}
}
