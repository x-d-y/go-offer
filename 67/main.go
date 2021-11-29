package main

import (
	"fmt"
)

func main() {
	fmt.Println(strToInt("-+1"))
}

func strToInt(str string) int {
	var sign bool
	var digital bool
	var res int
	var active = true
	for _, v := range str {
		if !sign && v == rune(' ') {
			continue
		}
		if !sign && (v == rune('-') || v == rune('+') ){
			sign = true
			if v == rune('-') {
				active = false
			}
			continue
		}
		if v <= rune('9') && v >= rune('0') {
			sign = true
			digital = true
			if res > 214748364 {
				res = 2147483648
				break
			}
			if res == 214748364 {
				if v > rune('7') {
					res = 2147483648
					break
				}
			}
			res = res*10 + int(v-'0')

			continue
		}
		if digital && (v <= rune('0') || v >= rune('9') ){
			if active {
				return res
			} else {
				return -res
			}
		}
		return res
	}
	if active {
		if res == 2147483648 {
			return 2147483647
		}
		return res
	} else {
		return -res
	}
}
