package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	ans := "Yes"
	m := map[string]int{}
	for _, c := range s {
		m[string(c)] += 1
	}
	for _, v := range m {
		if v != 2 {
			ans = "No"
			break
		}
	}
	fmt.Println(ans)
}
