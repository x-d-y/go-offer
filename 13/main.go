package main

import "fmt"

func main() {
	fmt.Println(movingCount(11, 11, 1))
}

var maxM, maxN, count int

func movingCount(m int, n int, k int) int {
	maxM, maxN, count = m, n, 0
	b := buildBool(m, n)
	r(0, 0, k, b)
	return count
}

func r(m, n, k int, b [][]bool) {
	if m < maxM && n < maxN && m >= 0 && n >= 0 && isTarget(m, n, k) && b[m][n] == false {
		b[m][n] = true
		count++
		r(m, n-1, k, b)
		r(m, n+1, k, b)
		r(m-1, n, k, b)
		r(m+1, n, k, b)
	}
}

func buildBool(m, n int) [][]bool {
	res := make([][]bool, m)
	for i := 0; i < m; i++ {
		res[i] = make([]bool, n)
	}
	return res
}

func isTarget(a, b, k int) bool {
	return sum(a)+sum(b) <= k
}

func sum(a int) int {
	if a == 100 {
		return 1
	}

	return a%10 + (a/10)%10

}
