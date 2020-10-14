package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var n, tmp int
	fmt.Scan(&n)
	max := int(math.Sqrt(float64(n)))
	fmin := len(strconv.Itoa(n))
	for i := 1; i <= max; i++ {
		if n%i == 0 {
			tmp = f(i, n/i)
			if tmp < fmin {
				fmin = tmp
			}
		}
	}
	fmt.Println(fmin)
}

func f(a, b int) int {
	alen := len(strconv.Itoa(a))
	blen := len(strconv.Itoa(b))
	if alen > blen {
		return alen
	} else {
		return blen
	}
}
