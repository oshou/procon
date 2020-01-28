package main

import "fmt"

func main() {
	var n, ans int
	fmt.Scan(&n)
	for i := n; i < n+10; i++ {
		if i >= 10 {
			ans = i % 10
		} else {
			ans = i
		}
		fmt.Println(ans)
	}
}
