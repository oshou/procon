package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	var n int
	var t, h string
	var tp, hp int = 0, 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s %s", &t, &h)
		for j := 0; j < min(len(t), len(h)); j++ {
			if t[j] < h[j] {
				hp += 3
				break
			}
			if t[j] > h[j] {
				tp += 3
				break
			}
			if t[j] == h[j] {
				if j < min(len(t), len(h))-1 {
					continue
				}
				switch {
				case len(t) == len(h):
					hp += 1
					tp += 1
					break
				case len(t) < len(h):
					hp += 3
					break
				case len(t) > len(h):
					tp += 3
					break
				}
			}
		}
	}
	fmt.Println(tp, hp)
}
