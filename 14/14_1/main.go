package main

import "fmt"

func main() {
	fmt.Println(cuttingRope(9))
}

var res []int
var c int

func cuttingRope(n int) int {
	res = make([]int, 0)
	if n <= 1 {
		return n
	}
	res = append(res, 0, 1)
	for i := 2; i <= n; i++ {
		c = 0
		for j := 1; j <= i/2; j++ {
			k := i - j
			c = maxer(maxer(k, res[k])*maxer(j, res[j]), c)
		}
		res = append(res, c)
	}
	return res[n]
}

func maxer(a, b int) int {
	if a > b {
		return a
	}
	return b
}
