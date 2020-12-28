package main

import "fmt"

func main() {
	var h, n, tmp, sum int
	fmt.Scan(&h, &n)
	for i := 0; i < n; i++ {
		fmt.Scan(&tmp)
		sum += tmp
	}
	if sum >= h {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
