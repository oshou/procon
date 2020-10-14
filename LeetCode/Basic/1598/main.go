package main

func minOperations(logs []string) int {
	var cnt int
	for _, log := range logs {
		switch log {
		case "../":
			if cnt != 0 {
				cnt -= 1
			}
		case "./":
			cnt += 0
		default:
			cnt += 1
		}
	}
	return cnt
}
