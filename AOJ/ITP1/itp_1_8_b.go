package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x string
	var sum int
	for {
		fmt.Scan(&x)
		if x == "0" {
			return
		}
		sum = 0
		for i := 0; i < len(x); i++ {
			r, _ := strconv.Atoi(string(x[i]))
			sum += r
		}
		fmt.Println(sum)
	}
}
