package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var n, a, b int
	fmt.Scan(&n)
	for i := 1; i <= 37; i++ {
		a = int(math.Pow(3, float64(i)))
		for j := 1; j <= 26; j++ {
			b = int(math.Pow(5, float64(j)))
			if n == a+b {
				fmt.Println(i, j)
				os.Exit(0)
			}
		}
	}
	fmt.Println(-1)
}
