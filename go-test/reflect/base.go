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

type Test struct {
	FiledA   int64    `required:"true"`
	FiledB   int      `required:"false"`
	FiledC   string   `required:"true"`
	FiledArr []string `required:"true"`
	FiledD   bool     `required:"true"`
	FiledE   TestC
	TestA
	*TestB
}

type TestA struct {
	FiledIn string `required:"false"`
}

type TestB struct {
	FiledIn string `required:"false"`
}

type TestC struct {
	FiledIn string `required:"true"`
}

type myInt int

func main() {
	// base test
	/*test := test{a: "hello", b: "world"}
	typeT := reflect.TypeOf(test)
	typePt := reflect.TypeOf(&typeT)
	fmt.Println("name: ", typeT.Name(),"kind: ",typeT.Kind())
	fmt.Println("name: ", typePt.Name(),"kind: ",typePt.Kind())*/

	// self define type
	/*var myInt myInt = 1
	typeMy := reflect.TypeOf(myInt)
	fmt.Println("name: ", typeMy.Name(),"kind: ", typeMy.Kind())*/

	//empty

	//test := Test{FiledA:1,FiledB:2,FiledD:false}
	/*path1 := "test"
	//path2 := "test"
	yes, msg := IsValidUrlPath(path1)
	fmt.Println(yes,msg)*/
	//isEmpty(test,test2)
	/*now := time.Now()
	fmt.Println(now.String())
	timeLocation, _ := time.LoadLocation("Asia/Shanghai")
	now2 := now.In(timeLocation)
	fmt.Println(now2.String())*/

	a, b := 1, 2

	switch {
	case a == b:
		fmt.Println("a")
		break
	case a != b:
		fmt.Println("b")
		break

	}

}

/*
func isEmpty(stt interface{})bool{
	//v := reflect.ValueOf(stt)
	t := reflect.TypeOf(stt)
	for i := 0;i<t.NumField() ;i++  {
		fieldType := t.Field(i)
		if fieldType.Tag.Get("required") == "true"{
			//fieldValue := t.Field(i)
			fmt.Println("filed type name: ",fieldType.Type.Kind())
			switch fieldType.Type.Kind() {
			case reflect.Map: //map
				break
			case reflect.Slice: //slice
			case reflect.Ptr: //point
			case reflect.Struct: //struct
			case reflect.Int: //int
			case reflect.Int64:

			default:


			}
		}
	}
	return false
}
*/

func isEmpty(field ...string) {
	for _, value := range field {
		t := reflect.TypeOf(value)
		v := reflect.ValueOf(value)
		name := v.Field(0).Type().Name()
		fmt.Println(t.String(), name)
	}
}

func IsValidUrlPath(path string) (bool, string) {
	if path == "" {
		return false, "path is empty"
	}
	if path[:1] != "/" {
		return false, "path should start with '/'"
	}
	return true, ""
}
