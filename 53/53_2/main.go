package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{0, 1, 3, 4, 5, 6, 7, 8, 9}))
}

func missingNumber(nums []int) int {
	if nums[len(nums)-1] == len(nums)-1 {
		return len(nums)
	}
	var i, j = 0, len(nums) - 1
	k := (i + j) / 2
	for i < j {
		if k == nums[k] {
			i = k + 1
			k = (i + j) / 2
		} else {
			j = k
			k = (i + j) / 2
		}
	}
	return i
}
