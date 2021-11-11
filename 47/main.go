package main

import "fmt"

func main() {
	fmt.Println(maxValue([][]int{{1,3,1},{1,5,1},{4,2,1}}))
	fmt.Println(maxValue([][]int{{}}))

}

func maxValue(grid [][]int) int {
	if len(grid[0]) == 0{
		return 0
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			cNum := 0
			if i-1 >= 0 {
				cNum = grid[i-1][j]
			}
			if j-1 >= 0 {
				if cNum < grid[i][j-1] {
					cNum = grid[i][j-1]
				}
			}
			grid[i][j] += cNum
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
