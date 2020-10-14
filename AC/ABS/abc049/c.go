package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	if canFetch(s) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func canFetch(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) >= 5 && (strings.HasPrefix(s, "dream") || strings.HasPrefix(s, "erase")) {
		if canFetch(s[5:]) {
			return true
		}
	}
	if len(s) >= 6 && (strings.HasPrefix(s, "eraser")) {
		if canFetch(s[6:]) {
			return true
		}
	}
	if len(s) >= 7 && (strings.HasPrefix(s, "dreamer")) {
		if canFetch(s[7:]) {
			return true
		}
	}
	return false
}
