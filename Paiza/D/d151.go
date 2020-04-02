package main

import "fmt"

func main() {
	var s string
	var n int
	fmt.Scan(&s, &n)
	switch s {
	case "chocolate":
		fmt.Println(n * 2)
	case "cake":
		fmt.Println(n * 5)
	}
}
