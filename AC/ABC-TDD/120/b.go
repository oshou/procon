package main

import "fmt"

func main() {
	var a, b, k, cnt, max int
	fmt.Scan(&a, &b, &k)
	max = maxint(a, b)
	for i := max; i >= 1; i-- {
		if a%i == 0 && b%i == 0 {
			cnt++
		}
		if cnt == k {
			fmt.Println(i)
			return
		}
	}
}

func maxint(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
