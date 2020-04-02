package main

import "fmt"

func main() {
	var n, a, b, blue, rem int
	fmt.Scan(&n, &a, &b)
	blue = n / (a + b) * a
	rem = n % (a + b)
	if rem >= a {
		blue += a
	} else {
		blue += rem
	}
	fmt.Println(blue)
}
