package main

import "fmt"

func main() {
	var x, y string
	fmt.Scan(&x, &y)
	hex := map[string]int{"A": 10, "B": 11, "C": 12, "D": 13, "E": 14, "F": 15}
	switch {
	case hex[x] > hex[y]:
		fmt.Println(">")
	case hex[x] < hex[y]:
		fmt.Println("<")
	case hex[x] == hex[y]:
		fmt.Println("=")
	}
}
