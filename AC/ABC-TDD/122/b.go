package main

import "fmt"

func main() {
	var s string
	var cnt, max int
	fmt.Scan(&s)
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "A", "C", "G", "T":
			cnt++
			continue
		default:
		}
		if max < cnt {
			max = cnt
		}
		cnt = 0
	}
	if max < cnt {
		max = cnt
	}
	cnt = 0
	fmt.Println(max)
}
