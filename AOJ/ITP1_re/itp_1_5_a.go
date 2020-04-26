package main

import (
	"fmt"
	"strings"
)

func main() {
	var h, w int
	for {
		fmt.Scan(&h, &w)
		if h == 0 && w == 0 {
			break
		}
		for i := 1; i <= h; i++ {
			fmt.Println(strings.Repeat("#", w))
		}
		fmt.Println()
	}
}
