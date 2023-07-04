package main

import "fmt"

func setZeroes(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i, v := range matrix {
		for j, c := range v {
			if c == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i, v := range matrix {
		for j := range v {
			if row[i] || col[j] {
				v[j] = 0
			}
		}
	}
}

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	fmt.Println(matrix)
}
