package main

import "fmt"

func main() {
	var s string
	var cnta, cntb, cntc int
	fmt.Scan(&s)
	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "a":
			cnta++
		case "b":
			cntb++
		case "c":
			cntc++
		}
	}
	if cnta == 1 && cntb == 1 && cntc == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
