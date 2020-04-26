package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for i := n - 1; i >= 0; i-- {
		if i != n-1 {
			fmt.Print(" ")
		}
		fmt.Print(a[i])
	}
	fmt.Println()
}
