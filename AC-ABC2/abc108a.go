package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)
	even := 0
	odd := 0
	even = k / 2
	if k%2 == 0 {
		odd = k / 2
	} else {
		odd = k/2 + 1
	}
	fmt.Println(even * odd)
}
