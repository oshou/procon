package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	var num, sum int
	for {
		fmt.Scan(&s)
		if s == "0" {
			break
		}
		sum = 0
		for _, c := range s {
			num, _ = strconv.Atoi(string(c))
			sum += num
		}
		fmt.Println(sum)
	}
}
