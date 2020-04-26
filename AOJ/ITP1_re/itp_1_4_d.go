package main

import "fmt"

func main() {
	var n, a, min, max, sum int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a)
		if i == 1 {
			min = a
			max = a
			sum = a
			continue
		}
		if a < min {
			min = a
		}
		if max < a {
			max = a
		}
		sum += a
	}
	fmt.Println(min, max, sum)
}
