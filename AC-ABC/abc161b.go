package main

import "fmt"

func main() {
	var n, m, sum, cnt int
	var line float64
	fmt.Scan(&n, &m)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		sum += a[i]
	}
	line = float64(sum) / (4 * float64(m))
	for i := 0; i < n; i++ {
		if float64(a[i]) >= line {
			cnt++
		}
	}
	if cnt >= m {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
