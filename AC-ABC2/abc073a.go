package main

import (
	"fmt"
	"strings"
)

func main() {
	var n string
	fmt.Scan(&n)
	if strings.Contains(n, "9") {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
