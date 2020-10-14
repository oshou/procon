package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	if N == 1 {
		fmt.Println("Hello World")
	} else {
		var A, B int
		fmt.Scan(&A, &B)
		fmt.Println(A + B)
	}
}
