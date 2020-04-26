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
				remi := i % 2
				remj := j % 2
				switch {
				case (remi == 1 && remj == 1) || (remi == 0 && remj == 0):
					fmt.Print("#")
				case (remi == 1 && remj == 0) || (remi == 0 && remj == 1):
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
