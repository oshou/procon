package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	var n int
	fmt.Scan(&s)
	fmt.Scan(&n)
	cnt := strings.Count(s, "R")
	if cnt >= n || cnt == len(s) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
