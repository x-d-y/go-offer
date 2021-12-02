package main

import "fmt"

func main() {
	fmt.Println(dicesProbability(2))
}

func dicesProbability(n int) []float64 {
	rate := 1.0 / 6.0
	dp := make([][]float64, n)
	//dp[0] = []float64{}
	dp[0] = []float64{rate, rate, rate, rate, rate, rate}
	for i := 1; i < n; i++ {
		dp[i] = make([]float64, 5*(i+1)+1)
		for j := 0; j < len(dp[i-1]); j++ {
			jr := dp[i-1][j] * rate
			for k := j; k < j+6 && k < len(dp[i]); k++ {
				dp[i][k] += jr
			}
		}
	}
	return dp[len(dp)-1]
}
