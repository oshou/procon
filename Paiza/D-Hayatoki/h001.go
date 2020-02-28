package main

import "fmt"

func main() {
	var d, s int
	fmt.Scan(&d, &s)
	if d*10/s >= 1 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
