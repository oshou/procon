package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	if s == strings.Repeat("A", len(s)) || s == strings.Repeat("B", len(s)) {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
