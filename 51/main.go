package main

import "fmt"

func main() {
	fmt.Println(reversePairs([]int{1,3,2,4}))
}

func reversePairs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	total = 0
	r(nums)
	return total
}

var total int

func r(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	left := r(nums[:len(nums)/2])

	right := r(nums[len(nums)/2:])

	res := []int{}

	var i, j, k int
	for ; i < len(right); i++ {
		for j = k; j < len(left); j++ {
			if left[j] > right[i] {
				total += len(left) - j
				break
			} else {
				k++
				res = append(res, left[j])
			}
		}
		if k == len(left)-1 {
			res = append(res, right[i:]...)
			break
		} else {
			res = append(res, right[i])
		}
	}
	res = append(res, left[k:]...)
	return res

}
