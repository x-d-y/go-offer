package main

import "fmt"

func main() {
	fmt.Println(findNumberIn2DArray([][]int{{1}}, 1))
	//{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	row := len(matrix)
	col := 0
	if row > 0 {
		col = len(matrix[0])
	}
	if row == 0 && col == 0 {
		return false
	}
	c := 0
	for i := row - 1; 0 <= i; i-- {
		for j := c; j < col; j++ {
			if target > matrix[i][j] {
				c = j
				continue
			}
			if target < matrix[i][j] {
				break
			}
			return true
		}
	}
	return false
}
