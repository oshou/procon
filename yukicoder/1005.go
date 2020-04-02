package main

import (
	"fmt"
)

func count(str, substr string) int {
	cnt := 0
	for i := 0; i < len(str)-len(substr)+1; i++ {
		if str[i:i+len(substr)] == substr {
			cnt++
		}
	}
	return cnt
}

func main() {
	var cnt int
	var s, t string
	fmt.Scan(&s, &t)
	cnt = count(s, t)
	if cnt >= 1 && len(t) <= 1 {
		fmt.Println(-1)
	} else {
		fmt.Println(cnt)
	}
}
