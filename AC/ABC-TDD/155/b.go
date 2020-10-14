package main

import (
	"fmt"
)

func main() {
	var n, a int
	fmt.Scan(&n)
	var ans bool = true
	for i := 1; i <= n; i++ {
		fmt.Scan(&a)
		if a%2 == 0 && (a%3 != 0 && a%5 != 0) {
			ans = false
		}
	}

	if ans {
		fmt.Println("APPROVED")
	} else {
		fmt.Println("DENIED")
	}
}
