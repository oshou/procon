package solution

import (
	"math"
)

func axb_c(n int) int {
	var cnt int
	divider := int(math.Sqrt(float64(n)))
	for i := 1; i <= divider; i++ {
		for j := i; j <= n; j++ {
			if i*j < n {
				if i == j {
					cnt += 1
				} else {
					cnt += 2
				}
			}
		}
	}
	return cnt
}
