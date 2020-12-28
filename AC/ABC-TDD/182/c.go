package main

import (
	"fmt"
	"strconv"
)

func GetDiv3Cnt(s string) int {
	memo := make(map[string]int)
	if len(s) == 0 {
		return -1
	}

	i, _ := strconv.Atoi(s)
	if i%3 == 0 {
		return 0
	}

	mincnt := len(s) + 1
	for i := 0; i < len(s); i++ {
		sd := s[:1] + s[i+1:]
		div := 0
		if memo[sd] != 0 {
			div = memo[sd]
		} else {
			div = GetDiv3Cnt(sd)
		}
		if div != -1 {
			mincnt = min(mincnt, 1+div)
		}
	}

	if mincnt == len(s)+1 {
		return -1
	} else {
		return mincnt
	}
}

func min(a, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(GetDiv3Cnt(s))
}
