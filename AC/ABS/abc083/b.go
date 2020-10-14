package main

import "fmt"

func main() {
	var n, a, b, tmp, dsum, sum int
	fmt.Scan(&n, &a, &b)
	for i := 1; i <= n; i++ {
		tmp = i
		dsum = 0
		for tmp != 0 {
			dsum += tmp % 10
			tmp /= 10
		}
		if a <= dsum && dsum <= b {
			sum += i
		}
	}
	fmt.Println(sum)
}
