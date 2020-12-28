package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, k int
	var s, ans string
	fmt.Scan(&n, &k)
	fmt.Scan(&s)
	for i := 0; i < n; i++ {
		if i == k-1 {
			ans += strings.ToLower(string(s[i]))
		} else {
			ans += string(s[i])
		}
	}
	fmt.Println(ans)
}
