package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	first := strings.Repeat("#", 10)
	second := strings.Repeat(".", 10)
	for i := 0; i < n; i++ {
		fmt.Println(first)
		fmt.Println(second)
	}
}
