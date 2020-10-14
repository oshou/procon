package main

import "fmt"

func main() {
	var n, cnt int
	var s string
	m := make(map[string]int)
	fmt.Scan(&n)
	fmt.Scan(&s)
	for i := 0; i < n; i++ {
		c := string(s[i])
		m[c]++
	}
	fmt.Println(m)

	for _, v1 := range m {
		for _, v2 := range m {
			for _, v3 := range m {
				v1--
				v2--
				v3--
				if v1 >= 0 && v2 >= 0 && v3 >= 0 {
					cnt++
				}
			}
		}
	}
	fmt.Println(cnt)
}
