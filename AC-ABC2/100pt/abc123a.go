package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c, d, e, k int
	fmt.Scan(&a, &b, &c, &d, &e, &k)
	arr := []int{a, b, c, d, e}
	sort.Ints(arr)
	if arr[len(arr)-1]-arr[0] <= k {
		fmt.Println("Yay!")
	} else {
		fmt.Println(":(")
	}
}
