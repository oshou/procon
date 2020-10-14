package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, asum, bsum int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			asum += arr[i]
		} else {
			bsum += arr[i]
		}
	}
	fmt.Println(asum - bsum)
}
