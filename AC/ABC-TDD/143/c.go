package main

import "fmt"

func main() {
	var n, cnt int
	var s string
	fmt.Scan(&n, &s)
	cnt = 1
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			continue
		} else {
			cnt += 1
		}
	}
	fmt.Println(cnt)
}
