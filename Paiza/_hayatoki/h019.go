package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	switch {
	case 20 <= n && n <= 15000:
		fmt.Println("yes")
	case 15000 < n && n <= 20000:
		fmt.Println("not sure")
	default:
		fmt.Println("no")
	}
}
