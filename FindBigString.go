package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "pwwkew"
	res := FindBigSting(str)
	fmt.Println(res)
}

func FindBigSting(str string) int {
	len := len(str)
	fmt.Println()
	var tmp string
	var check int
	var res int
	for i := 0; i < len; i++ {
		target := string(str[i])
		check = strings.Index(tmp, target)
		if check == -1 {
			tmp += target
		} else {
			tmp += target
			tmp = string(tmp[check+1:])
		}
		tmpLen := strings.Count(tmp, "") - 1
		if tmpLen > res {
			res = tmpLen
		}
	}
	return res
}
