package main

func majorityElement(nums []int) int {
	var ans int
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	border := len(nums) / 2
	for k, cnt := range m {
		if cnt > border {
			ans = k
		}
	}
	return ans
}
