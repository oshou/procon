package main

import "fmt"

func main() {
	var a, b string
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	if a == b {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}
}
