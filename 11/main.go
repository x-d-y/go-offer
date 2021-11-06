package main

import (
	"fmt"
)

func main() {
	fmt.Println(minArray([]int{}))
}

func minArray(numbers []int) int {
	l, r := 0, 0
	a := len(numbers) / 2
	l, r = a, a
	// if len(numbers)%2 == 1 {
	// 	l, r = a, a
	// } else {
	// 	l = a - 1
	// 	r = a
	// 	if numbers[l] > numbers[r] {
	// 		return numbers[r]
	// 	}
	// }
	li, ri := l, r
	for l >= 0 || r <= len(numbers)-1 {
		if l >= 0 {
			if numbers[l] <= numbers[li] {
				li = l
			} else {
				return numbers[li]
			}
			l--
		}
		if r < len(numbers) {
			if numbers[r] >= numbers[ri] {
				ri = r
			} else {
				return numbers[r]
			}
			r++
		}
	}
	return numbers[0]
}
