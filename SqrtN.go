package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(math.Sqrt(2))
}

func sqrt(x float64) float64 {
	if x == 0 {
		return 0
	}
	if x < 0 {
		return math.NaN()
	}

	// 设置精度要求
	precision := 0.000001

	// 定义搜索范围
	left, right := 0.0, x

	// 迭代计算
	for right-left > precision {
		mid := (left + right) / 2
		if mid*mid > x {
			right = mid
		} else {
			left = mid
		}
	}

	return left
}
