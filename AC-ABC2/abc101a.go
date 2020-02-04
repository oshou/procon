package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	ans := 0
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "+":
			ans++
		case "-":
			ans--
		}
	}
	fmt.Println(ans)
}
