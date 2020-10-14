package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	memo := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		memo[a[i]]++
	}
	for k1, v1 := range memo {
		if len(memo) == 0 {
			fmt.Println(0)
			return
		}
		if v1 > 1 {
			delete(memo, k1)
		}
		for k2, v2 := range memo {
			if v2 > 1 {
				delete(memo, k2)
			}
			if k1 < k2 {
				if k2%k1 == 0 {
					delete(memo, k2)
				}
			}
		}
	}
	fmt.Println(len(memo))
}
