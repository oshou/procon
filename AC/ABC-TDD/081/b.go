package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= 9; i++ {
		if n%i != 0 {
			continue
		}
		tmp := n / i
		if 1 <= tmp && tmp <= 9 {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
