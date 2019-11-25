package main

import (
	"fmt"
)

var info chan string

func Routine() {
	fmt.Println("routine1")
	fmt.Println("routine2")
	fmt.Println("routine3")

	info <- "info1"
	info <- "info2"
	info <- "info3"
	close(info)
}

func main() {
	info = make(chan string)
	go Routine()
	fmt.Println("Main")
	for receive := range info {
		fmt.Println(receive)
	}
}
