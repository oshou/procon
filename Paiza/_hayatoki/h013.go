package main

import "fmt"

func main() {
	var n float64
	fmt.Scan(&n)
	fmt.Printf("%d%%\n", int(n/140*100))
}
