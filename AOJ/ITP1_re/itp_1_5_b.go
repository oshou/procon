package main

import "fmt"

func main() {
	var h, w int
	for {
		fmt.Scan(&h, &w)
		if h == 0 && w == 0 {
			break
		}
		for i := 1; i <= h; i++ {
			for j := 1; j <= w; j++ {
				if i == 1 || i == h {
					fmt.Print("#")
					continue
				}
				if j == 1 || j == w {
					fmt.Print("#")
					continue
				}
				fmt.Print(".")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
