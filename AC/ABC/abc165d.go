package main

import (
	"fmt"
)

func main() {
	var a, b, n int
	var max, num int
	fmt.Scan(&a, &b, &n)
	for i := n; i >= 1; i-- {
		num = int(a*i/b) - a*int(i/b)
		fmt.Printf("i:%v, num:%v\n", i, num)
		if i == n {
			max = num
			continue
		}
		if max < num {
			max = num
		}
	}
	fmt.Println(max)
}
