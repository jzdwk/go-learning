package main

import "fmt"

type Top struct {
	msg string
}

func (t *Top) doSth() string {
	return t.msg
}
func (t *Top) doSth2() string {
	return t.msg + "2"
}

type Second struct {
	Top
}

func (s *Second) doSth() string {
	return s.msg + "second"
}

func main() {
	t := Second{Top: Top{msg: "test"}}
	fmt.Println(t.doSth())
	fmt.Println(t.doSth2())
	//无法向上转型，golang的继承通过组合实现，因此严格意义上说，不具备继承
	test(t)
}

func test(t Top) {
	return
}
