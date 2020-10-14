package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	cnt := strings.Count(s, "o")
	fmt.Println(700 + 100*cnt)
}
