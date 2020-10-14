package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	if strings.Contains(s, "7") {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
