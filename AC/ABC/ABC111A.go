package main

import "fmt"

func main() {
	var N, n string
	N = fmt.Scan(&N)
	for idx, value := range N {
		print(value)
	}
}
