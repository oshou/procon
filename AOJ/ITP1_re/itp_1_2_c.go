package main

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a > b {
		swap(&a, &b)
	}
	if b > c {
		swap(&b, &c)
	}
	if a > b {
		swap(&a, &b)
	}
	fmt.Println(a, b, c)
}
