package main

import "fmt"

func main() {
	var a, b int
	var op string
	for {
		fmt.Scan(&a, &op, &b)
		if op == "?" {
			break
		}
		switch op {
		case "+":
			fmt.Println(a + b)
		case "-":
			fmt.Println(a - b)
		case "*":
			fmt.Println(a * b)
		case "/":
			fmt.Println(a / b)
		}
	}
}
