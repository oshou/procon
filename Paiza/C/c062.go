package main

import "fmt"

func main() {
	var t, m, w int
	var n string
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		if n == "melon" && w == 0 {
			m++
			w = 10
		}
		if w > 0 {
			w--
		}
	}
	fmt.Println(m)
}
