package main

import "fmt"

func main() {
	fmt.Println(myPow(2.000, 10))
}


func myPow(x float64, n int) float64 {
	if n > 0 {
		return r(x, n)
	} else if n == 0 {
		return 1
	} else {
		return r(1/x, -n)
	}
}

func r(x float64, n int) float64 {

	if n == 0 {
		return 1
	}
    s := r(x, n>>1)
	if n&1 == 1 {
		return x * s*s
	}
	return s*s
}
