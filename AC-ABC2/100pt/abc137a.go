package main

import "fmt"

func max(nums ...int) int {
	var max int
	for i := range nums {
		if i == 0 {
			max = nums[i]
			continue
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(max(a+b, a-b, a*b))
}
