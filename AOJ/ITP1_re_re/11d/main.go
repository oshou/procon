package main

import "fmt"

func main() {
	var v int
	fmt.Scan(&v)
	h := v / 3600
	m := (v % 3600) / 60
	s := (v % 3600) % 60
	fmt.Printf("%d:%d:%d\n", h, m, s)
}
