package main

import "fmt"

func main() {
	hasharr := make(map[int]string)
	hasharr[1] = "12322"
	hasharr[2] = "andy"
	hasharr[3] = "男"

	fmt.Println(hasharr)

	for k, v := range hasharr {
		fmt.Println(k, v)
	}
}
