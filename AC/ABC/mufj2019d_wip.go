package main

import "fmt"

func main() {
	var n, cnt int
	var s string
	m := make(map[string]bool)
	fmt.Scan(&n)
	fmt.Scan(&s)
	for i := 0; i < n; i++ {
		c := string(s[i])
		if !m[c] {
			m[c] = true
		}
	}

	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			for k := 1; k <= n; k++ {
				substr := string(s[i-1]) + string(s[j-1]) + string(s[k-1])
				if !m[substr] {
					m[substr] = true
					cnt++
				}
			}
		}
	}
	fmt.Println(cnt)
}
