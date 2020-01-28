package main

import "fmt"

func main() {
	var n, min int
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &n)
		if i == 0 {
			min = n
		} else {
			if min > n {
				min = n
			}
		}
	}
	fmt.Println(min)
}
