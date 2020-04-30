package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	m := make(map[string]bool)
	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		m[s] = true
	}
	fmt.Println(len(m))
}
