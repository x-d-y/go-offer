package main

import "fmt"

func main() {
	fmt.Println(search([]int{5, 5, 9, 8, 7, 6}, 2))
}

func search(nums []int, target int) int {
	var l, r int
	for i, j, k := 0, len(nums)-1, len(nums)/2; i <= j; {
		if nums[k] <= target {
			i = k + 1
		} else {
			j = k - 1
		}
		k = (i + j) / 2
		r = i
	}
	if r == 0 || nums[r-1] != target {
		return 0
	}

	for i, j, k := 0, r-1, (r-1)/2; i <= j; {
		if nums[k] >= target {
			j = k - 1
		} else {
			i = k + 1
		}
		k = (i + j) / 2
		l = j
	}
	return r - l - 1
}
