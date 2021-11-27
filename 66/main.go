package main

import "fmt"

func main() {
	fmt.Println(constructArr([]int{1, 2, 3, 4, 5}))
}

func constructArr(a []int) []int {
	b := make([]int, len(a))
	last := 1
	b[0] = last
	a = append(a, 1)

	for i := 1; i < len(a)-1; i++ {
		last *= a[i-1]
		b[i] = last
	}
	last = 1
	for i := len(a) - 2; i > 0; i-- {
		a[i] = a[i] * a[i+1]
	}

	for i := 0; i < len(a)-1; i++ {
		b[i] = b[i] * a[i+1]

	}
	return b
}
