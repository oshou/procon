package main

import "fmt"

func min(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var n, b, f, r, v int
	fmt.Scan(&n)
	t := make([][][]int, 4)
	for i := 0; i < 4; i++ {
		t[i] = make([][]int, 3)
		for j := 0; j < 3; j++ {
			t[i][j] = make([]int, 10)
		}
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&b, &f, &r, &v)
		t[b-1][f-1][r-1] += v
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 10; k++ {
				fmt.Printf(" %d", min(0, t[i][j][k]))
			}
			fmt.Println()
		}
		if i < 3 {
			fmt.Println("####################")
		}
	}
}
