package main

import "fmt"

func main() {
	var h, w int
	s := make([]string)
	for i := 0; i < h; i++ {
		fmt.Scan(&s[i])
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w-1; j++ {
			if string(s[i][j])=="."
		}
	}
}
