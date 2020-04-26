package main

import "fmt"

func main() {
	var x, i int
	fmt.Scan(&x)
	i = 1
	for x != 0 {
		fmt.Printf("Case %d: %d\n", i, x)
		i++
		fmt.Scan(&x)
	}
}
