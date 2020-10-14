package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	if string(s[len(s)-1]) != "s" {
		fmt.Println(s + "s")
	} else {
		fmt.Println(s + "es")
	}
}
