package main

import "fmt"

func main() {
	var n, mid int
	var s string
	fmt.Scan(&n, &s)
	if n%2 == 1 {
		fmt.Println("No")
		return
	}
	mid = int(n / 2)
	for i := 0; i < mid; i++ {
		if s[i] != s[i+mid] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
