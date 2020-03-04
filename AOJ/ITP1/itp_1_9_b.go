package main

import (
	"fmt"
)

func main() {
	var s string
	var m, h int
	for {
		fmt.Scan(&s)
		if s == "-" {
			break
		}
		fmt.Scan(&m)
		for i := 0; i < m; i++ {
			fmt.Scan(&h)
			s = s[h:] + s[:h]
		}
		fmt.Println(s)
	}
}
