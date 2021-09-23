package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		hs, ws []int
		h, w   int
	)

	for {
		fmt.Scan(&h, &w)
		if h == 0 && w == 0 {
			break
		}
		hs = append(hs, h)
		ws = append(ws, w)

		for i := 0; i < h; i++ {
			fmt.Println(strings.Repeat("#", w))
		}
		fmt.Println()
	}
}
