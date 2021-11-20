package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abbac"))
	//lengthOfLongestSubstring("abcd")
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	l, r := 0, 0
	m := make(map[rune]int)
	for i, v := range s {

		if vv,ok:= m[v];ok {
			l = maxNum(vv +1,l)
		}
		m[v] = i
		max = maxNum(max, r-l+1)
		r++
}
	return max
}

func maxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
