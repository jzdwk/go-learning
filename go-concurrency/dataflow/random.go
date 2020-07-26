package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
)

func randomGen() <-chan uint64 {
	//no cache
	c := make(chan uint64)
	go func() {
		rnds := make([]byte, 8)
		for {
			_, err := rand.Read(rnds)
			if err != nil {
				close(c)
				break
			}
			fmt.Println("to write")
			c <- binary.BigEndian.Uint64(rnds)
			fmt.Println("write finish")
		}
	}()
	return c
}

func main() {
	c := randomGen()
	fmt.Println(fmt.Sprintf("get random1 %d", <-c))
	fmt.Println(fmt.Sprintf("get random2 %d", <-c))

}
