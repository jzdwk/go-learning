package main

import (
	"fmt"
	"strconv"
)

type book struct {
	id   int
	name string
}

func main() {
	//cap 3
	bookChan := make(chan book, 3)
	//select put
	for i := 0; i < cap(bookChan)*2; i++ {
		select {
		//write
		case bookChan <- book{i, "book_" + strconv.Itoa(i)}:
			fmt.Println("put book_" + strconv.Itoa(i))
		default:
			fmt.Println("put book_" + strconv.Itoa(i) + " failed")
		}
	}

	//select read
	for {
		select {
		//read
		case book := <-bookChan:
			fmt.Println("get book_" + strconv.Itoa(book.id))
		default:
			fmt.Println("get book failed")
		}
	}
}

//check if an EMPTY channel is closed
func IsClosed(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
	}
	return false
}
