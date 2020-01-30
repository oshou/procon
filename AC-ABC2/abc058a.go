package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	if b-a == c-b {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
