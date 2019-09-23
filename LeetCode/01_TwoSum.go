package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ret := []int{i, j}
				return ret
			}
		}
	}
	return nil
}

func main() {
	var nums []int = []int{1, 4, 8, 3, 2, 9, 15}
	var target int = 13
	fmt.Println(twoSum(nums, target))
}
