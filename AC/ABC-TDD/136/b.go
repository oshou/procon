package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, cnt int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		s := strconv.Itoa(i)
		if len(s)%2 != 0 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
