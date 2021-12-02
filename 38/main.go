package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(len(permutation("sabces")))
}

var ss []string
var res []string

func permutation(s string) []string {
	ss = strings.Split(s, "")
	res = []string{}
	r(0)
	return res
}

func r(k int) {
	if k == len(ss)-1 {
		res = append(res, strings.Join(ss, ""))
		return
	}
	m := make(map[string]bool)
	for i := k; i < len(ss); i++ {
		ss[i], ss[k] = ss[k], ss[i]
		if _, ok := m[ss[k]]; ok {
			ss[i], ss[k] = ss[k], ss[i]
			continue
		}
		m[ss[k]] = true
		r(k + 1)
		ss[i], ss[k] = ss[k], ss[i]
	}
}
