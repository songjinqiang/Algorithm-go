package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	target := 2
	res := BinaryFind(arr, target)
	fmt.Println(res)
}

func BinaryFind(arr []int, target int) int {
	len := len(arr)
	low := 0
	high := len
	var mid int

	for low < high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			high = mid
		}
		if arr[mid] < target {
			low = mid
		}
	}
	return mid
}
