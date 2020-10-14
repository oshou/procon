package main

import "fmt"

func main() {
	var n, k, sum int
	m := make(map[int]int)
	fmt.Scan(&n, &k)
	for bits := 0; bits < (1 << uint64(n+1)); bits++ {
		fmt.Println(bits)
		sum = 0
		for i := 0; i < n+1; i++ {
			if (bits>>uint64(i))&1 == 1 {
				sum += i + 100
			}
		}
		m[sum]++
	}
	fmt.Println(m)
	fmt.Println(len(m))
}
