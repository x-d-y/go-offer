package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("hhygygsdfjsgdghsq"))
}

func firstUniqChar(s string) byte {
	hash := [26]int{}

	for _, v := range s {
		hash[v-'a']++
	}

	for _, v := range s {
		if hash[v-'a'] == 1 {
			return byte(v)
		}
	}
	return ' '
}
