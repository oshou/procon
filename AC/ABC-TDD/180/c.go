package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, tmp int
	var ans []int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if i*i > n {
			break
		}
		if n%i == 0 {
			ans = append(ans, i)
			tmp = n / i
			if tmp != i {
				ans = append(ans, tmp)
			}
		}
	}
	sort.Ints(ans)
	for _, nums := range ans {
		fmt.Println(nums)
	}
}
