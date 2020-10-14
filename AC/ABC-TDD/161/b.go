package main

import "fmt"

func main() {
	var n, m, cnt, sum int
	fmt.Scan(&n, &m)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		sum += a[i]
	}
	border := float64(sum) / float64(4*m)
	for i := 0; i < n; i++ {
		if float64(a[i]) >= border {
			cnt++
		}
	}

	if cnt >= m {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
