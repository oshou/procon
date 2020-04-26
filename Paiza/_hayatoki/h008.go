package main

import "fmt"

func main() {
	var n float64
	fmt.Scan(&n)
	fmt.Printf("%.1f %.1f\n", n-18, n-18.5)
}
