package main

import "fmt"

func main() {
	var n, cnt int
	var s string
	fmt.Scan(&n, &s)
	for i := 0; i < n-2; i++ {
		if string(s[i:i+3]) == "ABC" {
			cnt++
		}
	}
	fmt.Println(cnt)
}
