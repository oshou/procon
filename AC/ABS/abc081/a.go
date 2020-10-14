package main

import "fmt"

func main() {
	var s string
	var cnt int
	fmt.Scan(&s)
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "1" {
			cnt++
		}
	}
	fmt.Println(cnt)
}
