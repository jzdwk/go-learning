/*
@Time : 20-10-12
@Author : jzd
@Project: go-learning
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "/tet/{info}/{id}/"
	elm := strings.Split(path, "/")
	for _, value := range elm {
		//fmt.Println(value)
		if value != "" {
			/*if value[0:]=="{"&&value[len(value)-1:]=="}"{
				fmt.Println(value)
			}*/
			if strings.HasPrefix(value, "{") && strings.HasSuffix(value, "}") {
				args := value[1 : len(value)-1]
				if args == "id" {
					path = strings.ReplaceAll(path, value, "\\\\d+")
				}
				if args == "info" {
					path = strings.ReplaceAll(path, value, ".*")
				}
			}
		}
	}
	fmt.Println(path)
}
