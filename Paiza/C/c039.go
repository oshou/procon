package main

import "fmt"

func main() {
	var s string
	var tmp, sum int
	fmt.Scan(&s)
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "<":
			tmp += 10
		case "/":
			tmp += 1
		case "+":
			sum += tmp
			tmp = 0
		}
	}
	sum += tmp
	fmt.Println(sum)
}
