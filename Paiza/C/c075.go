package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	var f, c, p int = 0, n, 0
	for i := 0; i < m; i++ {
		fmt.Scan(&f)
		if f <= p {
			p -= f
		} else {
			c -= f
			p += f / 10
		}
		fmt.Println(c, p)
	}
}
