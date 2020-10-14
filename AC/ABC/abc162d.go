package main

import (
	"fmt"
	"strings"
)

func rgb(a, b string) string {
	switch {
	case a != "R" && b != "R":
		return "R"
	case a != "G" && b != "G":
		return "G"
	case a != "B" && b != "B":
		return "B"
	}
	return ""
}

func main() {
	var n, cnt int
	var s, sk string
	fmt.Scan(&n)
	fmt.Scan(&s)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				continue
			}
			sk = rgb(string(s[i]), string(s[j]))
			cnt += strings.Count(string(s[j:]), sk)
			if 2*j-i < n && string(s[2*j-i]) == sk {
				cnt--
			}
		}
	}
	fmt.Println(cnt)
}
