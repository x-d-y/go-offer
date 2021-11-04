package main

import (
	"fmt"
)

func main() {
	fmt.Println(fib(45))
}

var m map[int]int

func fib(n int) int {
	res := 0
	defer func() {
		if m == nil {
			m = make(map[int]int)
		}
		m[n] = res
	}()
	if n <= 1 {
		res = n
		return res
	}
	a := fib(n - 1)
	b := m[n-2]
	if n == 45 {
		fmt.Print("debug")
	}
	res = (a + b)%1000000007
	return res
}
