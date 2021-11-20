package main

func main() {

}

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}
