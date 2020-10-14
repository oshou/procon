package main

import "fmt"

func main() {
	var a, b, c, x, y, tmp, mmin int
	mmin = -1
	fmt.Scan(&a, &b, &c, &x, &y)
	max := maxint(x, y)

	for i := 0; i <= max; i++ {
		tmp = 2*c*i + a*(maxint(0, x-i)) + b*(maxint(0, y-i))
		if mmin == -1 {
			mmin = tmp
			continue
		}
		if tmp < mmin {
			mmin = tmp
		}
	}
	fmt.Println(mmin)
}

func maxint(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
