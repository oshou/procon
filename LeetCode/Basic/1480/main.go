package main

func runningSum(nums []int) []int {
	var sum int
	var ans []int
	for _, num := range nums {
		sum += num
		ans = append(ans, sum)
	}
	return ans
}
