package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if x >= y {
		fmt.Println("Thank you")
	} else {
		fmt.Println(y - x)
	}
}
