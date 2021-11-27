package main

import "fmt"

func main() {
	fmt.Println(getLeastNumbers([]int{3, 2, 1}, 3))
}

// 使用堆或者快排

var nums []int

func getLeastNumbers(arr []int, k int) []int {
	nums = make([]int, 0)
	r(arr, k)
	return nums
}
func r(arr []int, k int) {
	if k == 0 {
		return
	}
	var flagl, flagr bool
	var i, j int
	for i, j = 0, len(arr)-2; i <= j; {
		if arr[i] <= arr[len(arr)-1] {
			i++
		} else {
			flagl = true
		}
		if arr[j] > arr[len(arr)-1] {
			j--
		} else {
			flagr = true
		}
		if flagl && flagr {
			arr[i], arr[j] = arr[j], arr[i]
			flagl, flagr = false, false
		}
	}
	list1 := append(append(make([]int, 0), arr[:i]...), arr[len(arr)-1])
	list2 := arr[i : len(arr)-1]
	if len(list1) == k {
		nums = append(nums, list1...)
		return
	} else if len(list1) < k {
		nums = append(nums, list1...)
		r(list2, k-len(list1))
	} else {
		r(list1[:len(list1)-1], k)
	}

}
