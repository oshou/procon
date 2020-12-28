package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	var s1i, s2i int
	var isYYMM, isMMYY bool
	fmt.Scan(&s)
	s1i, _ = strconv.Atoi(s[0:2])
	s2i, _ = strconv.Atoi(s[2:4])
	if 1 <= s1i && s1i <= 12 {
		isMMYY = true
	}
	if 1 <= s2i && s2i <= 12 {
		isYYMM = true
	}
	switch {
	case isYYMM && isMMYY:
		fmt.Println("AMBIGUOUS")
	case isYYMM:
		fmt.Println("YYMM")
	case isMMYY:
		fmt.Println("MMYY")
	default:
		fmt.Println("NA")
	}
}
