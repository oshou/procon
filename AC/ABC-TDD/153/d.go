package main

import "fmt"

func main() {
	var h int
	fmt.Scan(&h)
	fmt.Println(sep(h))
}

func sep(h int) int {
	if h == 1 {
		return 1
	}
	return 1 + 2*sep(h/2)
}
