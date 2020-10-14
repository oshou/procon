package main

import "fmt"

const MAX = 200001

func main() {
	var mmin, n int
	fmt.Scan(&n)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&p[i])
	}
	nums := make([]bool, MAX)
	for i := 0; i < n; i++ {
		nums[p[i]] = true
		if nums[mmin] == true {
			mmin = min(nums, mmin+1, p[i])
		}
		fmt.Println(mmin)
	}
}

func min(nums []bool, start, hit int) int {
	for i := start; i < len(nums); i++ {
		if nums[i] == false {
			return i
		}
	}
	return 0
}
