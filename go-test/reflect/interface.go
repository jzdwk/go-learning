package main

import (
	"fmt"
	"reflect"
)

type IA interface {
	methodA()
}

type IB interface {
	methodB()
}

type impl struct {
}

func NewA() IA {
	return impl{}
}

func (impl) methodA() {
	fmt.Println("impl A")
}

func (impl) methodA2() {
	fmt.Println("impl A2")
}

func (impl) methodB() {
	fmt.Println("impl B")
}

func main() {
	//im := impl{}
	//var im IA
	im := NewA()
	typeImpl := reflect.TypeOf(im)
	v := reflect.ValueOf(im)
	imA := v.Interface().(IA)
	imA.methodA()
	fmt.Println(typeImpl.Kind(), typeImpl.Name())
}
