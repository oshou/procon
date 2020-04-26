package main

import (
	"fmt"
)

func main() {
	var h, m, s int
	fmt.Scan(&s)
	h = s / 3600
	s %= 3600
	m = s / 60
	s %= 60
	fmt.Printf("%d:%d:%d\n", h, m, s)
}
