package main

import "fmt"

func main() {
	var n, m, idx, a, b int
	r := make(map[int]int, n)
	fmt.Scan(&n, &m)
	for i := 1; i <= m; i++ {
		fmt.Scan(&a, &b)
		if r[a] == 0 && r[b] == 0 {
			idx++
			r[a], r[b] = idx, idx
		}
		if r[a] != 0 && r[b] != 0 {
			min := min(r[a], r[b])

			idx--
		}
		if r[a] == 0 {
			r[a] = b
		}
		if r[b] == 0 {
			r[b] = a
		}
	}
	fmt.Println((cnt - 1) + (n - len(r)))
}
