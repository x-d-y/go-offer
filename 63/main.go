package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	min := math.MaxInt32
	max := 0

	for _, v := range prices {
		if profile := v - min; profile > max {
			max = profile
		}
		if v < min {
			min = v
		}
	}
	return max
}

func r(nums []int, no, min, max int) int {
	if no == len(nums) {
		return max
	}
	if profile := nums[no] - min; profile > max {
		max = profile
	}
	if nums[no] < min {
		min = nums[no]
	}
	return r(nums, no+1, min, max)
}
