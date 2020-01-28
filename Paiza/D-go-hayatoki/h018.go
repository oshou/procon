package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	switch s {
	case "Monday", "Tuesday", "Wednesday", "Thursday":
		fmt.Printf("Still %s\n", s)
	default:
		fmt.Print("TGIF\n")
	}
}
