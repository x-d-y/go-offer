package main

import "fmt"

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{3, 4, 1, 5, 2}))
}

func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	var j int
	for _, v := range popped {
		if len(stack) >= 1 && v == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			continue
		}
		for ; j < len(pushed) && pushed[j] != v; j++ {
			stack = append(stack, pushed[j])
		}
		j++
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
