package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	//if a == b+c || b == a+c || c == a+b {
	if a+b+c/2 == math.Max(a, b, c) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
