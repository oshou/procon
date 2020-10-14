package main

import "fmt"

func main() {
	x := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Scan(&x[i])
	}
	for i := 0; i < 5; i++ {
		if x[i] != i+1 {
			fmt.Println(i + 1)
			return
		}
	}
}
