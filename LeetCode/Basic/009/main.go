package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	strx := strconv.Itoa(x)
	var reversed string
	for i := 0; i < len(strx); i++ {
		reversed = string(strx[i]) + reversed
	}

	if strx == reversed {
		return true
	}
	return false
}
