package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	times := b / a
	if times > c {
		fmt.Println(c)
	} else {
		fmt.Println(times)
	}
}
