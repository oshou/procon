package main

func singleNumber(nums []int) int {
	var m = make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}
