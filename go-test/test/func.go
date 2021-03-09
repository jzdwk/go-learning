/*
@Time : 21-1-4
@Author : jzd
@Project: go-learning
*/
package test

import (
	"context"
	"fmt"
	"strings"
)

type RouteConfig struct {
	Uuid  string
	Trans *TransConfig
}

type TransConfig struct {
	Uuid   string
	Config map[string]interface{}
}

func main2() {
	/*p1 := "this is p1"
	p2 := 2021

	list := []funcValue{
		createFuncValue1(p1, p2),
		createFuncValue2(p1, p2),
	}

	for _, o := range list {
		o(context.Background(), p1, p2)
	}*/
	/*path := "/abc/get/info/*d+//"
	subPath := strings.Split(path,"/")
	if subPath[1] == "abc"{
		finalPath := ""
		for _, v := range subPath {
			if v != "abc"{
				finalPath = finalPath+v+"/"
			}
		}
		fmt.Println(finalPath[0:len(finalPath)])
	}*/
	/*maps := make(map[string]string)
	set(&maps)
	fmt.Println(maps)*/
	removePathPrefix("/cmcc2/delete/\\d+/info", "/cmcc2")
}

func removePathPrefix(path string, prefix string) string {
	if prefix == "" {
		return path
	}
	finalPath := ""
	subPath := strings.Split(path, "/")
	//old path has baseUrl
	if subPath[1] == prefix {
		for _, v := range subPath {
			if v != prefix {
				finalPath = finalPath + v + "/"
			}
		}
		finalPath = finalPath[0 : len(finalPath)-1]
	} else {
		finalPath = path
	}
	return finalPath
}

func set(maps *map[string]string) {
	(*maps)["a"] = "a"
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
