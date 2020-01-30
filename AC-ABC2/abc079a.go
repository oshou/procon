package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	if strings.Repeat(string(s[0]), 3) == s[0:3] || strings.Repeat(string(s[1]), 3) == s[1:4] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
