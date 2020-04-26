package main

import "fmt"

func main() {
	var n, s, p int
	fmt.Scan(&n, &s, &p)
	if s*p > n {
		fmt.Println("none")
		return
	}
	for i := s*(p-1) + 1; i <= s*p; i++ {
		if i != s*(p-1)+1 {
			fmt.Print(" ")
		}
		fmt.Print(i)
	}
	fmt.Println()
}
