package main

import (
	"fmt"
	"math"
)

func main() {
	var m, n, cnt int
	var l float64
	fmt.Scan(&m, &n)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			ifloat := float64(i)
			jfloat := float64(j)
			l = math.Sqrt(ifloat*ifloat + jfloat*jfloat)
			if math.Ceil(l) == l {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
