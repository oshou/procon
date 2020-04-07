package main

import "fmt"

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	var t, x, y, a, b int
	fmt.Scan(&t, &x, &y)
	var xn, yn int = x, y
	var xmax, xmin int = x, x
	for i := 0; i < t; i++ {
		fmt.Scan(&a, &b)
		xn += a
		yn += b
		if xn > xmax {
			xmax = xn
		}
		if xn < xmin {
			xmin = xn
		}
		if yn <= 0 {
			break
		}
	}
	fmt.Println(Max(xmax, xmin))
}
