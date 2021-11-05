package main

import "fmt"

func main() {
	fmt.Println(numWays(3))
	fmt.Println(numWays(4))
}

var num []int

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n <= 2 {
		return n
	}
	num = []int{}
	num = append(num, 0, 1, 2)
	for i := 3; i <= n; i++ {
		num = append(num, (num[i-1]+num[i-2])%1000000007)
	}
	return num[n]
}
