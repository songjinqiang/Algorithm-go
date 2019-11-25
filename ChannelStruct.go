package main

import (
	"fmt"
)

type user struct {
	name string
	sex  string
	age  int
}

var info chan user = make(chan user)

func Routine() {
	var userinfo user
	userinfo.name = "andy"
	userinfo.sex = "男"
	userinfo.age = 18

	info <- userinfo
	info <- userinfo
	close(info)
}

func main() {
	go Routine()
	fmt.Println("Main")

	for receive := range info {
		fmt.Println(receive)
	}
}
