package main

import "fmt"

func main() {
	var (
		v   int
		arr []int
	)

	for {
		fmt.Scan(&v)
		if v == 0 {
			break
		}
		arr = append(arr, v)
	}

	for i, v := range arr {
		fmt.Printf("Case %d: %d\n", i+1, v)
	}
}
