package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
	}
	fmt.Scan(&a[0][0], &a[0][1])
	fmt.Scan(&a[1][0], &a[1][1])
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if j >= 1 {
				a[i][j] = a[i][j-1] + (a[i][1] - a[i][0])
			}
			if i >= 1 {
				a[i][j] = a[i-1][j] + (a[1][j] - a[0][j])
			}
			if j != 0 {
				fmt.Print(" ")
			}
			fmt.Print(a[i][j])
		}
		fmt.Println()
	}
}
