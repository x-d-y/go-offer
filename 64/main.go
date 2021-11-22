package main

import "fmt"

func main() {
	fmt.Println(sumNums(3))
}

var ans int

func sumNums(n int) int {
	r(n)
	return ans
}

func r(n int) bool {
	ans += n
	return n > 0 && r(n-1)
}
