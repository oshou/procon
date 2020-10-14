package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	for i := 0; i <= x; i++ {
		if 4*i+2*(x-i) == y {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
