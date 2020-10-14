package main

import "fmt"

func main() {
	var n, s string
	fmt.Scan(&n)
	s = ""
	for _, c := range n {
		switch string(c) {
		case "9":
			s += "1"
		case "1":
			s += "9"
		}
	}
	fmt.Println(s)
}
