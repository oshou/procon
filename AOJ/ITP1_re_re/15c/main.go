package main

import "fmt"

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
			for j := 0; j < w; j++ {
				switch {
				case (i%2 == 0 && j%2 == 0) || (i%2 != 0 && j%2 != 0):
					fmt.Print("#")
				default:
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
