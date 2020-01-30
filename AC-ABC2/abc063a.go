package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	sum := a + b
	if sum < 10 {
		fmt.Println(sum)
	} else {
		fmt.Println("error")
	}
}
