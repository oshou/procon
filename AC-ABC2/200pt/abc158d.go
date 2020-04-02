package main

import "fmt"

func reverse(s string) string {
	ans := ""
	for i := len(s) - 1; i >= 0; i-- {
		ans += string(s[i])
	}
	return ans
}

func main() {
	var s, c string
	var q, t, f int
	var isReverse bool = false
	fmt.Scan(&s)
	fmt.Scan(&q)
	for i := 1; i <= q; i++ {
		fmt.Scan(&t)
		switch t {
		case 1:
			isReverse = !isReverse
		case 2:
			fmt.Scan(&f, &c)
			switch f {
			case 1:
				if isReverse {
					s = s + c
				} else {
					s = c + s
				}
			case 2:
				if isReverse {
					s = c + s
				} else {
					s = s + c
				}
			}
		}
	}
	if isReverse {
		s = reverse(s)
	}
	fmt.Println(s)
}
