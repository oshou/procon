package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if 1 <= a && a <= 9 && 1 <= b && b <= 9 {
		fmt.Println(a * b)
	} else {
		fmt.Println(-1)
	}
}
