package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		h, w int
	)
	for {
		fmt.Scan(&h, &w)
		if h == 0 && w == 0 {
			break
		}
		for i := 0; i < h; i++ {
			switch i {
			case 0, h - 1:
				fmt.Println(strings.Repeat("#", w))
			default:
				fmt.Printf("#%s#\n", strings.Repeat(".", w-2))
			}
		}
		fmt.Println()
	}
}
