package main

import "fmt"

func main() {
	var n, mmin int
	fmt.Scan(&n)
	mmin = n - 1
	for i := 1; i <= n; i++ {
		if i*i > n {
			break
		}
		if n%i == 0 {
			mmin = min(mmin, (i-1)+(n/i-1))
		}
	}
	fmt.Println(mmin)
}

func min(x, y int) int {
	if x >= y {
		return y
	} else {
		return x
	}
}
