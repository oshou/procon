package main

import "fmt"

func main() {
	var (
		n, b, f, r, v int
		address       [4][3][10]int
	)

	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&b, &f, &r, &v)
		address[b-1][f-1][r-1] = v
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 10; k++ {
				fmt.Printf(" %d", address[i][j][k])
			}
			fmt.Println()
		}
		fmt.Println("####################")
	}
}
