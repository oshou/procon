package main

import "fmt"

func main() {
	var s string
	var n int
	fmt.Scan(&s)
	fmt.Scan(&n)
	fmt.Printf("%s%s\n", s[:n-1], s[n:])
}
