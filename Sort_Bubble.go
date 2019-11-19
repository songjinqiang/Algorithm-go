package main

import "fmt"

func BubbleSort(arr []int) []int {
	len := len(arr)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if arr[i] > arr[j] {
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}
	return arr
}

func main() {
	arr := []int{2, 4, 1, 5, 6}
	res := BubbleSort(arr)
	fmt.Println(res)
}
