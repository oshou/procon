package main

import "fmt"

func main() {
	var (
		a, b, v int
		op      string
	)

	for {
		fmt.Scan(&a, &op, &b)
		if op == "?" {
			break
		}
		switch op {
		case "+":
			v = a + b
		case "-":
			v = a - b
		case "*":
			v = a * b
		case "/":
			v = a / b
		}
		fmt.Println(v)
	}
}
