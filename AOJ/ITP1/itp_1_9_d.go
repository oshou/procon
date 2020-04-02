package main

import (
	"fmt"
)

func main() {
	var s, order, p string
	var q, a, b int
	fmt.Scan(&s)
	fmt.Scan(&q)
	for i := 0; i < q; i++ {
		fmt.Scan(&order, &a, &b)
		switch order {
		case "replace":
			fmt.Scan(&p)
			s = s[:a] + p + s[b+1:]
		case "reverse":
			tmp := s[:a]
			for i := b; i >= a; i-- {
				tmp += string(s[i])
			}
			tmp += s[b+1:]
			s = tmp
		case "print":
			fmt.Println(string(s[a : b+1]))
		}
	}
}
