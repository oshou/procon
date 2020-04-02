package main

import "fmt"

func main() {
	var y, m, d, a, b, sum int
	fmt.Scan(&y, &m, &d)
	fmt.Scan(&a, &b)
	// year
	yy := (5 - (y % 4)) % 4
	sum += ((yy - 1) % 4) * 181
	// int
	if m%2 == 0 {
		sum += ((13-(m+1))/2)*28 + 13
		sum += 15 - d
	} else {
		sum += ((13 - (m + 1)) / 2) * 28
		sum += 13 - d
	}
	// out
	if a%2 == 0 {
		sum += ((a-1)/2)*28 + 13
		sum += b
	} else {
		sum += ((a - 1) / 2) * 28
		sum += b
	}
	fmt.Println(sum)
}
