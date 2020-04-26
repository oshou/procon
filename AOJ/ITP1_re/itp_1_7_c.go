package main

import "fmt"

func main() {
	var r, c int
	fmt.Scan(&r, &c)
	m := make([][]int, r+1)
	for i := 0; i < r+1; i++ {
		m[i] = make([]int, c+1)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			fmt.Scan(&m[i][j])
			m[i][c] += m[i][j]
			m[r][j] += m[i][j]
			m[r][c] += m[i][j]
		}
	}
	for i := 0; i < r+1; i++ {
		for j := 0; j < c+1; j++ {
			if j != 0 {
				fmt.Print(" ")
			}
			fmt.Print(m[i][j])
		}
		fmt.Println()
	}
}
