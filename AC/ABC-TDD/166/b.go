package main

import "fmt"

func main() {
	var n, k, d, tmp, cnt int
	fmt.Scan(&n, &k)
	sunuke := make([]bool, n)
	for i := 0; i < k; i++ {
		fmt.Scan(&d)
		for j := 0; j < d; j++ {
			fmt.Scan(&tmp)
			sunuke[tmp-1] = true
		}
	}
	for _, v := range sunuke {
		if !v {
			cnt++
		}
	}
	fmt.Println(cnt)
}
