package main

import "fmt"

func main() {
	foo, bar := make(chan struct{}, 1), make(chan struct{}, 1)
	//if close, all chan can <- with no block
	close(foo)
	close(bar) // for demo purpose
	x, y := 0.0, 0.0
	f := func() { x++ }
	g := func() { y++ }
	for i := 0; i < 100000; i++ {
		select {
		case <-foo:
			f()
		case <-foo:
			f()
		case <-bar:
			g()
		}
	}
	fmt.Println(x / y) // about 2
}
