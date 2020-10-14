package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	uniques := map[int]bool{}
	uniques[a] = true
	uniques[b] = true
	uniques[c] = true
	fmt.Println(len(uniques))
}
