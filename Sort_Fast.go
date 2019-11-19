package main

import "fmt"

func main() {
	arr := []int{2, 1, 5, 4, 7, 10, 8, 3, 6}
	res := FastSort(arr)
	fmt.Println(res)
}

func FastSort(arr []int) []int {
	len := len(arr)
	if len <= 1 {
		return arr
	}
	key := arr[0]
	var left []int
	var right []int
	for i := 1; i < len; i++ {
		if arr[i] < key {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	left = FastSort(left)
	right = FastSort(right)
	target := []int{key}
	res := append(append(left, target...), right...)
	return res
}
