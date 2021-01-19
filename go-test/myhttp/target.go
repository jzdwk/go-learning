/*
@Time : 21-1-11
@Author : jzd
@Project: go-learning
*/
package main

import (
	"flag"
	"fmt"
	"net/http"
)

var port = flag.Int("port", 80, "port")

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, fmt.Sprintf("Hello world from %d", *port))
}

func main() {
	flag.Parse()
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
