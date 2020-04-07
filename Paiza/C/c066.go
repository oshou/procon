package main

import "fmt"

func main() {
	var m, n, x, d, w, cnt int
	fmt.Scan(&m, &n, &x)
	d = x
	for i := 0; i < m; i++ {
		fmt.Scan(&w)
		if n <= 0 {
			continue
		}
		for {
			if d-w > 0 {
				d -= w
				cnt++
				break
			} else {
				n--
				if n <= 0 {
					break
				}
				d = x
			}
		}
	}
	fmt.Println(cnt)
}
