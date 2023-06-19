package main

import "fmt"

func main() {
	//第一遍
	numbers := []int{1, 1, 2, 2, 3, 3, 4, 5, 5, 6}
	xor := 0
	for _, number := range numbers {
		xor ^= number
	}

	//相同数异或后为0 xor为4和6异或后的值不为0 取最低位的1 xor有可能是多个1
	mask := xor & (-xor)

	//第二遍
	group1, group2 := 0, 0
	for _, number := range numbers {
		if number&mask != 0 {
			group1 ^= number
		} else {
			group2 ^= number
		}
	}

	fmt.Println(group1, group2)
}
