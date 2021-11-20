package main

import (

	"fmt"
	"strconv"
)

func main(){
	fmt.Println(translateNum(12258))
}

func translateNum(num int) int {
	s:=strconv.Itoa(num)
	last ,lastl:= 1,0
	for i,_ := range s{
		var c int
		c = last
		if i >= 1&& isNum(s[i-1],s[i]){
		c = last + lastl
		}
		lastl = last
		last = c
	}
	return last
}

func isNum(a,b byte)bool{
	if a == byte("1"[0]) || a == byte("2"[0]) && b <= "5"[0] {
		return true
	}
	return false
}




