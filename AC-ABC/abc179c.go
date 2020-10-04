package main

import (
	"fmt"
)

func main() {
	var n, cnt int
	fmt.Scan(&n)
	for i := 1; i < n; i++ {
		cnt += divider(i)
	}
	fmt.Println(cnt)
}

func divider(num int) int {
	var cnt int
	for j := 1; j*j <= num; j++ {
		if num%j == 0 {
			if j*j == num {
				cnt += 1
			} else {
				cnt += 2
			}
		}
	}
	return cnt
}
