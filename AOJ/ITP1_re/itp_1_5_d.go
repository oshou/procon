package main

import "fmt"

func main() {
	var n int
	var exists bool
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		exists = false
		tmp := i
		if tmp%3 == 0 {
			fmt.Printf(" %d", i)
			continue
		}
		for tmp != 0 {
			if tmp%10 == 3 {
				exists = true
			}
			tmp /= 10
		}
		if exists {
			fmt.Printf(" %d", i)
			continue
		}
	}
	fmt.Println()
}
