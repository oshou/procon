package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	var n, b, f, r, v int
	fmt.Scan(&n)
	var room [4][3][10]int
	for i := 0; i < n; i++ {
		fmt.Scan(&b, &f, &r, &v)
		room[b-1][f-1][r-1] = max(room[b-1][f-1][r-1]+v, 0)
	}

	for i := 0; i < 4; i++ {
		if i != 0 {
			fmt.Println("####################")
		}
		for j := 0; j < 3; j++ {
			for k := 0; k < 10; k++ {
				fmt.Printf(" %v", room[i][j][k])
			}
			fmt.Println()
		}
	}
}
