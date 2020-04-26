package main

import "fmt"

func main() {
	var a, b, c, cnt int
	fmt.Scan(&a, &b, &c)
	for i := a; i <= b; i++ {
		if c%i == 0 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
