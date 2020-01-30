package main

import "fmt"

func main() {
	var c1, c2 string
	fmt.Scan(&c1, &c2)
	if c1[0] == c2[2] && c1[1] == c2[1] && c1[2] == c2[0] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
