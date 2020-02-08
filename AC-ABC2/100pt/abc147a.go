package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a+b+c <= 21 {
		fmt.Println("win")
	} else {
		fmt.Println("bust")
	}
}
