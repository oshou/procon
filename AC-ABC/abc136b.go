package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, cnt int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if len(strconv.Itoa(i))%2 != 0 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
