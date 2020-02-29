package main

import "fmt"

func Amin(nums ...int) int {
	min := -1
	for _, num := range nums {
		if min == -1 {
			min = num
		}
		if min > num {
			min = num
		}
	}
	return min
}

func main() {
	var a, b, c, x, y int
	var high, low, highCost int
	fmt.Scan(&a, &b, &c, &x, &y)
	if x >= y {
		high, low = x, y
		highCost = a
	} else {
		high, low = y, x
		highCost = b
	}
	p1 := a*x + b*y
	p2 := low*2*c + (high-low)*highCost
	p3 := high * 2 * c
	fmt.Println(Amin(p1, p2, p3))
}
