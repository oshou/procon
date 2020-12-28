package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k, sum int
	fmt.Scan(&n, &k)
	h := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&h[i])
	}
	if n <= k {
		fmt.Println(0)
		return
	}
	sort.Ints(h)
	for i := 0; i < n-k; i++ {
		sum += h[i]
	}
	fmt.Println(sum)
}
