package main

import "fmt"

func main() {
	fmt.Println(cuttingRope(51))
}

// func cuttingRope(n int) int {
// 	a, b, p := n/3, n%3, 1000000007
// 	res := 1
// 	for i := 1; i <= a; i++ {
// 		res = (3 * res) % p
// 	}
// 	if b == 0 {
// 		return res
// 	}
// 	return res * b
// }

func cuttingRope(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = maxer(dp[i], (maxer(dp[j], j)*maxer(dp[i-j], i-j))%1000000007)
		}
	}
	return dp[n]
}

func maxer(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func remMulti(a, b, p int) int {
	return ((a % p) * (b % p)) % p
}
