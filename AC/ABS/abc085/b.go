package main

import "fmt"

func main() {
	var n, d int
	fmt.Scan(&n)
	m := make(map[int]bool)
	for i := 0; i < n; i++ {
		fmt.Scan(&d)
		m[d] = true
	}
	fmt.Println(len(m))
}
