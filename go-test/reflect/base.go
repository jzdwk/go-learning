package main

import (
	"fmt"
	"reflect"
)

type test struct {
	a  string
	b  string
	in `json:"test"`
}

type in struct {
	ina string
	inb string
}

type myInt int

func main2() {
	test := test{a: "hello", b: "world"}
	typeT := reflect.TypeOf(test)
	//typePt := reflect.TypeOf(&typeT)
	//var myInt myInt = 1
	//typeMy := reflect.TypeOf(myInt)
	//fmt.Println(typeT.Name(), typeT.Kind())
	//fmt.Println(typePt.Name(),typePt.Kind())
	//fmt.Println(typeMy.Name(),typeMy.Kind())
	field := typeT.Field(2)
	fmt.Println(typeT.NumField())
	fmt.Println(field.Anonymous, field.Tag.Get("json"))
}
