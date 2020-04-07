package main

import "fmt"

func main() {
	var n int
	var a string
	fmt.Scan(&n)
	m := make(map[string]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		m[a]++
	}
	switch {
	case m["paper"] > 0 && m["rock"] > 0 && m["scissors"] > 0:
		fmt.Println("draw")
	case m["paper"] == n || m["rock"] == n || m["scissors"] == n:
		fmt.Println("draw")
	case m["paper"] == 0:
		fmt.Println("rock")
	case m["rock"] == 0:
		fmt.Println("scissors")
	case m["scissors"] == 0:
		fmt.Println("paper")
	}
}
