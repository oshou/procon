package main

import "fmt"

func main() {
	var n, v, p, p0, p1, p2, p3, sum int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&v, &p)
		switch v {
		case 0:
			p0 += p
		case 1:
			p1 += p
		case 2:
			p2 += p
		case 3:
			p3 += p
		}
		sum = (p0 / 100 * 5) + (p1 / 100 * 3) + (p2 / 100 * 2) + (p3 / 100 * 1)
	}
	fmt.Println(sum)
}
