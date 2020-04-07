package main

import "fmt"

func main() {
	var s, a, b, i int
	var win string
	fmt.Scan(&s, &a, &b)
	for s <= a && s <= b {
		i++
		if i%2 == 1 {
			if s+10 <= a {
				s += 10
				win = "A"
			} else {
				break
			}
		} else {
			if s+1000 <= b {
				s += 1000
				win = "B"
			} else {
				break
			}
		}
	}
	fmt.Println(win, s)
}
