package main

import "fmt"

func main() {
	var k, n int
	fmt.Scan(&k, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	mid := k - a[n-1] + a[0]
	for i := 0; i < n; i++ {
		if i == 0 {
			continue
		}
		tmp := a[i] - a[i-1]
		if mid < tmp {
			mid = tmp
		}
	}
	fmt.Println(k - mid)
}
