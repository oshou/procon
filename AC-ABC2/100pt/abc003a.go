package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println((10000 * n * (n + 1) / 2) / n)
}
