package main

import "fmt"

func main() {
	var w string
	fmt.Scan(&w)

	m := make(map[string]int)
	for _, c := range w {
		m[string(c)]++
	}

	ans := "Yes"
	for _, cnt := range m {
		if cnt%2 != 0 {
			ans = "No"
			break
		}
	}
	fmt.Println(ans)
}
