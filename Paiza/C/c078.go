package main

import "fmt"

func main() {
	var n, c1, c2, p, num, profit int
	fmt.Scan(&n, &c1, &c2)
	for i := 1; i <= n; i++ {
		fmt.Scan(&p)
		if i == n {
			profit += num * p
			continue
		}
		switch {
		case p <= c1:
			num++
			profit -= p
		case c2 <= p:
			profit += num * p
			num = 0
		}
	}
	fmt.Println(profit)
}
