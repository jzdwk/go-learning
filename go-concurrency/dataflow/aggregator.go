package main

import (
	"fmt"
)

func aggregator(inputs ...<-chan uint64) <-chan uint64 {
	output := make(chan uint64)
	//var wg sync.WaitGroup
	for _, in := range inputs {
		//wg.Add(1)
		go func(int <-chan uint64) {
			for {
				x, ok := <-in
				if ok {
					output <- x
				} else {
					//wg.Done()
					close(output)
				}
			}
		}(in)
	}
	/*go func() {
		wg.Wait()
		close(output)
	}()*/
	return output
}

func main() {
	c := make(chan uint64, 10)
	var i uint64
	for i = 0; i < 10; i++ {
		c <- i
		i++
	}
	for {
		select {
		case i := <-aggregator(c):
			fmt.Println(i)
		default:
			fmt.Println("wait...")
		}
	}
}
