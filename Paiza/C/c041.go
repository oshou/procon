package main

import "fmt"

type Medal struct {
	gold   int
	silver int
	bronze int
}

func main() {
	var n int
	fmt.Scan(&n)
	m := make([]Medal, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m[i].gold, &m[i].silver, &m[i].bronze)
	}
	for i := 0; i < n; i++ {
		for j := n - 1; j >= i; j-- {
			if j == n-1 {
				continue
			}
			switch {
			case m[j+1].gold > m[j].gold:
				m[j+1], m[j] = m[j], m[j+1]
			case m[j+1].gold == m[j].gold && m[j+1].silver > m[j].silver:
				m[j+1], m[j] = m[j], m[j+1]
			case m[j+1].gold == m[j].gold && m[j+1].silver == m[j].silver && m[j+1].bronze > m[j].bronze:
				m[j+1], m[j] = m[j], m[j+1]
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(m[i].gold, m[i].silver, m[i].bronze)
	}
}
