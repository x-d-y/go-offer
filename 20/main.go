package main

import (
	"fmt"
)

func main() {
	fmt.Println(isNumber("e3"))
}

func isNumber(s string) bool {
	transfer := make([]map[string]int, 9)
	status0 := make(map[string]int)
	status0[" "] = 0
	status0["s"] = 1
	status0["d"] = 2
	status0["."] = 4
	transfer[0] = status0
	status1 := make(map[string]int)
	status1["."] = 4
	status1["d"] = 2
	transfer[1] = status1
	status2 := make(map[string]int)
	status2["e"] = 5
	status2["d"] = 2
	status2["."] = 3
	status2[" "] = 8
	transfer[2] = status2
	status3 := make(map[string]int)
	status3["e"] = 5
	status3["d"] = 3
	status3[" "] = 8
	transfer[3] = status3

	status4 := make(map[string]int)
	status4["d"] = 3
	transfer[4] = status4

	status5 := make(map[string]int)
	status5["d"] = 7
	status5["s"] = 6
	transfer[5] = status5

	status6 := make(map[string]int)
	status6["d"] = 7
	transfer[6] = status6

	status7 := make(map[string]int)
	status7["d"] = 7
	status7[" "] = 8
	transfer[7] = status7

	status8 := make(map[string]int)
	status8[" "] = 8
	transfer[8] = status8

	var t string
	var status int
	for _, v := range s {
		if v == rune(' ') {
			t = " "
			if v, ok := transfer[status][t]; ok {
				status = v
			} else {
				return false
			}
			continue
		}
		if v == rune('+') || v == rune('-') {
			t = "s"
			if v, ok := transfer[status][t]; ok {
				status = v
			} else {
				return false
			}
			continue
		}
		if v == rune('.') {
			t = "."
			if v, ok := transfer[status][t]; ok {
				status = v
			} else {
				return false
			}
			continue
		}

		if v >= rune('0') && v <= rune('9') {
			t = "d"
			if v, ok := transfer[status][t]; ok {
				status = v
			} else {
				return false
			}
			continue
		}
		if v == 'e' || v == 'E' {
			t = "e"
			if v, ok := transfer[status][t]; ok {
				status = v
			} else {
				return false
			}
			continue
		}
		return false
	}
	if status == 2 || status == 7 || status == 3 || status == 8 {
		return true
	}
	return false
}
