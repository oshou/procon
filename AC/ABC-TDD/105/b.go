package main

import "fmt"

func main() {
	var n, ans, cnt int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			continue
		}
		cnt = 0
		for j := 1; j <= n; j++ {
			if i%j == 0 {
				cnt++
			}
		}
		if cnt == 8 {
			ans++
		}
	}
	fmt.Println(ans)
}
