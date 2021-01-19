/*
@Time : 20-11-30
@Author : jzd
@Project: go-learning
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	rst := FindCgroupMountpoint("memory")
	fmt.Println(rst)
}

func FindCgroupMountpoint(subsystem string) string {
	f, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		return ""
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text()
		fields := strings.Split(txt, " ")
		for _, opt := range strings.Split(fields[len(fields)-1], ",") {
			if opt == subsystem {
				return fields[4]
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return ""
	}

	return ""
}
