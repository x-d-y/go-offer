package main

import "fmt"

func main() {
	fmt.Println(maxSlidingWindow([]int{1,3,1,2,0,5}, 3))
}

var dequeue []int
var queue []int
var res []int

func maxSlidingWindow(nums []int, k int) []int {
	if k > len(nums) || k == 0 || len(nums) == 0 {
		return []int{}
	}
	dequeue = make([]int, 0)
	queue = make([]int, 0)
	res = make([]int, 0)
	for i := 0; i < k; i++ {
		queue = append(queue, nums[i])
		if len(dequeue) == 0 {
			dequeue = append(dequeue, nums[i])
		} else {
			dequeue = insert(dequeue, nums[i])
		}
	}

	for i := k; i < len(nums); i++ {
		res = append(res, dequeue[0])
		pop := queue[0]
		queue = queue[1:]
		if pop == dequeue[0] {
			dequeue = dequeue[1:]
		}
		queue = append(queue, nums[i])
		if len(dequeue) == 0 {
			dequeue = append(dequeue, nums[i])
		} else {
			dequeue = insert(dequeue, nums[i])
		}
	}
	return append(res, dequeue[0])
}
func insert(dequeue []int, k int) []int {
	for i, v := range dequeue {
		if v < k {
			dequeue = append(make([]int, 0), dequeue[:i]...)
			break
		}
	}
	dequeue = append(dequeue, k)
	return dequeue
}
