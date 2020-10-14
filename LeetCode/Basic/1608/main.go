package main

func specialArray(nums []int) int {
	var cnt, max int
	max = -1
	for i := 0; i <= len(nums); i++ {
		cnt = 0
		for _, num := range nums {
			if num >= i {
				cnt++
			}
		}
		if i == cnt && max < i {
			max = i
		}
	}
	return max
}
