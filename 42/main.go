package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSubArray([]int{-2}))
}

func maxSubArray(nums []int) int {

	max, pre := -101, -101
	for _, v := range nums {
		pre = maxer(pre+v, v)
		max = maxer(max, pre)
	}

	return max
}

func maxer(a, b int) int {
	if a > b {
		return a
	}
	return b
}
