package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, cnt int
	var s string
	fmt.Scan(&n, &s)
	for i := 0; i < 1000; i++ {
		c := make([]int, 3)
		c[0] = i / 100
		c[1] = i/10 - c[0]*10
		c[2] = i % 10
		idx := 0
		for j := 0; j < n; j++ {
			if string(s[j]) == strconv.Itoa(c[idx]) {
				idx++
			}
			if idx == 2 {
				cnt++
				break
			}
		}
	}
	fmt.Println(cnt)
}
