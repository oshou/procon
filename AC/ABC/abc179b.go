package main

import "fmt"

func main() {
	var n, cnt int
	var isTriple bool
	fmt.Scan(&n)
	d1 := make([]int, n)
	d2 := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&d1[i], &d2[i])
		if d1[i] == d2[i] {
			cnt++
		} else {
			cnt = 0
		}
		if cnt >= 3 {
			isTriple = true
		}
	}

	if isTriple {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
