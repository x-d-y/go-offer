package main

func main() {

}

func add(a int, b int) int {
	for b != 0 {
		n := a ^ b
		c := a & b << 1
		a, b = n, c
	}
	return a
}
