package main

import "fmt"

func main() {
	var ans bool
	var sum int
	nums := make([]int, 4)
	total := 0
	for i := 0; i < 4; i++ {
		fmt.Scan(&nums[i])
		total += nums[i]
	}
	n := 4
	for i := 0; i < (1 << uint(n)); i++ {
		sum = 0
		for j := 0; j < n; j++ {
			if i>>uint(j)&1 == 1 {
				sum += nums[j]
			}
		}
		if sum*2 == total {
			ans = true
		}
	}
	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
