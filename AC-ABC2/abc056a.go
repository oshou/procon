package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	switch a {
	case "H":
		fmt.Println(b)
	case "D":
		switch b {
		case "H":
			fmt.Println("D")
		case "D":
			fmt.Println("H")
		}
	}
}
