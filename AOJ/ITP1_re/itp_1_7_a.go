package main

import "fmt"

func main() {
	var m, f, r, total int
	for {
		total = 0
		fmt.Scan(&m, &f, &r)
		if m == -1 && f == -1 && r == -1 {
			break
		}
		total = m + f
		switch {
		case m == -1 || f == -1:
			fmt.Println("F")
		case total >= 80:
			fmt.Println("A")
		case total >= 65:
			fmt.Println("B")
		case total >= 50 || (total >= 30 && r >= 50):
			fmt.Println("C")
		case total >= 30:
			fmt.Println("D")
		default:
			fmt.Println("F")
		}
	}
}
