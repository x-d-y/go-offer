package main

import "fmt"

func main() {
	fmt.Println(singleNumbers([]int{1, 3, 3, 8, 6, 6}))
}

func singleNumbers(nums []int) []int {
	s := 0
	for _, v := range nums {
		s = s ^ v
	}
	fmt.Println(s)
	m := 1
	for s&m==0{
		m <<= 1
	}
	s1, s2 := 0, 0
	for _, v := range nums {
		if m&v == 0 {
			s1 = s1 ^ v
		} else {
			s2 = s2 ^ v
		}
	}
	return []int{s1,s2}
}
