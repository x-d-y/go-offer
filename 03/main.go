package main

import "fmt"

func main() {
	fmt.Println(findRepeatNumber([]int{3, 4, 2, 0, 0, 1}))
}

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums);  {
		if i == nums[i] {
			i++
			continue
		}
		if nums[nums[i]] != nums[i] {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			i = nums[i]
		} else {
			return nums[i]
		}
	
	}
	return -1
}
