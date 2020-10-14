package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	if x == 3 || x == 5 || x == 7 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
