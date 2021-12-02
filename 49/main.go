package main

import "fmt"

func main() {
	fmt.Println(nthUglyNumber(10))
}

func nthUglyNumber(n int) int {
	dp := make([]int, n)

	dp[0] = 1

	a, b, c := 0, 0, 0

	for i := 1; i < n; i++ {
		ax := dp[a] * 2
		bx := dp[b] * 3
		cx := dp[c] * 5
		min := miner(miner(ax, bx), cx)
		dp[i] = min
		if min == ax {
			a++
		}
		if min == bx {
			b++
		}
		if min == cx {
			c++
		}
	}
	return dp[n-1]
}

func miner(a, b int) int {
	if a < b {
		return a
	}
	return b
}
