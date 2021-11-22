package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(7))
}

func hammingWeight(num uint32) int {
	if num == 0 {
		return 0
	}
	var i int
	for num&(num-1) != 0 {
		num = num & (num - 1)
		i++
	}
	return i + 1
}
