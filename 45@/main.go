package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(minNumber([]int{1}))
}

func minNumber(nums []int) string {
	numstrings := make([]string, 0)
	for _, v := range nums {
		numstrings = append(numstrings, strconv.Itoa(v))
	}
	res := qs(numstrings)
	return strings.Join(res, "")
}

func qs(nums []string) []string {
	if len(nums) == 0 {
		return []string{}
	}
	if len(nums) == 1 {
		return []string{nums[0]}
	}
	flagi, flagj := false, false
	i, j := 0, len(nums)-2
	for i <= j {
		if smaller(nums[i], nums[len(nums)-1]) {
			i++
		} else {
			flagi = true
		}

		if !smaller(nums[j], nums[len(nums)-1]) {
			j--
		} else {
			flagj = true
		}

		if flagi && flagj {
			nums[i], nums[j] = nums[j], nums[i]
			flagi, flagj = false, false
		}
	}
	l := qs(nums[:max(i, j)])
	r := qs(nums[max(i, j) : len(nums)-1])

	return append(append(l, nums[len(nums)-1]), r...)
}

func smaller(a, b string) bool {
	if a+b > b+a {
		return false
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
