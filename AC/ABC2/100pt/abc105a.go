package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	rem := n % k
	switch rem {
	case 0:
		fmt.Println(0)
	default:
		fmt.Println(1)
	}
}
